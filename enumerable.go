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
	return fmt.Sprintf("Enumerable Error: %s", e.message)
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

func rescue2(results *interface{}, err *Error) {
	if v := recover(); v != nil {
		err = &Error{}
		switch x := v.(type) {
		case string:
			err.set(x)
		case error:
			err.set(x.Error())
		}

		*results = nil
	}
}

func rescue3(err *Error) {
	if v := recover(); v != nil {
		err = &Error{}
		switch x := v.(type) {
		case string:
			err.set(x)
		case error:
			err.set(x.Error())
		}
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

// Map accepts a slice and a function that accepts an index and a value and returns another value.
// Map executes that function on each value in the slice and stores the returned value in a new slice.
func Map(slice, function interface{}) (results interface{}, err *Error) {
	defer rescue2(&results, err)

	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(slice)

		if s.Len() < 1 {
			results = nil
		} else {
			input := []reflect.Value{
				reflect.ValueOf(0),
				s.Index(0),
			}
			output := reflect.ValueOf(function).Call(input)
			results = make([]interface{}, s.Len())
			reflect.ValueOf(results).Index(0).Set(output[0])

			for idx := 1; idx < s.Len(); idx++ {
				input := []reflect.Value{
					reflect.ValueOf(idx),
					s.Index(idx),
				}

				output := reflect.ValueOf(function).Call(input)
				reflect.ValueOf(results).Index(idx).Set(output[0])
			}
		}
	default:
		results = nil
		err = &Error{}
		(*err).set("A slice needs to be the first parameter of Map.")
	}

	return results, err
}

// Select accepts a slice and a function that accepts an index and a value and returns a bool.
// Select executes that function on each value in the slice and stores the value at index in a new slice if the function returns true.
func Select(slice, function interface{}) (results interface{}, err *Error) {
	defer rescue2(&results, err)

	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(slice)

		if s.Len() < 1 {
			results = nil
		} else {
			results = make([]interface{}, 0)

			for idx := 0; idx < s.Len(); idx++ {
				input := []reflect.Value{
					reflect.ValueOf(idx),
					s.Index(idx),
				}

				output := reflect.ValueOf(function).Call(input)
				if output[0].Interface().(bool) {
					results = append(results.([]interface{}), s.Index(idx))
				}
			}
		}
	default:
		results = nil
		err = &Error{}
		(*err).set("A slice needs to be the first parameter of Select.")
	}

	return results, err
}

// Each accepts a slice and a function that accepts an index and a value and returns a bool.
// Each executes that function on each value in the slice.
func Each(slice, function interface{}) (err *Error) {
	defer rescue3(err)

	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(slice)

		if s.Len() > 0 {
			for idx := 0; idx < s.Len(); idx++ {
				input := []reflect.Value{
					reflect.ValueOf(idx),
					s.Index(idx),
				}

				reflect.ValueOf(function).Call(input)
			}
		}
	default:
		err = &Error{}
		(*err).set("A slice needs to be the first parameter of Select.")
	}

	return err
}
