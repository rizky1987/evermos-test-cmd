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

func CreatePayment() string{

	fmt.Println(">> Begin Create Payment <<")
	productRequst := &request.CallbackPaymentRequest{
		PaymentCode: helper.PaymentCodeTest,
		Status: helper.PaymentStatusSettlement,

	}
	postBody, _ := json.Marshal(productRequst)
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(helper.BaseURL +"/payment/callback", "application/json", responseBody)

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

	//sb := string(body)
	//fmt.Println(sb)

	commonResponse := response.CommonBaseResponse{}
	//jsonString := string(body)
	json.Unmarshal(body, &commonResponse)
	//fmt.Println(sb)

	if commonResponse.Code != 200 {
		fmt.Println("Gagal membayar tagihan")
	} else {
		fmt.Println("Pada Proses ini kita sudah berhasil membayar tagihan payment Code : ", helper.PaymentCodeTest)
	}

	fmt.Println(">> End Create Payment <<")

	return ""
}
