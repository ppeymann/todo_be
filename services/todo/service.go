package todos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	todo "github.com/ppeymann/todo_be.git"
	"github.com/ppeymann/todo_be.git/auth"
	"github.com/ppeymann/todo_be.git/models"
	"github.com/ppeymann/todo_be.git/utils"
)

type service struct {
	repo models.TodoRepository
}

func NewService(repo models.TodoRepository) models.TodoService {
	return &service{
		repo: repo,
	}
}

func (s *service) AddTodo(in *models.TodoInput, ctx *gin.Context) *todo.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &todo.BaseResult{
			Status: http.StatusOK,
			Errors: []string{todo.AuthorizationFailed},
		}
	}

	task, err := s.repo.CreateTodo(in, claims.Subject)
	if err != nil {
		return &todo.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	return &todo.BaseResult{
		Status: http.StatusOK,
		Result: task,
	}
}
