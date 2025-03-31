package repository

import (
	"errors"

	"github.com/ppeymann/todo_be.git/models"
	"gorm.io/gorm"
)

type todoRepository struct {
	pg       *gorm.DB
	database string
	table    string
}

func NewTodoRepository(db *gorm.DB, database string) models.TodoRepository {
	return &todoRepository{
		pg:       db,
		database: database,
		table:    "todo_entities",
	}
}

func (r *todoRepository) CreateTodo(in *models.TodoInput, id uint) (*models.TodoEntity, error) {
	accountRepo := NewAccountRepository(r.pg, r.database)

	acc, err := accountRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("account not found")
	}

	todo := &models.TodoEntity{
		Model:       gorm.Model{},
		Title:       in.Title,
		Description: in.Description,
		Status:      models.InProgress,
		AccountID:   acc.ID,
		Priority:    in.Priority,
	}

	err = r.Model().Create(todo).Error
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (r *todoRepository) GetAll(id uint) ([]models.TodoEntity, error) {
	var todos []models.TodoEntity

	err := r.Model().Where("account_id = ?", id).Find(&todos).Error
	if err != nil {
		return nil, err
	}

	return todos, nil
}

// Migrate implements models.todoRepository.
func (r *todoRepository) Migrate() error {

	return r.pg.AutoMigrate(&models.TodoEntity{})
}

// Model implements models.todoRepository.
func (r *todoRepository) Model() *gorm.DB {
	return r.pg.Model(&models.TodoEntity{})
}

// Name implements models.todoRepository.
func (r *todoRepository) Name() string {
	return r.table
}
