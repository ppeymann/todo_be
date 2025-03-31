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
	}

	// @TodoRepository represents method signatures for todo domain repository.
	// so any object that stratifying this interface can be used as todo domain repository.
	TodoRepository interface {
		// CreateTodo is method for creating todo
		CreateTodo(in *TodoInput, id uint) (*TodoEntity, error)

		todo.BaseRepository
	}

	// @TodoHandler represents method signatures for Todo handlers.
	// so any object that stratifying this interface can be used as Todo handlers.
	TodoHandler interface {
		AddTodo(ctx *gin.Context)
	}

	// @TodoEntity represents todo entity
	//
	// swagger:model TodoEntity
	TodoEntity struct {
		gorm.Model `swaggerignore:"true"`

		// Title
		Title string `json:"title" gorm:"column:title;index;unique"`

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
		Title       string `json:"title" gorm:"column:title;index;unique"`
		Description string `json:"description" gorm:"column:description"`
		Priority    uint32 `json:"priority" gorm:"column:priority"`
	}
)

const (
	InProgress string = "in_progress"
	Done       string = "complete"
	Cancel     string = "cancel"
)
