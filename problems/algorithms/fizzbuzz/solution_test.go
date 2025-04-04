/*
 ** ** ** ** ** **
  \ \ / / \ \ / /
   \ V /   \ V /
    | |     | |
    |_|     |_|
   Yasin   Yalcin
*/

package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	testCases := []struct {
		input    int
		expected []string
	}{
		{
			15,
			[]string{
				"1", "2", "Fizz", "4", "Buzz",
				"Fizz", "7", "8", "Fizz", "Buzz",
				"11", "Fizz", "13", "14", "FizzBuzz",
			},
		},
		{
			5,
			[]string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			3,
			[]string{"1", "2", "Fizz"},
		},
		{
			1,
			[]string{"1"},
		},
	}

	for _, tc := range testCases {
		result := FizzBuzz(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("FizzBuzz(%d) = %v; expected %v", tc.input, result, tc.expected)
		}
	}
}
