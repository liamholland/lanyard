package options

import "testing"

func TestInvalidOption(t *testing.T) {
	got := IsValidOptionName("fakeOption")
	want := false
	if got != want {
		t.Error()
	}
}

func TestValidOption(t *testing.T) {
	got := IsValidOptionName("start")
	want := true
	if got != want {
		t.Error()
	}
}

func TestGetOptionFromNameValid(t *testing.T) {
	got := *GetOptionFromName("start")
	want := StartDateOption
	if got.FileDeclaration != want.FileDeclaration {
		t.Error()
	}
}

func TestGetOptionFromNameInvalid(t *testing.T) {
	got := GetOptionFromName("fakeOption")
	if got != nil {
		t.Error()
	}
}
