package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Homepage default function
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok!")
	fmt.Println("Endpoint: homePage")
}

// A function to insert a device to the database
func createOrder(w http.ResponseWriter, r *http.Request) {
	// Limiting the length of incoming requests to 1 MB
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	// Passing the body of our POST request into a variable
	order, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
	}

	url := "https://integration.api.scalapay.com/v2/orders"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(order))
	if err != nil {
		w.WriteHeader(500)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer qhtfs87hjnc12kkos")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(500)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var response Response
	json.Unmarshal(body, &response)

	w.Write([]byte(response.CheckoutUrl))

}
