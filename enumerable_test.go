package enumerable

import (
	"reflect"
	"strings"
	"testing"
)

type TestEnumerable struct {
	in   interface{}
	want interface{}
}

// TestEnumerableAll tests the enumerable package's
// 'All' function. All executes a function on each
// element of a slice and returns true if that function
// returns true for all elements of the slice.
func TestEnumerableAll(t *testing.T) {
	testCases := []TestEnumerable{
		TestEnumerable{
			[]int{1, 2, 3},
			true,
		},
		TestEnumerable{
			[]int{0, 2, 3},
			false,
		},
		TestEnumerable{
			[]int{-1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			false,
		},
	}
	f := func(i, val int) bool {
		return val > 0
	}

	for idx, testCase := range testCases {
		got, err := All(testCase.in, f)
		if err != nil {
			t.Errorf("%s\n", err)
		}

		if got != testCase.want {
			t.Errorf("%d: All(%v) = %t, want = %t\n", idx+1, testCase.in, got, testCase.want)
		}
	}

	res, err := All(10, f)
	if err == nil {
		t.Errorf("All(%d) = %v, want error\n", 10, res)
	}
}

// TestEnumerableSome tests the enumerable package's
// 'Some' function. Some executes a function on each
// element of a slice and returns true if that function
// returns true for some (at least one) element(s) of the slice.
func TestEnumerableSome(t *testing.T) {
	testCases := []TestEnumerable{
		TestEnumerable{
			[]string{"hello", "world", "letter"},
			true,
		},
		TestEnumerable{
			[]string{"dog", "cat", "hat"},
			false,
		},
		TestEnumerable{
			[]string{"", "cat", "dog", "fish", "horse"},
			true,
		},
	}

	f := func(i int, val string) bool {
		return len(val) > 4
	}

	for idx, testCase := range testCases {
		got, err := Some(testCase.in, f)
		if err != nil {
			t.Errorf("%s\n", err)
		}

		if got != testCase.want {
			t.Errorf("%d: Some(%v) = %t, want = %t\n", idx+1, testCase.in, got, testCase.want)
		}
	}

	res, err := Some(10, f)
	if err == nil {
		t.Errorf("Some(%d) = %v, want error\n", 10, res)
	}
}

// TestEnumerableMap tests the enumerable package's
// 'Map' function. Map executes a function on each
// element of a slice and stores the result of that function
// in a new slice
func TestEnumerableMap(t *testing.T) {
	testCases := []TestEnumerable{
		TestEnumerable{
			[]string{"hello", "world", "letter"},
			[]string{"hello_map", "world_map", "letter_map"},
		},
		TestEnumerable{
			[]string{"dog", "cats", "jacket"},
			[]int{3, 4, 6},
		},
		TestEnumerable{
			[]string{"", "cat", "dog", "fish", "horse"},
			[]string{"cat", "catdog", "dogfish", "fishhorse", "horse"},
		},
	}

	f1 := func(i int, val string) string {
		return val + "_map"
	}

	f2 := func(i int, val string) int {
		return len(val)
	}

	f3 := func(i int, val string) string {
		if i < reflect.ValueOf(testCases[2].in).Len()-1 {
			return val + reflect.ValueOf(testCases[2].in).Index(i+1).String()
		}
		return val
	}

	got, err := Map(testCases[0].in, f1)
	if err != nil {
		t.Errorf("%s\n", err)
	}

	if !slicesEqual(got, testCases[0].want) {
		t.Errorf("%d: Map(%v) = %s, want = %s\n", 1, testCases[0].in, got, testCases[0].want)
	}

	got, err = Map(testCases[1].in, f2)
	if err != nil {
		t.Errorf("%s\n", err)
	}

	if !slicesEqual(got, testCases[1].want) {
		t.Errorf("%d: Map(%v) = %d, want = %d\n", 2, testCases[1].in, got, testCases[1].want)
	}

	got, err = Map(testCases[2].in, f3)
	if err != nil {
		t.Errorf("%s\n", err)
	}

	if !slicesEqual(got, testCases[2].want) {
		t.Errorf("%d: Map(%v) = %s, want = %s\n", 3, testCases[2].in, got, testCases[2].want)
	}

	res, err := Map(10, f1)
	if err == nil {
		t.Errorf("Map(%d) = %v, want error\n", 10, res)
	}
}

// TestEnumerableSelect tests the enumerable package's
// 'Select' function. Select executes a function on each
// element of a slice and adds that element to a new slice
// if that function returns true.
func TestEnumerableSelect(t *testing.T) {
	testCases := []TestEnumerable{
		TestEnumerable{
			[]string{"hello", "world", "letter"},
			[]string{"hello", "world"},
		},
		TestEnumerable{
			[]string{"dog", "cats", "jacket"},
			[]string{"dog"},
		},
		TestEnumerable{
			[]int{1, 2, -1, 5, 4, 8, 27},
			[]int{2, 4, 8},
		},
	}

	f1 := func(idx int, val string) bool {
		return len(val) <= 5
	}

	f2 := func(idx int, val string) bool {
		return strings.Contains(val, "d")
	}

	f3 := func(idx int, val int) bool {
		return (val % 2) == 0
	}

	res, err := Select(testCases[0].in, f1)
	if err != nil {
		t.Errorf("Received an error: %s", err)
	}
	if !slicesEqual(res, testCases[0].want) {
		t.Errorf("%d: Select(%v) = %v, want = %v", 1, testCases[0].in, res, testCases[0].want)
	}

	res, err = Select(testCases[1].in, f2)
	if err != nil {
		t.Errorf("Received an error: %s", err)
	}
	if !slicesEqual(res, testCases[1].want) {
		t.Errorf("%d: Select(%v) = %v, want = %v", 2, testCases[1].in, res, testCases[1].want)
	}

	res, err = Select(testCases[2].in, f3)

	if err != nil {
		t.Errorf("Received an error: %s", err)
	}
	if !slicesEqual(res, testCases[2].want) {
		t.Errorf("%d: Select(%v) = %v, want = %v", 3, testCases[2].in, res, testCases[2].want)
	}
}

// TestEnumerableEach tests the enumerable package's
// 'Each' function. Each executes a function on each
// element of a slice.
func TestEnumerableEach(t *testing.T) {
	testCases := []TestEnumerable{
		TestEnumerable{
			[]string{"hello", "world", "letter"},
			"letter",
		},
		TestEnumerable{
			[]string{"dog", "cats", "jacket"},
			6,
		},
	}

	longestString := ""
	f1 := func(idx int, val string) {
		if len(val) >= len(longestString) {
			longestString = val
		}
	}

	numberOfLetters := 0
	f2 := func(idx int, val string) {
		if len(val) >= numberOfLetters {
			numberOfLetters = len(val)
		}
	}

	err := Each(testCases[0].in, f1)
	if err != nil {
		t.Errorf("Received an error: %s", err)
	}
	if longestString != testCases[0].want {
		t.Errorf("%d: Each(%v) = %v, want = %v", 1, testCases[0].in, longestString, testCases[0].want)
	}

	err = Each(testCases[1].in, f2)
	if err != nil {
		t.Errorf("Received an error: %s", err)
	}
	if numberOfLetters != testCases[1].want {
		t.Errorf("%d: Each(%v) = %v, want = %v", 2, testCases[1].in, numberOfLetters, testCases[1].want)
	}

}

func slicesEqual(s1, s2 interface{}) bool {
	slice1 := reflect.ValueOf(s1)
	slice2 := reflect.ValueOf(s2)
	if slice1.Len() != slice2.Len() {
		return false
	}

	for index := 0; index < slice1.Len(); index++ {
		v1 := slice1.Index(index).Elem()
		v2 := slice2.Index(index)
		switch v1.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val := v1.Int()
			if val != v2.Int() {
				return false
			}
		case reflect.String:
			val := v1.String()
			if val != v2.String() {
				return false
			}
		}
	}
	return true
}
