package tpcodlElectricityCharges

import (
	"fmt"
	"math"
)

type Slab struct {
	Limit int
	Rate  float64
}

func Run() {
	fromUnitCharges := []Slab{
		{50, 3},
		{200, 4.8},
		{400, 5.8},
		{math.MaxInt, 6.2},
	}

	var UnitsConsumed = 200.0
	var MinimumCharges = 0.0
	totalCharge := MinimumCharges
	remainingUnits := UnitsConsumed

	for _, slab := range fromUnitCharges {
		if remainingUnits <= 0 {
			break
		}
		consumed := min(UnitsConsumed, float64(slab.Limit))
		remainingUnits -= consumed
		totalCharge += consumed * slab.Rate
	}

	MinimumCharges = totalCharge
	remainingUnits = calcMonthlyUnitsConsumed()

	fromUnitCharges = []Slab{
		{400, 5.8},
		{math.MaxInt, 6.2},
	}

	for _, slab := range fromUnitCharges {
		if remainingUnits <= 0 {
			break
		}
		consumed := min(UnitsConsumed, float64(slab.Limit))
		remainingUnits -= consumed
		totalCharge += consumed * slab.Rate
	}

	fmt.Println("Total Charges: ₹", totalCharge)
	fmt.Println("Total Extra Charges: ₹", totalCharge-MinimumCharges)
}

func calcMonthlyUnitsConsumed() float64 {
	const totalUnits = 72.75
	const avgPerDay = 20.02
	const totalHours = 100.12
	const perHour = totalUnits / totalHours
	const totalHoursInMonth = avgPerDay * 30
	const totalUnitsConsumed = perHour * totalHoursInMonth
	return totalUnitsConsumed
}
