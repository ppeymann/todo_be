package account

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	todo "github.com/ppeymann/todo_be.git"
	"github.com/ppeymann/todo_be.git/auth"
	"github.com/ppeymann/todo_be.git/env"
	"github.com/ppeymann/todo_be.git/models"
	"github.com/ppeymann/todo_be.git/utils"
	"github.com/segmentio/ksuid"
)

type service struct {
	repo   models.AccountRepository
	config *todo.Configuration
}

func NewService(repo models.AccountRepository, config *todo.Configuration) models.AccountService {
	return &service{
		repo:   repo,
		config: config,
	}
}

// SignUp implements services.AccountService.
func (s *service) SignUp(input *models.SignUpInput, ctx *gin.Context) *todo.BaseResult {
	// first hashing password if is production mode
	if env.IsProduction() {
		hash, err := utils.HashString(input.Password)
		if err != nil {
			return &todo.BaseResult{
				Status: http.StatusOK,
				Errors: []string{err.Error()},
			}
		}

		input.Password = hash
	}

	// create account
	account, err := s.repo.Create(input)
	if err != nil {
		return &todo.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	// add refresh to refresh table
	refresh := models.RefreshTokenEntity{
		TokenId:   ksuid.New().String(),
		UserAgent: ctx.Request.UserAgent(),
		IssuedAt:  time.Now().UTC().Unix(),
		ExpiredAt: time.Now().Add(time.Duration(s.config.Jwt.RefreshExpire) * time.Minute).UTC().Unix(),
	}

	account.Tokens = append(account.Tokens, refresh)

	err = s.repo.Update(account)
	if err != nil {
		return &todo.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	// create token and refresh token
	paseto, err := auth.NewPasetoMaker(env.GetStringDefault("JWT", ""))
	if err != nil {
		return &todo.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	tokenClaims := &auth.Claims{
		Subject:   account.ID,
		Issuer:    s.config.Jwt.Issuer,
		Audience:  s.config.Jwt.Audience,
		IssuedAt:  time.Unix(refresh.IssuedAt, 0),
		ExpiredAt: time.Now().Add(time.Duration(s.config.Jwt.TokenExpire) * time.Minute).UTC(),
	}

	refreshClaims := &auth.Claims{
		Subject:   account.ID,
		ID:        refresh.TokenId,
		Issuer:    s.config.Jwt.Issuer,
		Audience:  s.config.Jwt.Audience,
		IssuedAt:  time.Unix(refresh.IssuedAt, 0),
		ExpiredAt: time.Unix(refresh.ExpiredAt, 0),
	}

	// create token string with paseto maker
	tokenStr, err := paseto.CreateToken(tokenClaims)
	if err != nil {
		return &todo.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	refreshStr, err := paseto.CreateToken(refreshClaims)
	if err != nil {
		return &todo.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	return &todo.BaseResult{
		Status: http.StatusOK,
		Result: models.TokenBundleOutput{
			Token:   tokenStr,
			Refresh: refreshStr,
			Expire:  tokenClaims.ExpiredAt,
		},
	}
}

func (s *service) SignIn(in *models.LoginInput, ctx *gin.Context) *todo.BaseResult {
	acc, err := s.repo.FindByUserName(in.Username)
	if err != nil {
		return &todo.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	if env.IsProduction() {
		if ok := utils.CheckHashedString(in.Password, acc.Password); !ok {
			return &todo.BaseResult{
				Status: http.StatusOK,
				Errors: []string{models.ErrSignInFailed.Error()},
			}
		}
	} else {
		if in.Password != acc.Password {
			return &todo.BaseResult{
				Status: http.StatusOK,
				Errors: []string{models.ErrSignInFailed.Error()},
			}
		}
	}

	// add refresh to refresh table
	refresh := models.RefreshTokenEntity{
		TokenId:   ksuid.New().String(),
		UserAgent: ctx.Request.UserAgent(),
		IssuedAt:  time.Now().UTC().Unix(),
		ExpiredAt: time.Now().Add(time.Duration(s.config.Jwt.RefreshExpire) * time.Minute).UTC().Unix(),
	}

	acc.Tokens = append(acc.Tokens, refresh)

	err = s.repo.Update(acc)
	if err != nil {
		return &todo.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	// create token and refresh token
	paseto, err := auth.NewPasetoMaker(env.GetStringDefault("JWT", ""))
	if err != nil {
		return &todo.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	tokenClaims := &auth.Claims{
		Subject:   acc.ID,
		Issuer:    s.config.Jwt.Issuer,
		Audience:  s.config.Jwt.Audience,
		IssuedAt:  time.Unix(refresh.IssuedAt, 0),
		ExpiredAt: time.Now().Add(time.Duration(s.config.Jwt.TokenExpire) * time.Minute).UTC(),
	}

	refreshClaims := &auth.Claims{
		Subject:   acc.ID,
		ID:        refresh.TokenId,
		Issuer:    s.config.Jwt.Issuer,
		Audience:  s.config.Jwt.Audience,
		IssuedAt:  time.Unix(refresh.IssuedAt, 0),
		ExpiredAt: time.Unix(refresh.ExpiredAt, 0),
	}

	// create token string with paseto maker
	tokenStr, err := paseto.CreateToken(tokenClaims)
	if err != nil {
		return &todo.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	refreshStr, err := paseto.CreateToken(refreshClaims)
	if err != nil {
		return &todo.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	return &todo.BaseResult{
		Status: http.StatusOK,
		Result: models.TokenBundleOutput{
			Token:   tokenStr,
			Refresh: refreshStr,
			Expire:  tokenClaims.ExpiredAt,
		},
	}
}

func (s *service) ChangePassword(in *models.ChangePasswordInput, _ *gin.Context) *todo.BaseResult {
	acc, err := s.repo.FindByID(in.Subject)
	if err != nil {
		return &todo.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	if env.IsProduction() {
		if ok := utils.CheckHashedString(in.Old, acc.Password); !ok {
			return &todo.BaseResult{
				Status: http.StatusOK,
				Errors: []string{"Password is not correct"},
			}
		}

		hash, err := utils.HashString(in.New)
		if err != nil {
			return &todo.BaseResult{
				Status: http.StatusOK,
				Errors: []string{err.Error()},
			}
		}

		acc.Password = hash
	} else {
		if acc.Password != in.Old {
			return &todo.BaseResult{
				Status: http.StatusOK,
				Errors: []string{"Password is not correct"},
			}
		}

		acc.Password = in.New
	}

	err = s.repo.Update(acc)
	if err != nil {
		return &todo.BaseResult{
			Status: http.StatusOK,
			Errors: []string{err.Error()},
		}
	}

	return &todo.BaseResult{
		Status: http.StatusOK,
		Result: "Successful",
	}
}
