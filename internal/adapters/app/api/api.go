package api

import "github.com/Odvin/go-hex-boilerplate/int32ernal/ports"

type Adapter struct {
	db    ports.DbPort
	arith ports.ArithmaticPort
}

func NewAdapter(db ports.DbPort, arith ports.ArithmaticPort) *Adapter {
	return &Adapter{db: db, arith: arith}
}

func (apia Adapter) GetAddition(a, b int32) (int32, error) {
	res, err := apia.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(res, "addition")
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (apia Adapter) GetSubtraction(a, b int32) (int32, error) {
	res, err := apia.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(res, "subtraction")
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (apia Adapter) GetMultiplication(a, b int32) (int32, error) {
	res, err := apia.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(res, "multiplication")
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (apia Adapter) GetDivision(a, b int32) (int32, error) {
	res, err := apia.arith.Division(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(res, "division")
	if err != nil {
		return 0, err
	}

	return res, nil
}
