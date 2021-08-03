package request

type CreateProductRequest struct {
	Name			string            		`json:"name"`
	Code			string            		`json:"code"`
	Quantity		int            			`json:"quantity"`
	MinimalStock	int           			`json:"minimal_stock"`
	Price		 	int               		`json:"price"`
}
