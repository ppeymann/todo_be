package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	todo "github.com/ppeymann/todo_be.git"
	"github.com/ppeymann/todo_be.git/auth"
	"github.com/ppeymann/todo_be.git/models"
	"github.com/ppeymann/todo_be.git/utils"
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
func (a *authorizationService) SignUp(in *models.SignUpInput, ctx *gin.Context) *todo.BaseResult {
	return a.next.SignUp(in, ctx)
}

// SignIn implements service.AccountService.
func (a *authorizationService) SignIn(in *models.LoginInput, ctx *gin.Context) *todo.BaseResult {
	return a.next.SignIn(in, ctx)
}

// ChangePassword implements service.AccountService.
func (a *authorizationService) ChangePassword(in *models.ChangePasswordInput, ctx *gin.Context) *todo.BaseResult {
	claims := &auth.Claims{}
	err := utils.CatchClaims(ctx, claims)
	if err != nil {
		return &todo.BaseResult{
			Status: http.StatusOK,
			Errors: []string{todo.AuthorizationFailed},
		}
	}

	in.Subject = claims.Subject
	return a.next.ChangePassword(in, ctx)
}
