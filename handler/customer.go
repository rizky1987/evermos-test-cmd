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

func CreateCustomer() string{

	fmt.Println(">> Begin Create Customer <<")
	customerRequest := &request.CreateCustomerRequest{
		Code			: "CO.0001",
		Name			: "Rizky Mochammad Soleh",

	}
	postBody, _ := json.Marshal(customerRequest)
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(helper.BaseURL +"/customer", "application/json", responseBody)

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
		fmt.Println("User Sudah terdaftar, user code : ", customerRequest.Code)
	} else {
		fmt.Println("Pada Proses ini kita sudah berhasil membuat data customer dengan nama : ", customerRequest.Name)
	}
	fmt.Println(">> End Create Customer <<")

	return ""
}
