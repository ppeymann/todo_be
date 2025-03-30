package repository

import (
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/ppeymann/todo_be.git/models"
	"gorm.io/gorm"
)

type accountRepository struct {
	pg       *gorm.DB
	database string
	table    string
}

func NewAccountRepository(db *gorm.DB, database string) models.AccountRepository {
	return &accountRepository{
		pg:       db,
		database: database,
		table:    "account_entities",
	}
}

func (r *accountRepository) Create(in *models.SignUpInput) (*models.AccountEntity, error) {

	_, err := r.FindByUserName(in.Username)
	if err == nil {
		return nil, models.ErrAccountExist
	}

	account := &models.AccountEntity{
		Model:     gorm.Model{},
		Username:  in.Username,
		Password:  in.Password,
		LastName:  in.LastName,
		FirstName: in.FirstName,
	}

	// Create account
	err = r.pg.Transaction(func(tx *gorm.DB) error {
		if res := r.Model().Create(account).Error; res != nil {
			str := res.(*pgconn.PgError).Message
			if strings.Contains(str, "duplicate key value") {
				return models.ErrAccountExist
			}
			return res
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (r *accountRepository) FindByUserName(username string) (*models.AccountEntity, error) {
	acc := &models.AccountEntity{}

	err := r.Model().Where("user_name = ?", username).First(acc).Error
	if err != nil {
		return nil, err
	}

	return acc, err
}

// Update implements models.AccountRepository.
func (r *accountRepository) Update(account *models.AccountEntity) error {
	return r.pg.Save(&account).Error
}

// Migrate implements models.AccountRepository.
func (r *accountRepository) Migrate() error {
	err := r.pg.AutoMigrate(&models.RefreshTokenEntity{})
	if err != nil {
		return err
	}

	return r.pg.AutoMigrate(&models.AccountEntity{})
}

// Model implements models.AccountRepository.
func (r *accountRepository) Model() *gorm.DB {
	return r.pg.Model(&models.AccountEntity{})
}

// Name implements models.AccountRepository.
func (r *accountRepository) Name() string {
	return r.table
}
