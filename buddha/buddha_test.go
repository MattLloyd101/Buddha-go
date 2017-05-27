package buddha

import "testing"

func TestScale(t *testing.T) {
	type scaletest struct {
		t float64
		srcMin float64
		srcMax float64
		targMin float64
		targMax float64
		expected float64
	}
	var tests = []scaletest{
	  { 15, 10, 20, 0,  1, 0.5 },
	  { 15, 10, 20, 0, 10, 5 },
	  { 5, 	10, 20, 0, 10, -5 },
	  { 25, 10, 20, 0, 10, 15 },
	}

	for _, data := range tests {
		var actual = scale(data.t, data.srcMin, data.srcMax, data.targMin, data.targMax)
		if (actual != data.expected) {
			t.Error(
				"scale(", 
					data.t, ",", 
					data.srcMin, ",", 
					data.srcMax, ",", 
					data.targMin, ",", 
					data.targMax, 
				") was", actual, 
				"expected ", data.expected)
		}
	}
}