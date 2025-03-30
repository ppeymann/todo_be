package account

import (
	"github.com/gin-gonic/gin"
	todo "github.com/ppeymann/todo_be.git"
	"github.com/ppeymann/todo_be.git/models"
	validations "github.com/ppeymann/todo_be.git/validation"
)

type validationService struct {
	schemas map[string][]byte
	next    models.AccountService
}

func NewValidationService(schemaPath string, service models.AccountService) (models.AccountService, error) {
	schemas := make(map[string][]byte)
	err := validations.LoadSchema(schemaPath, schemas)
	if err != nil {
		return nil, err
	}

	return &validationService{
		schemas: schemas,
		next:    service,
	}, nil
}

// SignUp implements services.AccountService.
func (v *validationService) SignUp(input *models.SignUpInput, ctx *gin.Context) *todo.BaseResult {
	err := validations.Validate(input, v.schemas)
	if err != nil {
		return err
	}

	return v.next.SignUp(input, ctx)
}

// SignIn implements services.AccountService.
func (v *validationService) SignIn(input *models.LoginInput, ctx *gin.Context) *todo.BaseResult {
	err := validations.Validate(input, v.schemas)
	if err != nil {
		return err
	}

	return v.next.SignIn(input, ctx)
}

// ChangePassword implements services.AccountService.
func (v *validationService) ChangePassword(input *models.ChangePasswordInput, ctx *gin.Context) *todo.BaseResult {
	err := validations.Validate(input, v.schemas)
	if err != nil {
		return err
	}

	return v.next.ChangePassword(input, ctx)
}
