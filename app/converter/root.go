package converter

import (
	"fmt"
	"strconv"
)

func ConvertToRad(a, t string) (float64, error) {
	var rad float64
	var err error = nil

	switch t {
	case "dms":
		rad, err = ConvertDmsToRad(a)
	case "deg":
		rad, err = ConvertDegToRad(a)
	case "rad":
		rad, err = strconv.ParseFloat(a, 64)
	default:
		return 0, fmt.Errorf("incorrect format")
	}

	return rad, err
}
