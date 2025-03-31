package todo

import (
	"github.com/gin-gonic/gin"
	todo "github.com/ppeymann/todo_be.git"
	"github.com/ppeymann/todo_be.git/models"
)

type authorizationService struct {
	next models.TodoService
}

func NewAuthorizationService(service models.TodoService) models.TodoService {
	return &authorizationService{
		next: service,
	}
}

func (a *authorizationService) AddTodo(in *models.TodoInput, ctx *gin.Context) *todo.BaseResult {
	return a.next.AddTodo(in, ctx)
}
