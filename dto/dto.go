package dto

type SuccessResponse struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
}

type FailureResponse struct {
	Success bool   `json:"success"`
	Failure string `json:"failure"`
}
