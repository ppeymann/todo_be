package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	todo "github.com/ppeymann/todo_be.git"
	"github.com/ppeymann/todo_be.git/models"
	"github.com/ppeymann/todo_be.git/server"
)

type handler struct {
	service models.AccountService
}

func NewHandler(svc models.AccountService, s *server.Server) models.AccountHandler {
	handler := &handler{
		service: svc,
	}

	group := s.Router.Group("/api/v1/account")
	{
		group.POST("/signup", handler.SignUp)
		group.POST("/signin", handler.SignIn)
	}

	group.Use(s.Authenticate())
	{
		group.PATCH("/change_pass", handler.ChangePassword)
	}

	return handler
}

// SignUp handles signing up http request.
//
// @BasePath 		/api/v1/account/signup
// @Summary			signing up a new account
// @Description 	create new account with specified mobile and expected info
// @Tags 			account
// @Accept 			json
// @Produce 		json
//
// @Param			input		body		models.SignUpInput	true	"sign up input"
// @Success			200			{object}	todo.BaseResult{result=models.TokenBundleOutput}	"always returns status 200 but body contains error"
// @Router			/signup		[post]
func (h *handler) SignUp(ctx *gin.Context) {
	in := &models.SignUpInput{}

	// get input from Body
	if err := ctx.ShouldBindJSON(in); err != nil {
		ctx.JSON(http.StatusBadRequest, todo.BaseResult{
			Errors: []string{todo.ProvideRequiredJsonBody},
		})

		return
	}

	// call associated service method for expected request
	result := h.service.SignUp(in, ctx)
	ctx.JSON(result.Status, result)
}

// SignIn handles signing up http request.
//
// @BasePath 		/api/v1/account/signin
// @Summary			sign in to existing account
// @Description 	sign in to existing account with specified mobile and expected info
// @Tags 			account
// @Accept 			json
// @Produce 		json
//
// @Param			input		body		models.LoginInput	true	"sign in input"
// @Success			200			{object}	todo.BaseResult{result=models.TokenBundleOutput}	"always returns status 200 but body contains error"
// @Router			/signin		[post]
func (h *handler) SignIn(ctx *gin.Context) {
	in := &models.LoginInput{}

	// get input from Body
	if err := ctx.ShouldBindJSON(in); err != nil {
		ctx.JSON(http.StatusBadRequest, todo.BaseResult{
			Errors: []string{todo.ProvideRequiredJsonBody},
		})

		return
	}

	// call associated service method for expected request
	result := h.service.SignIn(in, ctx)
	ctx.JSON(result.Status, result)
}

// ChangePassword handles is for change password in setting http request.
//
// @BasePath 		/api/v1/account/change_pass
// @Summary			change password
// @Description 	change password with specified id
// @Tags 			account
// @Accept 			json
// @Produce 		json
//
// @Param			input		body		models.ChangePasswordInput	true	"change password input"
// @Success			200			{object}	todo.BaseResult{result=string}	"always returns status 200 but body contains error"
// @Router			/change_pass		[patch]
// @Security		Bearer Authenticate
func (h *handler) ChangePassword(ctx *gin.Context) {
	in := &models.ChangePasswordInput{}

	// get input from Body
	if err := ctx.ShouldBindJSON(in); err != nil {
		ctx.JSON(http.StatusBadRequest, todo.BaseResult{
			Errors: []string{todo.ProvideRequiredJsonBody},
		})

		return
	}

	// call associated service method for expected request
	result := h.service.ChangePassword(in, ctx)
	ctx.JSON(result.Status, result)
}
