package request

type CreateCartRequest struct {
	ProductId		string            		`json:"product_id"`
	CustomerId		string            		`json:"customer_id"`
	Quantity		int        				`json:"quantity"`
}

type CheckoutCartRequest struct {
	Ids			    []string            	`json:"cart_ids"`
}