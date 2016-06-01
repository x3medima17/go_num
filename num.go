package num

import "fmt"
import "strconv"
import "strings"

type Number struct{
	str string
	number int
	basic_map []string
	special_map []string
	rounds []string 
}

func main() {

	var tmp []int;
	tmp = append(tmp,0,2)


	s := NewNumber(0)
	// s.parse()
	fmt.Println(s.parse())
	// fmt.Println(s.parse_local("000"))
	// fmt.Println(s.basic_map[0])

}


func NewNumber(number int) *Number{
	current := new(Number)
	current.number = number
	current.basic_map = [] string {"zero","one","two","three","four","five","six","seven","eight","nine"}
	current.special_map = [] string {"ten","eleven","twelve","thirteen","fourteen","fifteen","sixteen","seventeen","eighteen","nineteen"}
	current.rounds = [] string {"twenty","thirty","forty","fifty","sixty","seventy","eighty","ninety"}
	return current
}

func(self Number) slice(str string) []string{
	var lst []string
	tmp := ""
	index := len(str)-1
	counter := 0

	for index >= 0 {

		counter++
		tmp = string(str[index])+tmp 
		if counter == 3 {
			counter = 0
			lst = append(lst,tmp)
			tmp = ""
		}
		index--
	}
	if counter != 0{
		lst = append(lst,tmp)
	}
	return	lst

 }

func(self Number) parse_local(val string) string {
	num, _ := strconv.Atoi(val)
	val = strconv.Itoa(num)
	
	first, _ := strconv.Atoi(string(int(val[0])))
	
	if len(val) == 3{
		tmp, _ := strconv.Atoi(val[1:])
		if tmp != 0{
			return self.basic_map[first] + " hundred "+ self.parse_local(val[1:])
		} else {
			return self.basic_map[first] + " hundred"
		}
	} else if len(val) == 2 {
		if num < 20 {
			return self.special_map[num - 10]
		} else {
			second, _ := strconv.Atoi(string(int(val[1])))
			if second != 0{
				return self.rounds[first-2] + "-" + self.parse_local(string(val[1]))
			} else {	
				return self.rounds[first-2]
			}
		}
	} else if len(val) == 1{
		return self.basic_map[first]
	}
	return ""
}

func(self Number) parse() string{
	str := strconv.Itoa(self.number)
	lst := self.slice(str)
	var final [][]string
	steps := []string  {"","thousand", "million"}

	for i,item := range lst{
		tmp, _ := strconv.Atoi(item)
		if tmp == 0 && len(final) != 0{
			continue
		}

		result := self.parse_local(item)
		current_tupple := [] string {result, steps[i]}
		final = append(final, current_tupple)
	}
	// fmt.Println(final)
	// cleanup
	var output  [] string
	
	for i, item := range final {
		item = final[len(final) - i - 1 ]
		if item[0] == "zero" &&  len(final) > 1 {
			continue
		} else {
			output = append(output, strings.Join(item, " "))
		}
	}
	
	return strings.TrimSpace(strings.Join(output, " "))
}


func fact(n int) int{
	if n == 0{
		return 1
	}else{
		return fact(n-1)*n
	}
}

