package response

type CommonBaseResponse struct {
	Message  []string `json:"message"`
	Code     int      `json:"code"`
	CodeType string   `json:"code_type"`
}
