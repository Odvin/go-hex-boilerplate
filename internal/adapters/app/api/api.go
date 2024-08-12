package api

import "github.com/Odvin/go-hex-boilerplate/internal/ports"

type Adapter struct {
	arith ports.ArithmaticPort
}

func NewAdapter(arith ports.ArithmaticPort) *Adapter {
	return &Adapter{arith: arith}
}

func (apia Adapter) GetAddition(a, b int) (int, error) {
	res, err := apia.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (apia Adapter) GetSubtraction(a, b int) (int, error) {
	res, err := apia.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (apia Adapter) GetMultiplication(a, b int) (int, error) {
	res, err := apia.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (apia Adapter) GetDivision(a, b int) (int, error) {
	res, err := apia.arith.Division(a, b)
	if err != nil {
		return 0, err
	}
	return res, nil
}
