package matrix

import (
	"strings"
	"strconv"
	"errors"
)

type matrix struct {
	rows [][]int 
	clos [][]int
}

func (m *matrix) Rows() [][]int {
	res := make([][]int, len(m.rows))
	for i := range m.rows {
		res[i] = append([]int{}, m.rows[i]...)
	}
	return res
}

func (m *matrix) Cols() [][]int {
	res := make([][]int, len(m.clos))
	for i := range m.clos {
		res[i] = append([]int{}, m.clos[i]...)
	}
	return res
}

func (m *matrix) Set(r, c, val int)  bool { 
	if r >= 0 && r < len(m.rows) && c >=0 &&(len(m.rows) > 0 && c < len(m.rows[0]) ) {
		m.rows[r][c] = val 
		m.clos[c][r] = val
		return true
	} 

	return false
}

func New(in string) (*matrix , error) {
	res := strings.Split(in, "\n") 
	//fmt.Println(res)
	if len(res) == 0 {
		return nil, errors.New(" line = 0")
	}

	rows := make([][]int, len(res))
	// 1.每行元素不等
	// 2.元素存在非数字
	var length int
	for i := range res {
		// such as {"1 2 3\n4 5 6\n7 8 9\n 8 7 6",
		// var index int
		// for index < len(res[i]){
		// 	if res[i][index] != ' ' {
		// 		break
		// 	}
		// 	index++
		// }
		// res[i] = res[i][index:]
		res1 := strings.Split(strings.TrimSpace(res[i]), " ")
		//res1 := strings.Split(res[i], " ") 
		
		
		if len(res1) == 0 {
			return nil, errors.New(" 0lll")
		}	
		if length == 0 {
			length = len(res1)
		}else if length != len(res1){
			
			return  nil, errors.New(" 1lll")
		}
			

		for j := range res1  { 
			
			num, err := strconv.Atoi(res1[j]) 
			if err != nil {
				return  nil, errors.New(" 2lll")
			}
			rows[i] = append(rows[i], num) 
		}
	}
	var clos [][]int
	
	if len(rows)  > 0 && len(rows[0]) > 0 {
		clos = make([][]int, len(rows[0]) ) // 3 
		for i := range clos {
			clos[i] = make([]int, len(rows)) // 2
		} 
		
		for i :=0; i < len(rows[0]); i++ {
			for j := 0; j < len(rows); j++ {
				clos[i][j] = rows[j][i]
			}
		}

	}
	

	m := matrix{
		rows:rows,
		clos:clos,
	}
	
	return &m, nil 
}