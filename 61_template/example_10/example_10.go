package main

import (
	"bytes"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"html/template"
	"os"
)

type Event struct {
	AggregateID      string `json:"aggregate_id"`
	AggregateVersion int    `json:"aggregate_version"`
	EventID          string `json:"event_id"`
}

type ProductOfferTest struct {
	Event
	OfferID           uuid.UUID       `json:"offer_id"`
	SourceID          string          `json:"source_id"`
	SellerID          string          `json:"seller_id"`
	SourceProductCode string          `json:"source_product_code"`
	SourceOfferCode   decimal.Decimal `json:"source_offer_code"`
	State             string          `json:"state"`
	OldState          string          `json:"old_state"`
	PrevSate          string          `json:"prev_state"`
}

func main() {
	// TEMPLATE
	var result bytes.Buffer
	fileName := "teste.json.template"

	decimal, _ := decimal.NewFromString("123")

	test := &ProductOfferTest{
		Event: Event{
			AggregateID:      "1",
			AggregateVersion: 1,
			EventID:          "2",
		},
		OfferID:           uuid.NewV4(),
		SellerID:          "4",
		SourceID:          "5",
		SourceOfferCode:   decimal,
		SourceProductCode: "123",
		State:             "disabled",
	}

	dir, err := os.Getwd()

	template, err := template.New(fileName).ParseFiles(dir + "/61_template/example_10/" + fileName)
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	template.ExecuteTemplate(&result, fileName, test)
	fmt.Println("--------------------\n\n\n\nRESULT", result.String(), "\n\n\n\n\n------------")
}
