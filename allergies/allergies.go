package allergies


var list = []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}

func Allergies(score uint) []string{
	res := make([]string, 0)

	for i:=0; i <=7; i++ {
		
		if (1<<uint(i)) & score > 0 {
			res = append(res, list[i])
		}
	}

	return res
}

func AllergicTo( score uint, substance string) bool {
	var need int 
	for i:=0; i <=7 ; i++ {
		if substance == list[i] {
			need = i
		}
	}

	return (1<<uint(need)) & score > 0 
}