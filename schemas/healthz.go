package schemas

type HealthzResponse struct {
	Status string `validate:"required" json:"status"`
}
