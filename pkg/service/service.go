package service

import "github.com/Calaghan1/todo-app/pkg/repository"

type Autorization interface {
}

type TodoList interface {
}
type TodoItem interface {
}

type Service struct {
	Autorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}