package converter

import (
	"angles/helpers"
	"math"
	"testing"
)

func TestConvertDegToRad(t *testing.T) {
	inputs := []string{
		"0",
		"90",
		"180",
		"360",
		"a", // This should fail
	}

	results := []struct {
		val float64
		err bool
	}{
		{0.0, false},
		{math.Pi / 2, false},
		{math.Pi, false},
		{math.Pi * 2, false},
		{0.0, true},
	}

	for i := 0; i < len(inputs); i++ {
		rad, err := ConvertDegToRad(inputs[i])
		if rad != results[i].val || helpers.HasError(err) != results[i].err {
			t.Errorf("%s érték konvertálása sikertelen. %s\n", inputs[i], err)
		}
	}
}

func TestConvertRadToDeg(t *testing.T) {
	inputs := []float64{
		0.0,
		math.Pi / 2,
		math.Pi,
		math.Pi * 2,
	}
	results := []float64{0.0, 90, 180, 360}

	for i := 0; i < len(inputs); i++ {
		deg := ConvertRadToDeg(inputs[i])
		if deg != results[i] {
			t.Errorf("%f érték konvertálása sikertelen. Out: %f Excepted: %f", inputs[i], deg, results[i])
		}
	}
}
