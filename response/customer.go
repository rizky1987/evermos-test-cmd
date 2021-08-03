package response

type CustomerResponse struct {
	Id          string  				`json:"id"`
	Name		string            		`json:"name" validate:"required"`
	Code		string            		`json:"code" validate:"required"`
}

type CustomerSuccessResponse struct {
	CommonBaseResponse
	Data 			CustomerResponse 		`json:"result"`
}
