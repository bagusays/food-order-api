package restspec

type HealthCheckResponse struct {
	DB HealthCheckDetail `json:"db"`
}

type HealthCheckDetail struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
