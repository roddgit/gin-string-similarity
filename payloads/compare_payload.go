package payloads

type CompareRequest struct {
	NamePMO  string `json:"name_pmo" validate:"required"`
	NameCORE string `json:"name_core" validate:"required"`
}

type CompareResponse struct {
	StatusCode           string  `json:"status_code"`
	StatusMessage        string  `json:"status_message"`
	LogsId               string  `json:"logs_id"`
	NameMatchingTreshold float64 `json:"name_matching_threshold"`
}

type CompareFault struct {
	StatusCode    string `json:"status_code"`
	StatusMessage string `json:"status_message"`
	Message       string `json:"message"`
}
