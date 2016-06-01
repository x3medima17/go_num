package num

import "testing"
// import "fmt"

var numTests = []struct {
  n        int // input
  expected string // expected result
}{
  {0, "zero"},
  {999, "nine hundred ninety-nine"},
  {1, "one"},
  {10, "ten"},
  {20, "twenty"},
  {17, "seventeen"},
  {112, "one hundred twelve"},
  {736, "seven hundred thirty-six"},
  {721, "seven hundred twenty-one"},
  {100000, "one hundred thousand"},
  {721256, "seven hundred twenty-one thousand two hundred fifty-six"},
  {999999, "nine hundred ninety-nine thousand nine hundred ninety-nine"},
  {100200, "one hundred thousand two hundred"},
  {1000000, "one million"},
  {99999, "ninety-nine thousand nine hundred ninety-nine"},
  {99998, "ninety-nine thousand nine hundred ninety-eight"},
  {131313, "one hundred thirty-one thousand three hundred thirteen"},
  {121212, "one hundred twenty-one thousand two hundred twelve"},
  {111111, "one hundred eleven thousand one hundred eleven"},


}

func TestNum(t *testing.T){
	for _, tt := range numTests {
		actual := NewNumber(tt.n)
		if actual.parse() != tt.expected {
			t.Errorf("Fib(%d): expected %s\n actual %s", tt.n, tt.expected, actual)
		}
	}
}