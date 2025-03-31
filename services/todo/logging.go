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
