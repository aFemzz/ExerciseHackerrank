package handlers

import (
	"sort"
)

func GetMaximumRevenue(quantity []int, m int) int64 {
	// reverse the qty to be descendingorder
	sort.Sort(sort.Reverse(sort.IntSlice(quantity)))

	// fmt.Println("qty :", quantity)

	var revenue int64
	for i := 0; i < m; i++ {
		if len(quantity) == 0 {
			break
		}

		maxQuantity := quantity[0]
		revenue += int64(maxQuantity)

		quantity[0]--
		if quantity[0] == 0 {

			quantity = quantity[1:]
		} else {

			sort.Sort(sort.Reverse(sort.IntSlice(quantity)))
		}
	}

	return revenue
}
