package main

import (
	"errors"
	"fmt"
)

var (
	EmptyStringEror = errors.New("Name is  empty error")
	ValidationError = errors.New("Validation error")
)

func validateName(name string) error {
	if name == "" {
		return EmptyStringEror
	} else if name != "Tejas" {
		return ValidationError
	}
	return nil
}

func main() {

	name1 := "Tejas"
	name2 := "Amar"

	err := validateName(name1)

	printError(err, name1)
	err = validateName(name2)
	printError(err, name2)

}

func printError(err error, name string) {
	if err != nil {
		if errors.Is(err, EmptyStringEror) {
			fmt.Println("name is String")
		} else if errors.Is(err, ValidationError) {
			fmt.Printf("name : %s  is not validate", name)
		}
		return
	}
	fmt.Println("All right ")
}
