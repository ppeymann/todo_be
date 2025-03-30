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

// SignIn implements services.Accountservices.
func (l *loggingServices) SignIn(input *models.LoginInput, ctx *gin.Context) (result *todo.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "SignIn",
			"errors", strings.Join(result.Errors, " ,"),
			"input", input,
			"result", result,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.SignIn(input, ctx)
}

// ChangePassword implements services.Accountservices.
func (l *loggingServices) ChangePassword(in *models.ChangePasswordInput, ctx *gin.Context) (result *todo.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "ChangePassword",
			"errors", strings.Join(result.Errors, " ,"),
			"input", in,
			"result", result,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.ChangePassword(in, ctx)
}

// Account implements services.Accountservices.
func (l *loggingServices) Account(ctx *gin.Context) (result *todo.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "Account",
			"errors", strings.Join(result.Errors, " ,"),
			"result", result,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.Account(ctx)
}
