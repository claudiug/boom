package calc

import (
	"testing"
)

func testAdd(t *testing.T) {
	result := AddMe(1, 3)
	if result != 4 {
		t.Errorf("this is bad. expect 4 got %d", result)
	}
}
