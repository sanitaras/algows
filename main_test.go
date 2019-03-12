package main

import (
	"reflect"
	"testing"
)

// TestGetFibs testing
func TestGetFibs(t *testing.T) {

	a5 := []int{0, 1, 1, 2, 3}
	a10 := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	a15 := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377}

	b5 := getFibs(5)
	b10 := getFibs(10)
	b15 := getFibs(15)

	z5 := reflect.DeepEqual(a5, b5)
	z10 := reflect.DeepEqual(a10, b10)
	z15 := reflect.DeepEqual(a15, b15)

	if !z5 {
		t.Fatal("z5 test fail")
	}

	if !z10 {
		t.Fatal("z10 test fail")
	}

	if !z15 {
		t.Fatal("z15 test fail")
	}
}
