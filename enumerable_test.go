package enumerable

import (
	"reflect"
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

	if slicesEqual(got, testCases[0].want) {
		t.Errorf("%d: Map(%v) = %s, want = %s\n", 1, testCases[0].in, got, testCases[0].want)
	}

	got, err = Map(testCases[1].in, f2)
	if err != nil {
		t.Errorf("%s\n", err)
	}

	if slicesEqual(got, testCases[1].want) {
		t.Errorf("%d: Map(%v) = %d, want = %d\n", 2, testCases[1].in, got, testCases[1].want)
	}

	got, err = Map(testCases[2].in, f3)
	if err != nil {
		t.Errorf("%s\n", err)
	}

	if slicesEqual(got, testCases[2].want) {
		t.Errorf("%d: Map(%v) = %s, want = %s\n", 3, testCases[2].in, got, testCases[2].want)
	}

	res, err := Map(10, f1)
	if err == nil {
		t.Errorf("Map(%d) = %v, want error\n", 10, res)
	}
}

func slicesEqual(s1, s2 interface{}) bool {
	slice1 := reflect.ValueOf(s1)
	slice2 := reflect.ValueOf(s2)
	if slice1.Len() != slice2.Len() {
		return false
	}

	for index := 0; index < slice1.Len(); index++ {
		if slice1.Index(index) != slice2.Index(index) {
			return false
		}
	}
	return true
}
