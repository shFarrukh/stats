package stats

import (
	"fmt"

	"github.com/shFarrukh/bank/pkg/bank/types")


func ExampleTotalInCategory() {
	peyments := []types.Payment {
		{ 
			Amount : 10_000,
			Category: "car",
		},
		{
			Amount:  20_000,
			Category: "car",
		},
	}
	result := TotalInCategory(peyments, "car")
	fmt.Println(result)
	// Output
	// 15000
}