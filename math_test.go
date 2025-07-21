package main

import "testing"

func TestSum(t *testing.T){

	total := Sum(15,15)

	if total != 30 {
		t.Errorf("result invalid: result %d. should be: %d", total,30)
	}
}