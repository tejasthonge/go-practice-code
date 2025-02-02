package main

/*
The errors package helps in:

Creating new errors
Wrapping errors with additional context
Checking error types and unwrapping errors
*/

/*
The errors package in Go provides powerful error handling tools:

errors.New() and fmt.Errorf() for creating errors
errors.Is() for comparing errors
errors.As() for type assertion
errors.Unwrap() for extracting wrapped errors
*/

import (
	"errors"
	"fmt"
	"reflect"
)



func main() {
	err := errors.New("This is custom error")

	fmt.Println(err.Error())
	
	fmt.Println(reflect.TypeOf(err)) //*errors.errorString
	fmt.Println(reflect.TypeOf(err.Error())) //string

	//fomating string error or wrapped errors

	wrappedErr := fmt.Errorf("service error: %w", err)
	fmt.Println(wrappedErr)  //service error: This is custom error


	//unwapping error
	fmt.Println(errors.Unwrap(wrappedErr)) //This is custom error

}
