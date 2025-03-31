package pkg

import (
	"log"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	kitLog "github.com/go-kit/log"
	todo "github.com/ppeymann/todo_be.git"
	"github.com/ppeymann/todo_be.git/models"
	"github.com/ppeymann/todo_be.git/repository"
	"github.com/ppeymann/todo_be.git/server"
	"github.com/ppeymann/todo_be.git/services/account"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"gorm.io/gorm"
)

func InitTodoService(db *gorm.DB, logger kitLog.Logger, config *todo.Configuration, server *server.Server) models.AccountService {
	accountRepo := repository.NewAccountRepository(db, config.Database)

	err := accountRepo.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	// accountService create service
	accountService := account.NewService(accountRepo, config)

	// getting path
	path := getSchemaPath("account")
	accountService, err = account.NewValidationService(path, accountService)
	if err != nil {
		log.Fatal(err)
	}

	// @Inject logging service to chain
	accountService = account.NewLoggingServices(kitLog.With(logger, "component", "account"), accountService)

	// @Inject instrumenting service to chain
	accountService = account.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "account",
			Name:      "request_count",
			Help:      "num of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "account",
			Name:      "request_latency_microseconds",
			Help:      "total duration of requests (ms).",
		}, fieldKeys),
		accountService,
	)

	// @Inject authorization service to chain and return it
	accountService = account.NewAuthorizationService(accountService)

	_ = account.NewHandler(accountService, server)

	return accountService
}
