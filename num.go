package num

import "strconv"
import "strings"

type Number struct {
	str        string
	number     int
	basicMap   []string
	specialMap []string
	roundsMap  []string
}

// Constructor of Number struct
// Accepts a number to be translated
// Initializes all auxiliary maps
func NewNumber(number int) *Number {
	current := new(Number)
	current.number = number
	current.basicMap = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	current.specialMap = []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	current.roundsMap = []string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	return current
}

// Slices a number in millions, thousands and units
// Returns slice of strings of maximum length of three
func (self Number) slice(str string) []string {
	var lst []string
	tmp := ""
	index := len(str) - 1
	counter := 0

	for index >= 0 {
		counter++
		tmp = string(str[index]) + tmp
		if counter == 3 {
			counter = 0
			lst = append(lst, tmp)
			tmp = ""
		}
		index--
	}
	if counter != 0 {
		lst = append(lst, tmp)
	}

	return lst
}

// Method used to parse local number not bigger than three digits
// Returns translated number
func (self Number) parse_local(val string) string {
	num, _ := strconv.Atoi(val)
	val = strconv.Itoa(num)

	first, _ := strconv.Atoi(string(int(val[0])))

	var returnValue string

	if len(val) == 3 {
		tmp, _ := strconv.Atoi(val[1:])
		if tmp != 0 {
			returnValue = self.basicMap[first] + " hundred " + self.parse_local(val[1:])
		} else {
			returnValue = self.basicMap[first] + " hundred"
		}
	} else if len(val) == 2 {
		if num < 20 {
			returnValue = self.specialMap[num-10]
		} else {
			second, _ := strconv.Atoi(string(int(val[1])))
			if second != 0 {
				returnValue = self.roundsMap[first-2] + "-" + self.parse_local(string(val[1]))
			} else {
				returnValue = self.roundsMap[first-2]
			}
		}
	} else if len(val) == 1 {
		returnValue = self.basicMap[first]
	}
	return returnValue
}

// Public method that parses the given number
// Returns translated number
func (self Number) Parse() string {
	str := strconv.Itoa(self.number)
	lst := self.slice(str)
	var final [][]string
	steps := []string{"", "thousand", "million"}

	for i, item := range lst {
		tmp, _ := strconv.Atoi(item)
		if tmp == 0 && len(final) != 0 {
			continue
		}

		result := self.parse_local(item)
		current_tupple := []string{result, steps[i]}
		final = append(final, current_tupple)
	}

	var output []string

	for i, item := range final {
		item = final[len(final)-i-1]
		if item[0] == "zero" && len(final) > 1 {
			continue
		} else {
			output = append(output, strings.Join(item, " "))
		}
	}

	return strings.TrimSpace(strings.Join(output, " "))
}
