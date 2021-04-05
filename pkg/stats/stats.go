package stats

import (
	"github.com/shFarrukh/bank/pkg/types"
)


func Avg(payments []types.Payment) types.Money{
	s := 0
	couter := 0
	for _, i := range payments {

		s+= int(i.Amount)
		couter ++
	}
	return types.Money(s/couter)

}
func TotalInCategory (Payments []types.Payment, category types.Category) types.Money{
	s:=0

	for _, i := range Payments{
		if i.Category == category{
			s += int(i.Amount)
			
		}
	}
	return types.Money(s)
}