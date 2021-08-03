package response

import "gopkg.in/mgo.v2/bson"

type ProductResponse struct {
	Id          		bson.ObjectId   		`json:"id"`
	Code				string            		`json:"code"`
	Name				string            		`json:"name"`
	Quantity		 	int               		`json:"quantity"`
	OnHoldQuantity		int               		`json:"on_hold_quantity"`
	SoldQuantity		int               		`json:"sold_quantity"`
	Message				string					`json:"message"`
}

type ProductSuccessResponse struct {
	CommonBaseResponse
	Data 			ProductResponse 		`json:"result"`
}
