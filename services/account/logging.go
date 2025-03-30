package account

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
	next   models.AccountService
}

func NewLoggingServices(logger log.Logger, services models.AccountService) models.AccountService {
	return &loggingServices{
		logger: logger,
		next:   services,
	}
}

// SignUp implements services.Accountservices.
func (l *loggingServices) SignUp(input *models.SignUpInput, ctx *gin.Context) (result *todo.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "SignUp",
			"errors", strings.Join(result.Errors, " ,"),
			"input", input,
			"result", result,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.SignUp(input, ctx)
}
