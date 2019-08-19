package trilean

import (
	"strconv"
)

type Trilean uint8

const (
	// This will assign unset to 0, which is the default value in interpolation
	TriUnset Trilean = iota
	TriTrue
	TriFalse
)

func ToString(t Trilean) string {
	if t == TriTrue {
		return "TriTrue"
	} else if t == TriFalse {
		return "TriFalse"
	}
	return "TriUnset"
}

func FromString(s string) (Trilean, error) {
	if s == "" {
		return TriUnset, nil
	}

	b, err := strconv.ParseBool(s)
	if err != nil {
		return TriUnset, err
	} else if b == true {
		return TriTrue, nil
	} else {
		return TriFalse, nil
	}
}

func FromBool(b bool) Trilean {
	if b {
		return TriTrue
	}
	return TriFalse
}

func boolPointer(b bool) *bool {
	return &b
}

func ToBoolPointer(t Trilean) *bool {
	if t == TriTrue {
		return boolPointer(true)
	} else if t == TriFalse {
		return boolPointer(false)
	}
	return nil
}
