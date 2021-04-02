package main

import (
	"bank/pkg/bank/card"
	"bank/pkg/bank/payment"
	"fmt"
)

func main() {

	cardsList := card.GenerateCard()
	fmt.Println(card.Total(cardsList))
	payments := payment.GeneratePayments()
	fmt.Println(payment.Max(payments))
}
