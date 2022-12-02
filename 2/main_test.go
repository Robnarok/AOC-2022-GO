package main

import "testing"

func Test_checkWin_AY(t *testing.T) {
	got := checkWin("A", "Y")
	want := 6

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_checkWin_BX(t *testing.T) {
	got := checkWin("B", "X")
	want := 0

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_checkWin_CZ(t *testing.T) {
	got := checkWin("C", "Z")
	want := 3

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_scoreChoice_X(t *testing.T) {
	got := scoreChoice("X")
	want := 1

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_scoreChoice_Y(t *testing.T) {
	got := scoreChoice("Y")
	want := 2

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_scoreChoice_Z(t *testing.T) {
	got := scoreChoice("Z")
	want := 3

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_playGame_AY(t *testing.T) {
	got := playGame("A Y")
	want := 8

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_playGame_BX(t *testing.T) {
	got := playGame("B X")
	want := 1

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_playGame_CZ(t *testing.T) {
	got := playGame("C Z")
	want := 6

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
