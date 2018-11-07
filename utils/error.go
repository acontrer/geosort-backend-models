package utils

import "errors"

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func NewError(err error, trace string) error {
	return errors.New(trace + ": " + err.Error())
}
