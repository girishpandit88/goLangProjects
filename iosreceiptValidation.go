package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func validateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	// log.Println(string(body))

	var t Message
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	// log.Println(t, t.receipt)
	validate(t.receipt, t.transactionId)
}

func validate(receipt, transactionId string) bool {
	hc := http.Client{}
	data := url.Values{}
	data.Set("receipt-data", receipt)

	log.Println(data)
	prodUrlStr := "https://buy.itunes.apple.com/verifyReceipt"
	sandboxUrlStr := "https://sandbox.itunes.apple.com/verifyReceipt"

	r, _ := http.NewRequest("POST", prodUrlStr, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/json")

	resp, _ := hc.Do(r)
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	respBody := string(contents)
	log.Println("resp from prod", respBody)
	if strings.Contains(respBody, "21002") {
		log.Println("Going for sandbox")
		r, _ := http.NewRequest("POST", sandboxUrlStr, strings.NewReader(data.Encode()))
		r.Header.Add("Content-Type", "application/json")

		resp, _ := hc.Do(r)
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		respBody := string(contents)
		log.Println("resp from sandbox", respBody)
		if strings.Count(respBody, "0") == 1 {
			log.Println("valid receipt")
		}

	} else {
		return true
	}
	return false
}

type Message struct {
	receipt       string
	transactionId string
}

func main() {
	http.HandleFunc("/receipts/validate", validateHandler)
	http.ListenAndServe(":8080", nil)
}
