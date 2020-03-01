package converter

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ConvertDmsToRad(a string) (float64, error) {
	valXParts := strings.Split(a, "-")
	d, err := strconv.Atoi(valXParts[0])
	if err != nil {
		return 0, err
	}

	m, err := strconv.Atoi(valXParts[1])
	if err != nil {
		return 0, err
	}

	s, err := strconv.Atoi(valXParts[2])
	if err != nil {
		return 0, err
	}

	return math.Pi * (float64(d) + float64(m)/60.0 + float64(s)/3600.0) / 180.0, nil
}

func ConvertRadToDms(r float64) string {
	deg := ConvertRadToDeg(r)
	d := math.Floor(deg)
	m := math.Floor((deg - d) * 60)
	s := (deg - d - m/60) * 3600

	return fmt.Sprintf("%d-%d-%d", int(d), int(m), int(s))
}
