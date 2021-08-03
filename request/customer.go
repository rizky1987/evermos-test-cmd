package request

type CreateCustomerRequest struct {
	Name		string            		`json:"name"`
	Code		string            		`json:"code"`
}