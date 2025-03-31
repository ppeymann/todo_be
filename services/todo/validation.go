package todos

import (
	"github.com/gin-gonic/gin"
	todo "github.com/ppeymann/todo_be.git"
	"github.com/ppeymann/todo_be.git/models"
	validations "github.com/ppeymann/todo_be.git/validation"
)

type validationService struct {
	schemas map[string][]byte
	next    models.TodoService
}

func NewValidationService(schemaPath string, service models.TodoService) (models.TodoService, error) {
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

func (v *validationService) AddTodo(in *models.TodoInput, ctx *gin.Context) *todo.BaseResult {
	err := validations.Validate(in, v.schemas)
	if err != nil {
		return err
	}

	return v.next.AddTodo(in, ctx)
}
