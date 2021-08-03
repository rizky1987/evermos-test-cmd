package response

type InventoryAdjustmentResponse struct {
	Id          		string  				`json:"_id"`
	ProductId			string            		`json:"product_id"`
	Process				string            		`json:"process"`
	Quantity		 	int               		`json:"quantity"`
	Note				string               		`json:"note"`
}

type InventoryAdjustmentSuccessResponse struct {
	CommonBaseResponse
	Data 			InventoryAdjustmentResponse 		`json:"result"`
}

type InventoryAdjustmentSuccessWithPagingResponse struct {
	CommonBaseResponse
	Result 			*InventoryAdjustmentSearchResponse `json:"result"`
}

type InventoryAdjustmentSearchResponse struct {
	Data 			[]*InventoryAdjustmentResponse 			`json:"data"`
	CurrentPage 	int 						`json:"current_page"`
	Limit 			int 						`json:"limit"`
	TotalRecords 	int 						`json:"total_records"`
	TotalPages 		int 						`json:"total_page"`
}