package conversion

import (
	"errors"
	"strconv"
)

func StringToFloats(strings []string) ([]float64, error) {
	var floats []float64
	for _, stringVal := range strings {
		floatVal, err := strconv.ParseFloat(stringVal, 64)
		if err != nil {
			return nil, errors.New("failed to conver string to float")
		}

		floats = append(floats, floatVal)

	}
	return floats, nil
}
