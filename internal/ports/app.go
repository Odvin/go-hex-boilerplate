package ports

type APIPort interface {
	GetAddition(a, b int) (int, error)
	GetSubtraction(a, b int) (int, error)
	GetMultiplication(a, b int) (int, error)
	GetDivision(a, b int) (int, error)
}
