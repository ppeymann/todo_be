package account

import (
	"github.com/gin-gonic/gin"
	todo "github.com/ppeymann/todo_be.git"
	"github.com/ppeymann/todo_be.git/models"
)

type authorizationService struct {
	next models.AccountService
}

func NewAuthorizationService(service models.AccountService) models.AccountService {
	return &authorizationService{
		next: service,
	}
}

// SignUp implements service.AccountService.
func (a *authorizationService) SignUp(input *models.SignUpInput, ctx *gin.Context) *todo.BaseResult {
	return a.next.SignUp(input, ctx)
}

// SignIn implements service.AccountService.
func (a *authorizationService) SignIn(input *models.LoginInput, ctx *gin.Context) *todo.BaseResult {
	return a.next.SignIn(input, ctx)
}
