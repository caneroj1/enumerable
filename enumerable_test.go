package enumerable

import "testing"

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
