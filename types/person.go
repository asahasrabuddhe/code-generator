package types

//go:generate jsongen -package=types -name=p -type=Person
type Person struct {
	Name string
	Age int
}
