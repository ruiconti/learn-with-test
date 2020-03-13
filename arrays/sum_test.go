package arrays

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{30, 50}

		got := Sum(numbers)
		want := 80

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}
