package main

import "time"

/* Struct for order details, uncomment when necessary

type Order struct {
	TotalAmount AmountQty `json:"total_amount"`
	Consumer    `json:"consumer"`
	Shipping    `json:"shipping"`
	ItemsArray  []Items `json:"items"`
	Merchant    `json:"merchant"`
}

type AmountQty struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Consumer struct {
	PhoneNumber string `json:"phone_number"`
	GivenNames  string `json:"given_names"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
}

type Shipping struct {
	PhoneNumber string `json:"phone_number"`
	CountryCode string `json:"country_code"`
	Name        string `json:"name"`
	Postcode    string `json:"postcode"`
	Line1       string `json:"line_one"`
}

type Items struct {
	Quantity string    `json:"quantity"`
	Price    AmountQty `json:"price"`
	Name     string    `json:"name"`
	Category string    `json:"category"`
	Sku      string    `json:"sku"`
}

type Merchant struct {
	RedirectCancelUrl  string `json:"redirect_cancel_url"`
	RedirectConfirmUrl string `json:"redirect_confirm_url"`
}*/

type Response struct {
	Token       string    `json:"token"`
	Expires     time.Time `json:"expires"`
	CheckoutUrl string    `json:"checkoutUrl"`
}
