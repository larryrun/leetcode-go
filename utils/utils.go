package utils

import "fmt"

func AssertTrue(result bool) {
	if !result {
		panic("failed")
	}
}

func AssertEq(expected, actual interface{}) {
	if actual != expected {
		panic(fmt.Sprintf("expected: %v, but got: %v", expected, actual))
	}
}
