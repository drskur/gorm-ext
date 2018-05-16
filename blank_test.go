package gorm_ext

import "testing"

func TestIsBlank(t *testing.T) {
	strA := ""
	strB := "%"
	strC := "not blank"

	intA := 0

	if !isBlank(strA) {
		t.Fail()
	}

	if !isBlank(strB) {
		t.Fail()
	}

	if isBlank(strC) {
		t.Fail()
	}

	if !isBlank(intA) {
		t.Fail()
	}

}
