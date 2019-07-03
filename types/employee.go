package types

//go:generate jsongen -package=types -name=e -type=Employee -pointer
type Employee struct {
	Name   string
	Age    int
	Salary float64
}
