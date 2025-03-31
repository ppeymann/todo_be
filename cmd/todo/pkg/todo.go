package pkg

import (
	"log"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	kitLog "github.com/go-kit/log"
	todo "github.com/ppeymann/todo_be.git"
	"github.com/ppeymann/todo_be.git/models"
	"github.com/ppeymann/todo_be.git/repository"
	"github.com/ppeymann/todo_be.git/server"
	todos "github.com/ppeymann/todo_be.git/services/todo"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"gorm.io/gorm"
)

func InitTodoService(db *gorm.DB, logger kitLog.Logger, config *todo.Configuration, server *server.Server) models.TodoService {
	todoRepo := repository.NewTodoRepository(db, config.Database)

	err := todoRepo.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	// @TodoService create service
	todoService := todos.NewService(todoRepo)

	// getting path
	path := getSchemaPath("todo")
	todoService, err = todos.NewValidationService(path, todoService)
	if err != nil {
		log.Fatal(err)
	}

	// @Inject logging service to chain
	todoService = todos.NewLoggingServices(kitLog.With(logger, "component", "todo"), todoService)

	// @Inject instrumenting service to chain
	todoService = todos.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "todo",
			Name:      "request_count",
			Help:      "num of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "todo",
			Name:      "request_latency_microseconds",
			Help:      "total duration of requests (ms).",
		}, fieldKeys),
		todoService,
	)

	// @Inject authorization service to chain and return it
	todoService = todos.NewAuthorizationService(todoService)

	_ = todos.NewHandler(todoService, server)

	return todoService
}
