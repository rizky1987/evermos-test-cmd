package request

type SearchParamInventoryAdjustmentRequest struct {
	ProductId		string            		`json:"product_id"`
}

type CreateInventoryAdjustmentRequest struct {
	ProductId         		string  		`json:"product_id"`
	Process           		string        	`json:"process"`
	Quantity          		int           	`json:"quantity"`
	Note    				string        	`json:"note"`
}
