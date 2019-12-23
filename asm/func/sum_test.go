package sum

import "testing"

func TestSum(t *testing.T) {
	result := Sum(1,1)
	if result != 2 {
		t.Errorf("%d does not equal 2", result)
	}
}