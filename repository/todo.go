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

func (r *todoRepository) GetByID(id, accountID uint) (*models.TodoEntity, error) {
	t := &models.TodoEntity{}

	err := r.Model().Where("id = ? AND account_id = ?", id, accountID).First(t).Error
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (r *todoRepository) Update(todo *models.TodoEntity) error {
	return r.Model().Save(todo).Error
}

func (r *todoRepository) DeleteTodo(id, accountID uint) error {
	_, err := r.GetByID(id, accountID)
	if err != nil {
		return err
	}

	return r.Model().Where("id = ? AND account_id = ?", id, accountID).Delete(&models.TodoEntity{}).Error
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
