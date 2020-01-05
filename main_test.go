package main

import (
	"testing"
)

func Test_nps(t *testing.T) {
	scores :=[]int{10,7,10,10,10}
	want :=0
	got :=nps(scores)
	if want!=got {
		t.Error(scores, want, got)
	}
}