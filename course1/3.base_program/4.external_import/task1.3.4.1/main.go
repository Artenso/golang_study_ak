package main

import (
	"github.com/shopspring/decimal"
)

func DecimalSum(a, b string) (string, error) {
	dec1, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}

	dec2, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}

	res := decimal.Sum(dec1, dec2)

	return res.String(), nil
}

func DecimalSubtract(a, b string) (string, error) {
	dec1, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}

	dec2, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}

	res := dec1.Sub(dec2)

	return res.String(), nil
}

func DecimalMultiply(a, b string) (string, error) {
	dec1, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}

	dec2, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}

	res := dec1.Mul(dec2)

	return res.String(), nil
}

func DecimalDivide(a, b string) (string, error) {
	dec1, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}

	dec2, err := decimal.NewFromString(b)
	if err != nil {
		return "", err
	}
	res := dec1.Div(dec2)

	return res.String(), nil
}

func DecimalRound(a string, precision int32) (string, error) {
	dec1, err := decimal.NewFromString(a)
	if err != nil {
		return "", err
	}

	res := dec1.Round(precision)

	return res.String(), nil
}

func DecimalGreaterThan(a, b string) (bool, error) {
	dec1, err := decimal.NewFromString(a)
	if err != nil {
		return false, err
	}

	dec2, err := decimal.NewFromString(b)
	if err != nil {
		return false, err
	}
	res := dec1.GreaterThan(dec2)

	return res, nil
}

func DecimalLessThan(a, b string) (bool, error) {
	dec1, err := decimal.NewFromString(a)
	if err != nil {
		return false, err
	}

	dec2, err := decimal.NewFromString(b)
	if err != nil {
		return false, err
	}
	res := dec1.LessThan(dec2)

	return res, nil
}

func DecimalEqual(a, b string) (bool, error) {
	dec1, err := decimal.NewFromString(a)
	if err != nil {
		return false, err
	}

	dec2, err := decimal.NewFromString(b)
	if err != nil {
		return false, err
	}
	res := dec1.Equal(dec2)

	return res, nil
}
