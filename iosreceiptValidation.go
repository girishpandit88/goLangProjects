package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"bytes"
	"errors"
)

type AppReceipt struct {
	receipt_type 					string `json:"receipt_type"`
    adam_id 						string `json:"adam_id"`
    app_item_id 					string `json:"app_item_id"`
    bundle_id 						string `json:"bundle_id"`
	application_version 			string `json: "application_version"`
    download_id 					string `json:"download_id"`
    version_external_identifier 	string `json:"version_external_identifier"`
    receipt_creation_date 			string `json:"receipt_creation_date"`
    receipt_creation_date_ms 		string `json:"receipt_creation_date_ms"`
    receipt_creation_date_pst 		string `json:"receipt_creation_date_pst"`
    request_date 					string `json:"request_date"`
    request_date_ms 				string `json:"request_date_ms"`
    request_date_pst 				string `json:"request_date_pst"`
    original_purchase_date 			string `json:"original_purchase_date"`
    original_purchase_date_ms 		string `json:"original_purchase_date_ms"`
    original_purchase_date_pst 		string `json:"original_purchase_date_pst"`
    original_application_version 	string `json:"original_application_version"`
    in_app 							[]PurchaseReceipt `json:"in_app"`
}

type PurchaseReceipt struct {
	Quantity                  string `json:"quantity"`
	ProductId                 string `json:"product_id"`
	TransactionId             string `json:"transaction_id"`
	OriginalTransactionId     string `json:"original_transaction_id"`
	PurchaseDate              string `json:"purchase_date"`
	OriginalPurchaseDate      string `json:"original_purchase_date"`
	ExpiresDate               string `json:"expires_date"`
	AppItemId                 string `json:"app_item_id"`
	VersionExternalIdentifier string `json:"version_external_identifier"`
	WebOrderLineItemId        string `json:"web_order_line_item_id"`
}

type receiptRequestData struct {
	Receiptdata string `json:"receipt-data"`
}

const (
	prodUrlStr string = "https://buy.itunes.apple.com/verifyReceipt"
	sandboxUrlStr string = "https://sandbox.itunes.apple.com/verifyReceipt"
)

type Error struct {
	error
}

type Message struct {
	receipt 	string `json:"receipt"`
}


func validateHandler(w http.ResponseWriter, r *http.Request) {
	debug(httputil.DumpRequestOut(r, true))
	decoder := json.NewDecoder(r.Body)
	log.Println(decoder)
    var t Message   
    err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	// log.Println(string(t))

	log.Println(t, t.receipt, err)
	receipt, err := validate(t.receipt, prodUrlStr)
	if err!=nil{
		receipt, err := validate(t.receipt, sandboxUrlStr)
		log.Println(receipt)
		log.Println(err)
	}

	log.Println(receipt)
}

func debug(data []byte, err error) {
    if err == nil {
        fmt.Printf("%s\n\n", data)
    } else {
        log.Println(err)
    }
}

func validate(receipt, url string) (*AppReceipt, error) {
	log.Println("receipt: ",receipt)
	requestData, err := json.Marshal(receiptRequestData{receipt})
	if err!=nil{
		return nil, err
	}
	var t receiptRequestData
	json.Unmarshal(requestData, &t)
	log.Println("requestData: ", t, url)
	toSend :=bytes.NewBuffer(requestData)

	resp, err := http.Post(url, "application/json", toSend)

	if err!=nil{
		return nil,err
	}

	body, err := ioutil.ReadAll(resp.Body)
	
	var responseData struct {
		Status float64 `json:"status"`
		ReceiptContent *AppReceipt `json:"receipt"`
	}

	responseData.ReceiptContent = new(AppReceipt)

	err = json.Unmarshal(body, &responseData)

	if err!=nil {
		return nil,err
	}

	log.Println("resp: ",responseData.Status)
	if responseData.Status !=0 {
		return nil, &Error{errors.New("invalid receipt")}
	}

	return responseData.ReceiptContent,nil

}

func main() {
	http.HandleFunc("/receipts/validate", validateHandler)
	http.ListenAndServe(":8080", nil)
}
