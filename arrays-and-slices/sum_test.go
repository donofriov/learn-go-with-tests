package sums

import (
	"reflect"
	"testing"
)

func checkDeepEqual(t *testing.T, got, want interface{}) {
	// reflect.DeepEqual is not type safe
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		checkDeepEqual(t, got, want)
	})

}

func TestSumAll(t *testing.T) {
	numbers := []int{1, 1, 1}
	nums := []int{0, 2, 9}

	got := SumAll(numbers, nums)
	want := []int{3, 11}

	checkDeepEqual(t, got, want)
}

func TestSumAllTails(t *testing.T) {
	t.Run("make sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkDeepEqual(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkDeepEqual(t, got, want)
	})
}
