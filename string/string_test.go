package mystr

import "testing"

func TestReverseString(t *testing.T) {
	if ReverseString("中国人") != "人国中" {
		t.Fail()
	}
}