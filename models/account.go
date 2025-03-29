package models

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	todo "github.com/ppeymann/todo_be.git"
	"gorm.io/gorm"
)

// Errors
var ErrAccountExist = errors.New("account with specified params already exist")
var ErrSignInFailed = errors.New("account not found or password error")
var ErrPermissionDenied = errors.New("specified role is not available for user")
var ErrAccountNotExist = errors.New("specified account does not exist")

type (
	// AccountService represents method signatures for api account endpoint.
	// so any object that stratifying this interface can be used as account service for api endpoint.
	AccountService interface {
		// SignUp is service for sign up and create account
		SignUp(in *SignUpInput, ctx *gin.Context) *todo.BaseResult
	}

	// AccountRepository represents method signatures for account domain repository.
	// so any object that stratifying this interface can be used as account domain repository.
	AccountRepository interface {
		// Create a account
		Create(in *SignUpInput) (*AccountEntity, error)

		// FindByUserName
		FindByUserName(username string) (*AccountEntity, error)

		// Update is for updating account
		Update(account *AccountEntity) error
	}

	// AccountHandler represents method signatures for account handlers.
	// so any object that stratifying this interface can be used as account handlers.
	AccountHandler interface {
		// SignUp is handler for sign up and create account
		SignUp(ctx *gin.Context)
	}

	// AccountEntity Contains account info and is entity of user account that stored on database.
	//
	// swagger:model AccountEntity
	AccountEntity struct {
		gorm.Model `swaggerignore:"true"`

		// Username
		Username string `json:"user_name" gorm:"column:user_name;index;unique"`

		// Password
		Password string `json:"password" gorm:"password" mapstructure:"password"`

		// LastName
		LastName string `json:"last_name" gorm:"column:last_name"`

		// FirstName
		FirstName string `json:"first_name" gorm:"column:first_name"`

		// Tokens list of current account active session
		Tokens []RefreshTokenEntity `json:"-" gorm:"foreignKey:AccountID;references:ID"`
	}

	// RefreshTokenEntity is entity to store accounts active session
	RefreshTokenEntity struct {
		gorm.Model
		AccountID uint
		TokenId   string `json:"token_id" gorm:"column:token_id;index"`
		UserAgent string `json:"user_agent" gorm:"column:user_agent"`
		IssuedAt  int64  `json:"issued_at" gorm:"column:issued_at"`
		ExpiredAt int64  `json:"expired_at" gorm:"column:expired_at"`
	}

	// LoginInput is request model for login endpoint
	//
	// swagger:model LoginRequest
	LoginInput struct {
		Username string `json:"user_name" gorm:"column:user_name;index;unique"`
		Password string `json:"password" gorm:"password" mapstructure:"password"`
	}

	// SignUpInput is request model for sign up endpoint
	//
	// swagger:model SignUpInput
	SignUpInput struct {
		Username  string `json:"user_name" gorm:"column:user_name;index;unique"`
		Password  string `json:"password" gorm:"password" mapstructure:"password"`
		LastName  string `json:"last_name" gorm:"column:last_name"`
		FirstName string `json:"first_name" gorm:"column:first_name"`
	}

	// TokenBundleOutput Contains Token, Refresh Token, Date and Token Expire time for Login/Verify response DTO.
	//
	// swagger:model TokenBundleOutput
	TokenBundleOutput struct {
		// Token is JWT/PASETO token staring for storing in client side as access token
		Token string `json:"token"`

		// Refresh token string used for refreshing authentication and give fresh token
		Refresh string `json:"refresh"`

		// Expire time of Token and CentrifugeToken
		Expire time.Time `json:"expire"`
	}
)
