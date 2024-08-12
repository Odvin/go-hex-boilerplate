package arithmetic

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (arith Adapter) Addition(a, b int) (int, error) {
	return a + b, nil
}

func (arith Adapter) Subtraction(a, b int) (int, error) {
	return a - b, nil
}

func (arith Adapter) Multiplication(a, b int) (int, error) {
	return a * b, nil
}

func (arith Adapter) Division(a, b int) (int, error) {
	return a / b, nil
}
