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

func CreateCart() string{

	fmt.Println(">> Begin Create Cart <<")
	createCartRequest := &request.CreateCartRequest{
		ProductId		: helper.ProductIdForTest,
		CustomerId		: helper.CustomerIdForTest,
		Quantity		: 10,

	}
	postBody, _ := json.Marshal(createCartRequest)
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(helper.BaseURL +"/cart", "application/json", responseBody)

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
		fmt.Println("Gagal memasukan product ke keranjang")
	} else {
		fmt.Println("Pada Proses ini kita sudah berhasil memasukan product ke keranjang dengan Quantity : ", createCartRequest.Quantity)
	}

	//sb := string(body)
	//fmt.Println(sb)

	fmt.Println(">> End Create Cart <<")

	return ""
}

func CheckoutCreateCart() string{

	fmt.Println(">> Begin Checkout Cart <<")
	productRequst := &request.CheckoutCartRequest{
		Ids :[]string{helper.CartIdForTest},

	}
	postBody, _ := json.Marshal(productRequst)
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(helper.BaseURL +"/cart/checkout", "application/json", responseBody)

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
	//
	commonResponse := response.CommonBaseResponse{}
	//jsonString := string(body)
	json.Unmarshal(body, &commonResponse)
	//fmt.Println(sb)


	fmt.Println("Checkout Berhasil: ")


	fmt.Println(">> End Checkout Cart <<")

	return ""
}