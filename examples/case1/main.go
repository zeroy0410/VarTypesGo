// The example context insensitive analysis can not be solved.

package main

func main() {
	processBool := func(i interface{}) bool {
		return i.(bool) // type assertion
	}

	processint := func(i interface{}) int {
		return i.(int) // type assertion
	}

	processString := func(i interface{}) string {
		return i.(string) // type assertion
	}
	var boolInterface interface{} = true
	var intInterface interface{} = 1
	var stringInterface interface{} = "hello"
	process := func(i interface{}) interface{} {
		return i
	}
	boolInterface = process(boolInterface)
	intInterface = process(intInterface)
	stringInterface = process(stringInterface)
	processBool(boolInterface)
	processint(intInterface)
	processString(stringInterface)
}
