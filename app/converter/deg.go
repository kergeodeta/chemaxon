package converter

import (
	"math"
	"strconv"
)

func ConvertDegToRad(a string) (float64, error) {
	deg, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return 0, err
	}

	return deg * math.Pi / 180, nil
}

func ConvertRadToDeg(r float64) float64 {
	return r * 180 / math.Pi
}
