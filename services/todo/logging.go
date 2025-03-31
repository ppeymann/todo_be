package todos

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
	todo "github.com/ppeymann/todo_be.git"
	"github.com/ppeymann/todo_be.git/models"
)

type loggingServices struct {
	logger log.Logger
	next   models.TodoService
}

func NewLoggingServices(logger log.Logger, services models.TodoService) models.TodoService {
	return &loggingServices{
		logger: logger,
		next:   services,
	}
}

func (l *loggingServices) AddTodo(in *models.TodoInput, ctx *gin.Context) (result *todo.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "AddTodo",
			"errors", strings.Join(result.Errors, " ,"),
			"input", in,
			"result", result,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.AddTodo(in, ctx)
}

func (l *loggingServices) GetAll(ctx *gin.Context) (result *todo.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "GetAll",
			"errors", strings.Join(result.Errors, " ,"),
			"result", result,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.GetAll(ctx)
}

func (l *loggingServices) GetByID(id uint, ctx *gin.Context) (result *todo.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "GetByID",
			"errors", strings.Join(result.Errors, " ,"),
			"result", result,
			"id", id,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.GetByID(id, ctx)
}

func (l *loggingServices) DeleteTodo(id uint, ctx *gin.Context) (result *todo.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "DeleteTodo",
			"errors", strings.Join(result.Errors, " ,"),
			"result", result,
			"id", id,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.DeleteTodo(id, ctx)
}

func (l *loggingServices) UpdateTodo(in *models.TodoInput, id uint, ctx *gin.Context) (result *todo.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "UpdateTodo",
			"errors", strings.Join(result.Errors, " ,"),
			"input", in,
			"id", id,
			"result", result,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.UpdateTodo(in, id, ctx)
}

func (l *loggingServices) ChangeStatus(status string, id uint, ctx *gin.Context) (result *todo.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "ChangeStatus",
			"errors", strings.Join(result.Errors, " ,"),
			"status", status,
			"id", id,
			"result", result,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.ChangeStatus(status, id, ctx)
}
