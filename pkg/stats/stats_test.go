package stats

import (
	"reflect"
	"testing"

	"github.com/shFarrukh/bank/v2/pkg/types"
)

func TestCategoriesAvg(t *testing.T) {

	first := map[types.Category]types.Money{
		"food": 1000,
		"auto": 5000,
		"mobile" : 100,
	}
	second := map[types.Category]types.Money{
		"food": 2000,
		"auto": 4000,
		"mobile" : 100,
		"beauty" :100,
	}
	expected := map[types.Category]types.Money{
		"food": 1000,
		"auto": -1000,
		"mobile" : 0,
		"beauty" :100,
	}
	result := PeriodsDynamic(first, second)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}
