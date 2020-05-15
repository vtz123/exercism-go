package matrix

import (
	"strings"
	"errors"
	"strconv"
)

func New(str string) (*Matrix, error) {
	lines := strings.Split(str, "\n")

	s := make(Matrix, len(lines)) 
	length := len(strings.Split(lines[0], " "))
	for i := range s {
		nums := strings.Split(lines[i], " ")
		if len(nums) != length {
			return &Matrix{}, errors.New("false input")
		}
		s[i] = make([]int, length)

		for j := range s[i] {
			num, err := strconv.Atoi(nums[j])
			if err != nil {
				return &Matrix{}, err
			}
			s[i][j] = num
		}
	}

	return &s, nil
}


type Matrix [][]int

func (s *Matrix) Saddle() []Pair {
	res := make([]Pair, 0)

	m, _ := len(*s), len((*s)[0])

	list := make([]int, 0)
	for i := 0; i < m; i++ {
		list = Max((*s)[i])
		
		for _, k := range list {
			
			if s.ok(i,k) {
				res = append(res, Pair{i, k})
			}
		}
	}

	return res
}

func (s Matrix ) ok(i,j int) bool {
	for m :=0; m < len(s); m++ {
		
		if s[m][j] < s[i][j] {
			return false
		}
	}

	return true
}

type Pair struct {
	raw, col int
}

func Max(rows []int) []int {
	//index := -1
	max := 0
	res := make([]int , 0)
	for i := range rows {
		if rows[i] > max {
			//index = i
			max = rows[i]
		}
	}

	for i := range rows {
		if rows[i] == max {
			res = append(res, i)
		}
	}

	return res
}
