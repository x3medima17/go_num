# go_num
[![Build Status](https://travis-ci.org/x3medima17/go_num.svg?branch=master)](https://travis-ci.org/x3medima17/go_num)


This is an implemetation of Number Translation, written for iAGE as test task.
The folowing best practices were used.
  - Code documenting
  - Tests
  - Continuous Integration
  - Github
 
To use this module/package you need only `num.go` file.

Example:
```go
	
	numberToParse := 157369
	currentObj := NewNumber(numberToParse)
	result := currentObj.Parse()
	fmt.Println(result)
```
You will get `one hundred fifty-seven thousand three hundred sixty-nine`