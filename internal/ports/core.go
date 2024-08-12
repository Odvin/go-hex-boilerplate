package ports

type ArithmaticPort interface {
	Addition(a, b int) (int, error)
	Subtraction(a, b int) (int, error)
	Multiplication(a, b int) (int, error)
	Division(a, b int) (int, error)
}
