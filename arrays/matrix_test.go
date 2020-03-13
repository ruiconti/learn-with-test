package arrays

import (
	"reflect"
	"testing"
)

/*
	Requirements: Create empty matrix
	Requirements: Iterate through matrix and updates specific column with
		slice as arg
*/

func TestCreateMatrix(t *testing.T) {
	t.Run("Creates 2x2 matrix", func(t *testing.T) {
		got := CreateMatrix(0)
		want := [][]float64{{0.0, 0.0}, {0.0, 0.0}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Creates an 5x5 matrix", func(t *testing.T) {
		got := CreateMatrix(5)
		want := [][]float64{
			{0.0, 0.0, 0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0, 0.0, 0.0},
			{0.0, 0.0, 0.0, 0.0, 0.0}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

/*
	Requirements: create identity matrix
*/

func TestCreateIdentityMatrix(t *testing.T) {
	t.Run("Creates identity of size N", func(t *testing.T) {
		got := CreateIdentityMatrix(3)
		want := [][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}

		if reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

/*
	Requirements: Iterate through matrix and updates specific column with
		slice as arg
*/
