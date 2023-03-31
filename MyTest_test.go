package main

import "testing"

func TestCheckname(t *testing.T) {
	got := Checkname("John")
	want := true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestUname(t *testing.T) {
	got := Uname("John")
	want := "John"
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	got = Uname("John1")
	want = ""
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	got = Uname("John2")
	want = ""
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestLetterAge(t *testing.T) {
	got := LetterAge("18")
	want := "18"
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	got = LetterAge("D18")
	want = ""
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	got = LetterAge("...")
	want = ""
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
func TestUage(t *testing.T) {
	got := Uage("18")
	want := true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	got = Uage("16")
	want = false
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	got = Uage("1")
	want = false
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
