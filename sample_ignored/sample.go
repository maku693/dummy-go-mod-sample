package sample

//go:generate stringer -type=Animal
// Node: go:generate command above will not be invoked

type Animal int

const (
	Dog Animal = iota
	Cat
	Bird
)
