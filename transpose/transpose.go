package transpose


func Transpose(input []string) []string {
	result := make([][]byte, 0)

	for i, line := range input {

		for j, l := range line {
			for len(result) <= j {
				result = append(result, []byte{})
			}

			for len(result[j]) < i {
				result[j]  =  append(result[j], ' ')
			
			}

			result[j] = append(result[j], byte(l))

		}


	}
	res := make([]string, 0)

	for i := range result {
		res = append(res, string(result[i]))
	}
	return res
}