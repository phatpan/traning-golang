package roman_test

import (
	"roman"
	"testing"
)

func TestOne(t *testing.T){

	result := roman.Roman(1)

	if result != "I" {
		t.Errorf("It should be I but go %v" , result )
	}
}
