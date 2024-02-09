package main

import "testing"

func TestEmptySearch(t *testing.T) {
	var arr []byte
	valueRaw := []byte("e")
	value := valueRaw[0]
	var want bool
	got, err := Search(arr, value)
	if want != got || err != nil {
		t.Fatalf(`Search(%s, %b) = %t, %v, want match for %t, nil`, arr, value, got, err, want)
	}
}

func TestFoundSearch(t *testing.T) {
	arr := []byte("abcdefghijklmnoprstuwxyz")
	valueRaw := []byte("e")
	value := valueRaw[0]
	want := true
	got, err := Search(arr, value)
	if want != got || err != nil {
		t.Fatalf(`Search(%s, %b) = %t, %v, want match for %t, nil`, arr, value, got, err, want)
	}
}

func TestNotFoundSearch(t *testing.T) {
	arr := []byte("bcdfghjklmnprstwxz")
	valueRaw := []byte("e")
	value := valueRaw[0]
	var want bool
	got, err := Search(arr, value)
	if want != got || err != nil {
		t.Fatalf(`Search(%s, %b) = %t, %v, want match for %t, nil`, arr, value, got, err, want)
	}
}
