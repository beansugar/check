package main

import (
	"fmt"

	"github.com/beansugar/check/names"
)

func main() {
	name := &names.NameCheck{}
	name = name.DigitNumber("1234567rrc")
	fmt.Println(name.HaveNumber)
	fmt.Println(name.NumberCount)
}
