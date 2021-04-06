package stats

import (
	"github.com/shFarrukh/bank/v2/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	s := 0
	couter := 0
	for _, i := range payments {

		if i.Status != types.StatusFail {
			s += int(i.Amount)
			couter++
		}
	}
	return types.Money(s / couter)

}
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	ans := map[types.Category]types.Money{}
	for _, i := range payments {
		ans[i.Category] = TotalInCategory(payments, i.Category)
	}
	return ans
}

func TotalInCategory(Payments []types.Payment, category types.Category) types.Money {
	s := 0
	c := 0

	for _, i := range Payments {
		if i.Category == category && i.Status != types.StatusFail {
			s += int(i.Amount)
			c++;
		}
	}
	return types.Money(s/c)
}


