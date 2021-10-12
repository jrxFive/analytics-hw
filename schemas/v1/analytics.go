package v1

type AnalyticsRequest struct{}

type AnalyticMetric struct {
	Value              string  `json:"value"`
	Clicks             int64   `json:"clicks"`
	AverageClicksByDay float64 `json:"averageClicksByDay"`
}

type AnalyticMetrics struct {
	Metrics []AnalyticMetric `json:"metrics"`
}

type AnalyticsResponse struct {
	Data  map[string]AnalyticMetrics `json:"data"`
	Group string                     `json:"group"`
	Unit  string                     `json:"unit"`
	Units string                     `json:"units"`
	Facet string                     `json:"facet"`
}
