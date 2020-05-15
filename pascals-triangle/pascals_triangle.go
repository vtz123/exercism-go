package pascal

func Triangle(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}

	res := make([][]int, n)
	res[0] = []int{1}
	if n == 1 {
		return res
	}
	res[1] = []int{1,1}
	if n == 2 {
		return res
	}

	for i:=2; i < n; i++ {
		res[i] = make([]int, i+1)
		res[i][0] = 1

		for j:=1; j < len(res[i])-1; j++ {
			res[i][j] = res[i-1][j] + res[i-1][j-1]
		}

		res[i][len(res[i])-1] = 1
	}


	return res
}
