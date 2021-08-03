package handler

import (
	"bytes"
	"encoding/json"
	"evermos-test-cmd/helper"
	"evermos-test-cmd/request"
	"evermos-test-cmd/response"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func FindProduct() string{

	fmt.Println(">> Begin Inventory Adjustment Product <<")
	productRequst := &request.SearchParamInventoryAdjustmentRequest{
		ProductId: helper.ProductIdForTest,
	}
	postBody, _ := json.Marshal(productRequst)
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(helper.BaseURL + "/inventory-adjustment/find-all", "application/json", responseBody)

	//Handle Error
	if err != nil {
		return fmt.Sprintf("An Error Occured %v ", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	commonResponse := response.InventoryAdjustmentSuccessWithPagingResponse{}
	//jsonString := string(body)
	json.Unmarshal(body, &commonResponse)
	//fmt.Println(jsonString)

	if commonResponse.Code != 200 {
		fmt.Println("Terjadi kesalahan saat mengambil data")
	} else {
			fmt.Println("berhasil")

			totalProduct := 0
			for i:=0; i<len(commonResponse.Result.Data); i++ {

				message := fmt.Sprintf("Product Id %v melakukan proses %v dengan Qty %d. Catatan : %v",
					commonResponse.Result.Data[i].ProductId,
					commonResponse.Result.Data[i].Process,
					commonResponse.Result.Data[i].Quantity,
					commonResponse.Result.Data[i].Note,
					)

				if commonResponse.Result.Data[i].Process == "in" {
					totalProduct = totalProduct + commonResponse.Result.Data[i].Quantity
				} else {
					totalProduct = totalProduct - commonResponse.Result.Data[i].Quantity
				}
				fmt.Println(message)
			}

			fmt.Println("Product yang ada di gudang saat ini adalah ", totalProduct)
	}

	fmt.Println(">> End Inventory Adjustment Product <<")

	return ""
}

func CreateInventoryAdjustment(proses string) string{

	fmt.Println(">> Begin Inventory Adjustment Product <<")
	productRequst := &request.CreateInventoryAdjustmentRequest{
		ProductId: helper.ProductIdForTest,
		Note:"Test Inventory Adjustment",
		Quantity:10,
		Process:proses,
	}
	postBody, _ := json.Marshal(productRequst)
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(helper.BaseURL + "/inventory-adjustment", "application/json", responseBody)

	//Handle Error
	if err != nil {
		return fmt.Sprintf("An Error Occured %v ", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	commonResponse := response.InventoryAdjustmentSuccessWithPagingResponse{}
	//jsonString := string(body)
	json.Unmarshal(body, &commonResponse)
	//fmt.Println(jsonString)

	if commonResponse.Code != 200 {
		fmt.Println("Terjadi kesalahan saat mengambil data")
	} else {
		fmt.Println("berhasil Menambahkan inventory adjustment")
	}

	fmt.Println(">> End Inventory Adjustment Product <<")

	return ""
}
