package tempconv

import (
	"testing"
)

func TestTempConv(t *testing.T) {
	tests := []struct {
		c Celcius
		f Fahrenheit
		k Kelvin
	}{
		{-100, -148, 173.15},
		{0, 32, 273.15},
		{100, 212, 373.15},
	}

	fErr := 0.0001

	for _, test := range tests {
		if r := CToF(test.c); r != test.f {
			t.Errorf("CToF want %v got %v", test.f, r)
		}
		if r := CToK(test.c); float64(r-test.k) > fErr {
			t.Errorf("CToK want %v got %v", test.k, r)
		}
		if r := FToC(test.f); float64(r-test.c) > fErr {
			t.Errorf("FToC want %v got %v", test.c, r)
		}
		if r := FToK(test.f); float64(r-test.k) > fErr {
			t.Errorf("FToK want %v got %v", test.k, r)
		}
		if r := KToC(test.k); float64(r-test.c) > fErr {
			t.Errorf("KToC want %v got %v", test.c, r)
		}
		if r := KToF(test.k); float64(r-test.f) > fErr {
			t.Errorf("KToF want %v got %v", test.f, r)
		}
	}
}
