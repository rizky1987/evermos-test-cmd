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

func CreateProduct() string{

	fmt.Println(">> Begin Create Product <<")
	productRequst := &request.CreateProductRequest{
		Code			: "PO.0001",
		Name			: "Bedak",
		Quantity		: 100,
		Price			: 30000,
		MinimalStock	: 10,

	}
	postBody, _ := json.Marshal(productRequst)
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(helper.BaseURL +"/product", "application/json", responseBody)

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

	commonResponse := response.CommonBaseResponse{}
	//jsonString := string(body)
	json.Unmarshal(body, &commonResponse)
	//fmt.Println(sb)

	if commonResponse.Code != 200 {
		fmt.Println("Product code sudah ada")
	} else {
		fmt.Println("Pada Proses ini kita sudah berhasil membuat data product dengan nama : ", productRequst.Name)
	}

	fmt.Println(">> End Create Product <<")

	return ""
}
