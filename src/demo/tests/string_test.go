package test

import (
	"strconv"
	"testing"
)

func TestStringtoInt(t *testing.T) {
	str := "8"
	i, err := strconv.Atoi(str)
	if err != nil {
		t.Error(err)
	}
	if i != 8 {
		t.Errorf("i:%v!=8", i)
	}
}
