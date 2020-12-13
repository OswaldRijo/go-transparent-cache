package view

import (
	"encoding/json"
	"fmt"
	"io"
)

type bodyPrices struct {
	ItemCodes []string `json:"item_codes"`
}

type Price struct {
	Code  string `json:"item_code"`
	Value float64 `json:"price"`
}

type Prices struct {
	Values  []Price `json:"prices"`
}


func serializePrice(itemCode string, price float64) *Price  {
	return &Price{Code: itemCode, Value: price}
}

func serializePrices(itemCode []string, price []float64) *Prices  {
	var prices Prices

	for i, code:= range itemCode {
		price := Price{Code: code, Value: price[i]}
		prices.Values = append(prices.Values, price)
	}

	return &prices
}

func deserializeBody(body io.ReadCloser) *bodyPrices  {
	var b bodyPrices
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&b)
	if err != nil{
		fmt.Printf("got error : %v \n", err)
	}

	return &b
}
