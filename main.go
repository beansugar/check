package main

import (
	"check/do"
	"fmt"
	"log"
)

func main() {
	//example  err can tell what's wrong
	ok, err := do.NewCheck("sssq412").CheckUser(3, 10)
	if err != nil {
		log.Println("error", err)
	}
	fmt.Println(ok)
	ok, err = do.NewCheck("1831qqqe").CheckPassword(5, 15, do.Number(1), do.Character(1, 3))
	if err != nil {
		log.Println(err)
	}
	fmt.Println(ok)

}
