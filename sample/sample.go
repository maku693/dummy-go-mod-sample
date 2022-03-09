package sample

//go:generate stringer -type=Animal

type Animal int

const (
	Dog Animal = iota
	Cat
	Bird
)
