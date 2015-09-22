package enumerable

import (
	"fmt"
	"reflect"
)

// Error is a wrapper for errors related to the enumerable package.
type Error struct {
	message string
}

func (e Error) Error() string {
	return fmt.Sprintf("Enumerable Error: %s\n", e.message)
}

func (e *Error) set(message string) {
	e.message = message
}

func rescue(result *bool, err *Error) {
	if v := recover(); v != nil {
		err = &Error{}
		switch x := v.(type) {
		case string:
			err.set(x)
		case error:
			err.set(x.Error())
		}

		*result = false
	}
}

// All accepts a slice and a function that accepts an index and a value and returns a bool.
// All executes that function on each value in the slice and returns true if the function returns
// true for every element of the slice.
func All(slice, function interface{}) (result bool, err *Error) {
	result = true

	defer rescue(&result, err)

	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(slice)

		if s.Len() < 1 {
			result = false
		} else {
			for idx := 0; idx < s.Len(); idx++ {
				input := []reflect.Value{
					reflect.ValueOf(idx),
					s.Index(idx),
				}
				output := reflect.ValueOf(function).Call(input)
				result = result && output[0].Interface().(bool)
			}
		}
	default:
		result = false
		err = &Error{}
		(*err).set("A slice needs to be the first parameter of All.")
	}

	return result, err
}

// Some accepts a slice and a function that accepts an index and a value and returns a bool.
// Some executes that function on each value in the slice and returns true if the function returns
// true for at least one element of the slice.
func Some(slice, function interface{}) (result bool, err *Error) {
	result = false

	defer rescue(&result, err)

	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(slice)

		if s.Len() < 1 {
			result = false
		} else {
			for idx := 0; idx < s.Len(); idx++ {
				input := []reflect.Value{
					reflect.ValueOf(idx),
					s.Index(idx),
				}
				output := reflect.ValueOf(function).Call(input)
				result = result || output[0].Interface().(bool)
			}
		}
	default:
		result = false
		err = &Error{}
		(*err).set("A slice needs to be the first parameter of Some.")
	}

	return result, err
}