package ocr

import (
	"strings"
	"fmt"
)

var tmp = map[string]string{
	zero: "0", one: "1", two: "2", three: "3", four: "4",
	five: "5", six: "6", seven: "7", eight: "8", nine: "9",
}

//zero, one, two, three, etc. define the OCR digits
const zero string = `
 _ 
| |
|_|
   `
const one string = `
   
  |
  |
   `
const two string = `
 _ 
 _|
|_ 
   `
const three string = `
 _ 
 _|
 _|
   `
const four string = `
   
|_|
  |
   `
const five string = `
 _ 
|_ 
 _|
   `
const six string = `
 _ 
|_ 
|_|
   `
const seven string = `
 _ 
  |
  |
   `
const eight string = `
 _ 
|_|
|_|
   `
const nine string = `
 _ 
|_|
 _|
   `

func recognizeDigit(in string) string {

	if val,ok := tmp[in]; ok {
		return val
	}
	return "?"
}

func Recognize(ins string) []string {
	str := make([]string, 0)
	test := strings.Split(ins , "\n")
	
	
	for i := 0; i < len(test)-1; i=i+4 {
		
		str = append(str, fmt.Sprintf("\n%s\n%s\n%s\n", test[i+1], test[i+2], test[i+3] ) )
	}
	

	res := make([]string, 0)

	for i := range str {
		res = append(res, helper(str[i]))
	}

	
	return res
}

func helper(in string) string{
	out := Split(in)
	var res strings.Builder
	for i:= 0; i < len(out); i++ {
		
		res.WriteString( recognizeDigit(out[i]))
	}

	
	return res.String()
}

func Split(in string) []string{
	
	splitStr := strings.Split( in, "\n") // 5 line
	
	res := make([]string, len(splitStr[1])/3)

	for i := 0; i < len(splitStr[1])/3; i++ {
		res[i] = fmt.Sprintf("\n%s\n%s\n%s\n   ", splitStr[1][3*i:3*i+3], splitStr[2][3*i:3*i+3], splitStr[3][3*i:3*i+3])
	}
	

	return res
}
