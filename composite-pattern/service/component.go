package service

type Component interface {
	Search(string)
	GetName() string
}
