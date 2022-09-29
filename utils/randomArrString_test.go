package utils

import (
	"testing"
)

func TestRandomArrString(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Logf("字符串:%v", RandomArrString(10))
	}
}
