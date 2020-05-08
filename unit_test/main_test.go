package main

import "testing"

func testCalc(t *testing.T) {
	if calc(2) !=4 {
		t.Error("Ërror expected 2+2 = 4")
	}
}
