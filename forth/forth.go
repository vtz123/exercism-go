package forth

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Forth(in []string) ([]int, error) {
	var input []string
	for i := range in {
		in[i] = strings.ToLower(in[i])
	}

	if len(in) == 1 {
		input = strings.Split(in[0]," ")
	} else if len(in) >= 2 {
		input = strings.Split(replace(in)," ")
	}

	return deal(input)
}

func deal(input []string) ([]int, error) {
	output := make([]int, 0)

	for i := 0; i < len(input); i++ {
		num, err := strconv.Atoi(input[i])

		if err == nil {
			output = append(output, num)
			continue
		}
		switch input[i] {
		case "+","-","*","/": if len(output) >= 2 {
			last := len(output)-1
			output[last-1],err  = helper(output[last-1], output[last], input[i])
			if err != nil {
				return []int{}, errors.New(" integer divide by zero")
			}
			output = output[:last]
		}else {
			return []int{}, errors.New("not enough numbers")
		}

		case "dup": if len(output) >= 1 {
			output = append(output, output[len(output)-1])
		}else {
			return []int{}, errors.New("not enough numbers")
		}

		case "drop": if len(output) >= 1 {
			output =  output[:len(output)-1]
		}else {
			return []int{}, errors.New("not enough numbers")
		}

		case "swap": if len(output) >= 2 {
			output[len(output)-1], output[len(output)-2] = output[len(output)-2], output[len(output)-1]
		}else {
			return []int{}, errors.New("not enough numbers")
		}

		case "over": if len(output) < 2 {
			return []int{}, errors.New("not enough numbers")
		} else {
			output = append(output, output[len(output)-2])
		}

		default:
			return []int{}, errors.New("panic")
		}
	}



	return output, nil
}

func helper(a1, a2 int, cha string) (int, error){
	switch cha {
	case "+": return a1+a2, nil
	case "-": return a1-a2, nil
	case "*": return a1*a2, nil
	case "/": if a2 == 0 {
		return 0, errors.New("integer divide by zero")
	}
		return a1/a2, nil
	}
	return 0, errors.New("panic")
}

func replace(in []string) string {
	var res string
	tmp := make(map[string]string)
	for i := 0; i < len(in)-1; i++ {
		kcopy,vcopy := find(in[i])
		for k,v := range tmp {
			vcopy = strings.ReplaceAll(vcopy, k, v)
		}
		in[i] = fmt.Sprintf(": %s %s ;",kcopy, vcopy)

		k,v := find(in[i])
		tmp[k] = v
		var v1 = v
		for {

			if _,ok := tmp[v1]; ok {
				tmp[k] = tmp[v1]
				v1 = tmp[k]
			} else {
				break
			}
		}
	}

	res = in[len(in)-1]
	for k,v := range tmp {
		res = strings.ReplaceAll(res, k, v)
	}
	return res
}

func find(res string) (k, v string) {
	str := strings.Split(res, " ")

	return str[1], res[len(str[1])+3: len(res)-2 ]
}