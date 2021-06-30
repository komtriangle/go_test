package main

import (
	"errors"
	"strconv"
)

func ValidateParameter(param string) (int, error) {
	if param == "" {
		return 0, errors.New("parameter is emptry; ")
	}
	intValue, err := strconv.Atoi(param)
	if err != nil {
		return 0, errors.New("can not convert parameter to int; ")
	}
	return intValue, nil
}
