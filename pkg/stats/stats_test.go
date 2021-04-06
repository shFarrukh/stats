package stats

import (
	"reflect"
	"testing"

	"github.com/shFarrukh/bank/v2/pkg/types"
)

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{
			Category: "food",
			Amount:   1000,
		},
		{
			Category: "food",
			Amount:   1000,
		},
		{
			Category: "auto",
			Amount:   5000,
		},
		{
			Category: "auto",
			Amount:   5000,
		},
	}
	expected := map[types.Category]types.Money{
		"food": 1000,
		"auto": 5000,
	}
	result := CategoriesAvg(payments)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}
