package summultiples

import (
	"sort"
)

func SumMultiples(limit int, divisors ...int) int {
	sort.Ints(divisors)

	// remove zero number
	var index = -1
	for i := 0; i < len(divisors); i++ {
		if divisors[i] == 0 {
			index = i
		}else {
			break
		}
	}
	
	divisors = divisors[index+1:]
	

	window := make([]int, len(divisors) )
	for i := range window {
		window[i] = 1
	}

	sum := 0
	
	for {
		var min = limit

		for i := 0; i < len(divisors); i++ {
			if  divisors[i]*window[i] < min {
				min = divisors[i]*window[i]
			}
		}
		
		if min >= limit {
			break
		}

		for i:=0; i < len(divisors); i++ {
			if min == divisors[i]*window[i] {
				window[i]++
				
			}
		}
		
		sum += min
	}
	

	return sum
}