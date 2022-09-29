package main

import (
	"reflect"
	"testing"
)

func LargestNumber(t *testing.T) {
	got := []int{1, 2, 3, 5, 6, 9, 10}

	want := []int{10}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
