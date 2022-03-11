package sample

//go:generate stringer -type=Animal
// Note: go:generate command above will not be invoked

type Animal int

const (
	Dog Animal = iota
	Cat
	Bird
)
