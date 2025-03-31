package models

import (
	"github.com/gin-gonic/gin"
	todo "github.com/ppeymann/todo_be.git"
	"gorm.io/gorm"
)

type (
	// @TodoService represents method signatures for api todo endpoint.
	// so any object that stratifying this interface can be used as todo service for api endpoint.
	TodoService interface {
		// AddTodo is method for adding todo
		AddTodo(in *TodoInput, ctx *gin.Context) *todo.BaseResult

		// GetAll is method for getting all todo with specific account ID
		GetAll(ctx *gin.Context) *todo.BaseResult

		// GetByID
		GetByID(id uint, ctx *gin.Context) *todo.BaseResult

		// DeleteTodo
		DeleteTodo(id uint, ctx *gin.Context) *todo.BaseResult
	}

	// @TodoRepository represents method signatures for todo domain repository.
	// so any object that stratifying this interface can be used as todo domain repository.
	TodoRepository interface {
		// CreateTodo is method for creating todo
		CreateTodo(in *TodoInput, id uint) (*TodoEntity, error)

		// GetAll is method for getting all todo with specific account ID
		GetAll(id uint) ([]TodoEntity, error)

		// GetByID is method for get one todo by that ID
		GetByID(id, accountID uint) (*TodoEntity, error)

		// Update is for update todo by ID
		Update(*TodoEntity) error

		// DeleteTodo is for delete a todo by that ID
		DeleteTodo(id, accountID uint) error

		todo.BaseRepository
	}

	// @TodoHandler represents method signatures for Todo handlers.
	// so any object that stratifying this interface can be used as Todo handlers.
	TodoHandler interface {
		// AddTodo is handler for add new todo task http request.
		AddTodo(ctx *gin.Context)

		// GetAll is handler for get all todo task http request.
		GetAll(ctx *gin.Context)

		// GetByID is handler for get one todo by specific ID
		GetByID(ctx *gin.Context)

		// DeleteTodo is handler for delete one todo by specific ID
		DeleteTodo(*gin.Context)
	}

	// @TodoEntity represents todo entity
	//
	// swagger:model TodoEntity
	TodoEntity struct {
		gorm.Model `swaggerignore:"true"`

		// Title
		Title string `json:"title" gorm:"column:title"`

		// Description
		Description string `json:"description" gorm:"column:description"`

		// Status [in_progress, complete, ]
		Status string `json:"status" gorm:"column:status"`

		// AccountID
		AccountID uint `json:"account_id" gorm:"column:account_id"`

		// Priority	[1 = not important, 2 = important, 3 = very important]
		Priority uint32 `json:"priority" gorm:"column:priority"`
	}

	// @TodoInput represents todo input DTO
	//
	// swagger:model TodoInput
	TodoInput struct {
		Title       string `json:"title" gorm:"column:title"`
		Description string `json:"description" gorm:"column:description"`
		Priority    uint32 `json:"priority" gorm:"column:priority"`
	}
)

const (
	InProgress string = "in_progress"
	Done       string = "complete"
	Cancel     string = "cancel"
)
