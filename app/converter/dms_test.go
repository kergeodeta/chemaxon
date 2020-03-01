package converter

import (
	"angles/helpers"
	"math"
	"testing"
)

func TestConvertDmsToRad(t *testing.T) {
	inputs := []string{
		"182-12-33",
		"180-0-0",
		"94-23-22",
		"a",
	}

	results := []struct {
		val float64
		err bool
	}{
		{3.180149886, false},
		{math.Pi, false},
		{1.647406585, false},
		{0, true},
	}

	for i := 0; i < len(inputs); i++ {
		rad, _ := ConvertDmsToRad(inputs[i])
		if !helpers.CompareWithTreshold(rad, results[i].val, 1e-8) {
			t.Errorf("%s érték konvertálása sikertelen. Elvárt: %f Aktuális: %f.", inputs[i], results[i].val, rad)
		}
	}
}

func TestConvertRadToDms(t *testing.T) {
	inputs := []float64{
		1.789654,
		2.134567,
		3.987654,
		4.123987,
	}
	results := []string{
		"102-32-22",
		"122-18-6",
		"228-28-32",
		"236-17-13",
	}

	for i := 0; i < len(inputs); i++ {
		dms := ConvertRadToDms(inputs[i])
		if dms != results[i] {
			t.Errorf("%f érték konvertálása sikertelen. Out: %s Excepted: %s", inputs[i], dms, results[i])
		}
	}
}
