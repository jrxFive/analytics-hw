package v1

import (
	"context"
	"net/http"
	"sync"

	"github.com/jrxFive/analytics-hw/types"
	"github.com/spf13/cast"

	"github.com/jrxFive/analytics-hw/internal/settings"
	schemas "github.com/jrxFive/analytics-hw/schemas/v1"

	openapi "github.com/jrxFive/analytics-hw/pkg/bitly"
	"github.com/labstack/echo/v4"
)

const (
	// QsGroup group querystring parameter
	QsGroup = "group"

	// QsUnit unit querystring parameter
	QsUnit = "unit"

	// QsUnits units querystring parameter
	QsUnits = "units"

	// FacetCountry only country is supported currently
	FacetCountry = "country"
)

var (
	// PaginationChannelBuffer size of buffer channel when obtaining group bitlinks
	PaginationChannelBuffer = 100

	// DefaultQsParams currently supported QueryString parameters and their defaults
	DefaultQsParams = map[string]string{
		QsGroup: "default",
		QsUnit:  "day",
		QsUnits: "30",
	}
)

type Analytics struct {
	client   *openapi.APIClient
	settings settings.Settings
}

func NewAnalytics(client *openapi.APIClient, s settings.Settings) Analytics {
	return Analytics{
		client:   client,
		settings: s,
	}
}

func (a *Analytics) Initialize(g *echo.Group) {
	g.GET("/countries", a.GetByCountries)
}

func (a *Analytics) getQsParamOrDefault(param string, c echo.Context) string {
	var p string

	if p = c.QueryParams().Get(param); p == "" {
		return DefaultQsParams[param]
	}

	return p
}

// fetchUser obtain user associated with api key
func (a *Analytics) fetchUser(ctx context.Context) (openapi.User, error) {
	userRequest := a.client.UserApi.GetUser(ctx)
	user, response, err := a.client.UserApi.GetUserExecute(userRequest)
	defer response.Body.Close()
	if err != nil {
		return user, err
	}

	return user, nil
}

// getGroupBitLinksWithChannel fetch all bitlinks of given groupId. The response can potentially be paginated (when using default size).
// Use Channel implementation to avoid waiting for all results to be obtained to further process
func (a *Analytics) getGroupBitLinksWithChannel(ctx context.Context, groupID string, ch chan PaginatedGroupBitLinksResponse) {
	groupBitlinksRequest := a.client.BitlinksApi.GetBitlinksByGroup(ctx, groupID)

	// Fetch All Group Bitlinks, this can potentially be paginated based on the number of links
	// Use Channel implementation to avoid waiting for all results
	pagintedGroupLinksRequest := PaginatedGroupBitLinksRequest{req: groupBitlinksRequest}

	pagintedGroupLinksRequest.ChannelResponse(ch, func() (openapi.Bitlinks, *http.Response, error) {
		bitlinks, response, err := a.client.BitlinksApi.GetBitlinksByGroupExecute(groupBitlinksRequest)
		if err != nil {
			a.settings.Logger.Warn(err)
			return bitlinks, response, err
		}

		return bitlinks, response, err
	})
}

// Iterate over channel of group's bitlinks using waitgroup to do concurrent requests
func (a *Analytics) getGroupBitLinkCountryMetricsFromChannel(ctx context.Context, timeUnit openapi.TimeUnit, units string, linksChannel chan PaginatedGroupBitLinksResponse) (map[string]schemas.AnalyticMetrics, chan error) {
	wg := sync.WaitGroup{}
	mu := sync.RWMutex{}
	wgErr := make(chan error, cap(linksChannel))
	defer close(wgErr)

	aggregatedBitLinkMetrics := make(map[string]schemas.AnalyticMetrics)

	for bitlink := range linksChannel {
		wg.Add(1)

		go func(b PaginatedGroupBitLinksResponse) {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					a.settings.Logger.Warn("panic occurred", r)
					wgErr <- echo.NewHTTPError(http.StatusInternalServerError)
				}
			}()

			if b.Err != nil {
				a.settings.Logger.Error(b.Err)
				wgErr <- echo.NewHTTPError(http.StatusInternalServerError, b.Err.Error())
			}

			bitlinkMetricsCountriesRequest := a.client.BitlinksApi.GetMetricsForBitlinkByCountries(ctx, *b.Bitlink.Id)
			bitlinkMetricsCountriesRequest = bitlinkMetricsCountriesRequest.Unit(timeUnit)
			bitlinkMetricsCountriesRequest = bitlinkMetricsCountriesRequest.Units(cast.ToInt32(units))

			bitlinkClickMetrics, response, err := a.client.BitlinksApi.GetMetricsForBitlinkByCountriesExecute(bitlinkMetricsCountriesRequest)
			response.Body.Close()
			if err != nil {
				a.settings.Logger.Error(err)
				wgErr <- echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			// If bitlink has no click no metrics are given in response
			for _, bitlinkMetric := range *bitlinkClickMetrics.Metrics {
				mu.RLock()
				v, ok := aggregatedBitLinkMetrics[*b.Bitlink.Id]
				mu.RUnlock()

				if ok {
					mu.Lock()
					v.Metrics = append(v.Metrics, schemas.AnalyticMetric{
						Value:              *bitlinkMetric.Value,
						Clicks:             cast.ToInt64(*bitlinkMetric.Clicks),
						AverageClicksByDay: cast.ToFloat64(cast.ToFloat64(*bitlinkMetric.Clicks) / cast.ToFloat64(units)),
					})

					aggregatedBitLinkMetrics[*b.Bitlink.Id] = schemas.AnalyticMetrics{
						Metrics: v.Metrics,
					}
					mu.Unlock()
				} else {
					mu.Lock()
					aggregatedBitLinkMetrics[*b.Bitlink.Id] = schemas.AnalyticMetrics{
						Metrics: []schemas.AnalyticMetric{
							{
								Value:              *bitlinkMetric.Value,
								Clicks:             cast.ToInt64(*bitlinkMetric.Clicks),
								AverageClicksByDay: cast.ToFloat64(cast.ToFloat64(*bitlinkMetric.Clicks) / cast.ToFloat64(units)),
							},
						},
					}
					mu.Unlock()
				}
			}
		}(bitlink)
	}

	wg.Wait()
	return aggregatedBitLinkMetrics, wgErr
}

func (a *Analytics) GetByCountries(c echo.Context) error {
	var (
		apiKey   string
		group    string
		unit     string
		units    string
		timeUnit *openapi.TimeUnit
	)

	// Middleware will add to context to be able to fetch key for individual requests
	// Use context.Background, http.Server is configured to have upperlimit to write response
	// If this were to change use WithDeadline/WithCancel
	apiKey = c.Get(types.ContextBitlyAPIKey).(string)

	// obtain params and defaults if not given
	group = a.getQsParamOrDefault(QsGroup, c)
	unit = a.getQsParamOrDefault(QsUnit, c)
	units = a.getQsParamOrDefault(QsUnits, c)

	// verify that given parameter is a valid time unit
	timeUnit, err := openapi.NewTimeUnitFromValue(unit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, types.ErrInvalidUnit)
	}

	// initialize response struct
	response := schemas.AnalyticsResponse{
		Data:  map[string]schemas.AnalyticMetrics{},
		Unit:  unit,
		Units: units,
		Facet: FacetCountry,
	}

	// group Querystring parameter was not given or specified (default), find User's default group
	if group == DefaultQsParams[QsGroup] {
		ctx := context.WithValue(context.Background(), openapi.ContextAccessToken, apiKey)
		user, err := a.fetchUser(ctx)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		response.Group = *user.DefaultGroupGuid
	} else {
		response.Group = group
	}

	ctx := context.WithValue(context.Background(), openapi.ContextAccessToken, apiKey)
	allGroupBitLinksChannel := make(chan PaginatedGroupBitLinksResponse, PaginationChannelBuffer)
	a.getGroupBitLinksWithChannel(ctx, response.Group, allGroupBitLinksChannel)

	ctx = context.WithValue(context.Background(), openapi.ContextAccessToken, apiKey)
	aggregatedBitLinkMetrics, errs := a.getGroupBitLinkCountryMetricsFromChannel(ctx, *timeUnit, units, allGroupBitLinksChannel)
	if len(errs) > 0 {
		errorMsgSlice := make([]string, 0, len(errs)-1)

		for err := range errs {
			errorMsgSlice = append(errorMsgSlice, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, errorMsgSlice)
	}

	response.Data = aggregatedBitLinkMetrics
	return c.JSON(http.StatusOK, response)
}
