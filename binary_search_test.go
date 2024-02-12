package main

import "testing"

func TestEmptyBinarySearch(t *testing.T) {
	var arr []int
	value := 1
	var want bool
	got, err := BinarySearch(arr, value)
	if want != got || err != nil {
		t.Fatalf(`BinarySearch(%d, %b) = %t, %v, want match for %t, nil`, arr, value, got, err, want)
	}
}

func TestFoundFirstHalfBinarySearch(t *testing.T) {
	arr := []int{3, 6, 9, 10, 11, 12, 13, 14, 15, 16}
	value := 10
	want := true
	got, err := BinarySearch(arr, value)
	if want != got || err != nil {
		t.Fatalf(`BinarySearch(%d, %b) = %t, %v, want match for %t, nil`, arr, value, got, err, want)
	}
}

func TestFoundSecondHalfBinarySearch(t *testing.T) {
	arr := []int{3, 6, 9, 10, 11, 12, 16, 18, 20, 22}
	value := 20
	want := true
	got, err := BinarySearch(arr, value)
	if want != got || err != nil {
		t.Fatalf(`BinarySearch(%d, %b) = %t, %v, want match for %t, nil`, arr, value, got, err, want)
	}
}

func TestNotFoundFirstHalfBinarySearch(t *testing.T) {
	arr := []int{3, 6, 9, 10, 11, 12, 13, 14, 15, 16}
	value := 20
	var want bool
	got, err := BinarySearch(arr, value)
	if want != got || err != nil {
		t.Fatalf(`BinarySearch(%d, %b) = %t, %v, want match for %t, nil`, arr, value, got, err, want)
	}
}
