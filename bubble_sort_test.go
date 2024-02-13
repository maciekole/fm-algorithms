package main

import (
	"reflect"
	"testing"
)

func TestEmptyBubbleSort(t *testing.T) {
	var arr []int
	var want []int
	got, err := BubbleSort(arr)
	if !reflect.DeepEqual(want, got) || err != nil {
		t.Fatalf(`BubbleSort(%d) = %d, %v, want match for %d, nil`, arr, got, err, want)
	}
}

func TestSortedBubbleSort(t *testing.T) {
	arr := []int{3, 6, 9, 10, 11, 12, 13, 14, 15, 16}
	want := []int{3, 6, 9, 10, 11, 12, 13, 14, 15, 16}
	got, err := BubbleSort(arr)
	if !reflect.DeepEqual(want, got) || err != nil {
		t.Fatalf(`BubbleSort(%d) = %d, %v, want match for %d, nil`, arr, got, err, want)
	}
}

func TestReversedBubbleSort(t *testing.T) {
	arr := []int{16, 15, 14, 13, 12, 11, 10, 9, 6, 3}
	want := []int{3, 6, 9, 10, 11, 12, 13, 14, 15, 16}
	got, err := BubbleSort(arr)
	if !reflect.DeepEqual(want, got) || err != nil {
		t.Fatalf(`BubbleSort(%d) = %d, %v, want match for %d, nil`, arr, got, err, want)
	}
}

func TestBubbleSort(t *testing.T) {
	arr := []int{16, 12, 3, 6, 10, 9, 11, 15, 13, 14}
	want := []int{3, 6, 9, 10, 11, 12, 13, 14, 15, 16}
	got, err := BubbleSort(arr)
	if !reflect.DeepEqual(want, got) || err != nil {
		t.Fatalf(`BubbleSort(%d) = %d, %v, want match for %d, nil`, arr, got, err, want)
	}
}
