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

func (s *service) GetAll(ctx *gin.Context) *todo.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &todo.BaseResult{
			Status: http.StatusOK,
			Errors: []string{todo.AuthorizationFailed},
		}
	}

	todos, err := s.repo.GetAll(claims.Subject)
	if err != nil {
		return &todo.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	return &todo.BaseResult{
		Status: http.StatusOK,
		Result: todos,
	}
}

func (s *service) GetByID(id uint, ctx *gin.Context) *todo.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &todo.BaseResult{
			Status: http.StatusOK,
			Errors: []string{todo.AuthorizationFailed},
		}
	}

	t, err := s.repo.GetByID(id, claims.Subject)
	if err != nil {
		return &todo.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	return &todo.BaseResult{
		Status: http.StatusOK,
		Result: t,
	}
}

func (s *service) DeleteTodo(id uint, ctx *gin.Context) *todo.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &todo.BaseResult{
			Status: http.StatusOK,
			Errors: []string{todo.AuthorizationFailed},
		}
	}

	err = s.repo.DeleteTodo(id, claims.Subject)
	if err != nil {
		return &todo.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	return &todo.BaseResult{
		Status: http.StatusOK,
		Result: "Successful",
	}
}

func (s *service) UpdateTodo(in *models.TodoInput, id uint, ctx *gin.Context) *todo.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &todo.BaseResult{
			Errors: []string{todo.AuthorizationFailed},
			Status: http.StatusOK,
		}
	}

	t, err := s.repo.GetByID(id, claims.Subject)
	if err != nil {
		return &todo.BaseResult{
			Errors: []string{err.Error()},
			Status: http.StatusOK,
		}
	}

	t.Description = in.Description
	t.Title = in.Title
	t.Priority = in.Priority
	t.Status = in.Status

	err = s.repo.Update(t)
	if err != nil {
		return &todo.BaseResult{
			Errors: []string{err.Error()},
			Status: http.StatusOK,
		}
	}

	return &todo.BaseResult{
		Status: http.StatusOK,
		Result: t,
	}
}

func (s *service) ChangeStatus(status string, id uint, ctx *gin.Context) *todo.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &todo.BaseResult{
			Errors: []string{todo.AuthorizationFailed},
			Status: http.StatusOK,
		}
	}

	t, err := s.repo.GetByID(id, claims.Subject)
	if err != nil {
		return &todo.BaseResult{
			Errors: []string{err.Error()},
			Status: http.StatusOK,
		}
	}

	t.Status = status

	err = s.repo.Update(t)
	if err != nil {
		return &todo.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	return &todo.BaseResult{
		Status: http.StatusOK,
		Result: t,
	}
}
