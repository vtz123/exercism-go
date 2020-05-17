package secret
 
var steps = []string{"wink", "double blink", "close your eyes", "jump"}

func Handshake(code uint)  []string{
	res := make([]string, 0) 

	if code == 0 {
		return res
	}

	for s, step := range steps {
		if (1<<uint(s))&code > 0 {
			res = append(res, step)
		}
	}

	if (1<<uint(len(steps)))&code > 0 {
		reverse(res)
	}
	return res
}

func reverse(strings []string) {
	for i, j := 0, len(strings)-1; i < j; i, j = i+1, j-1 {
		strings[i], strings[j] = strings[j], strings[i]
	}
}