package payment

import (
	"bank/pkg/bank/card"
	"bank/pkg/bank/types"
	"fmt"
)

func PaymentTesting() *[]types.Card {
	return card.GenerateCard()
}

func ExamplePaymentSources() {
	data := *PaymentTesting()
	for _, elem := range PaymentSources(data) {
		fmt.Println(elem.Number)
	}
	// Output:
	// 123 456 789
	// 123 222 789
	// 123 999 789
}
