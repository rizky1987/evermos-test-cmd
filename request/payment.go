package request

type CallbackPaymentRequest struct {
	PaymentCode		string            		`json:"payment_code"`
	Status			string            		`json:"status"`
}
