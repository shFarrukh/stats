package stats

import (
	"github.com/shFarrukh/bank/v2/pkg/types"
)


func Avg(payments []types.Payment) types.Money{
	s := 0
	couter := 0
	for _, i := range payments {
		
		if i.Status != types.StatusFail{
			s+= int(i.Amount)
		couter ++
		}
	}
	return types.Money(s/couter)

}
func TotalInCategory (Payments []types.Payment, category types.Category) types.Money{
	s:=0

	for _, i := range Payments{
		if i.Category == category && i.Status != types.StatusFail{
			s += int(i.Amount)
			
		}
	}
	return types.Money(s)
}