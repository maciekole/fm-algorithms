package main

import "testing"

func TestEmptyCrystalBallSearch(t *testing.T) {
	var arr []bool
	want := -1
	got, err := CrystalBallsSearch(arr)
	if want != got || err != nil {
		t.Fatalf(`CrystalBallsSearch(%t) = %d, %v, want match for %d, nil`, arr, got, err, want)
	}
}

func TestFoundFirstHalfCrystalBallSearch(t *testing.T) {
	arr := []bool{false, false, true, true, true, true}
	want := 2
	got, err := CrystalBallsSearch(arr)
	if want != got || err != nil {
		t.Fatalf(`CrystalBallsSearch(%t) = %d, %v, want match for %d, nil`, arr, got, err, want)
	}
}

func TestFoundSecondHalfCrystalBallSearch(t *testing.T) {
	arr := []bool{false, false, false, false, false, false, true, true, true, true}
	want := 6
	got, err := CrystalBallsSearch(arr)
	if want != got || err != nil {
		t.Fatalf(`CrystalBallsSearch(%t) = %d, %v, want match for %d, nil`, arr, got, err, want)
	}
}

func TestNotFoundCrystalBallSearch(t *testing.T) {
	arr := []bool{false, false, false, false, false, false}
	want := -1
	got, err := CrystalBallsSearch(arr)
	if want != got || err != nil {
		t.Fatalf(`CrystalBallsSearch(%t) = %d, %v, want match for %d, nil`, arr, got, err, want)
	}
}
