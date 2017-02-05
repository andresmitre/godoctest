package testpkg

import "fmt"

func untestFunc() {
	/*
		this is something other, not a doctest comment
	*/
}

func testFunc1(param string) string {
	//
	//  @test = {
	//  	{"", ""},
	//  }
	//
	return ""
}

func testFunc2(param string) string {
	/*
		@test = {
			{"", ""},
			{"foo", "bar"},
			{"baz", "quux"},
		}
	*/
	switch param {
	case "":
		return ""
	case "foo":
		return "bar"
	case "baz":
		return "quux"
	}
	return ""
}

func multiline(aVeryLongParameterOne string, aVeryLongParameterTwo float64,
	aVeryLongParameterThree string) (veryValuableReturnValue string,
	errDamn error) {
	/*
		@test = {
			{"param1", 3.14, "param2", "return value", nil},
		}
	*/
	return "return value", nil
}

func main() {
	fmt.Println("vim-go")
}
