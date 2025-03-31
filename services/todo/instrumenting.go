package todos

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/metrics"
	todo "github.com/ppeymann/todo_be.git"
	"github.com/ppeymann/todo_be.git/models"
)

type instrumentingServices struct {
	requestCounter metrics.Counter
	requestLatency metrics.Histogram
	next           models.TodoService
}

func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, service models.TodoService) models.TodoService {
	return &instrumentingServices{
		requestCounter: counter,
		requestLatency: latency,
		next:           service,
	}
}

func (i instrumentingServices) AddTodo(in *models.TodoInput, ctx *gin.Context) *todo.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "AddTodo").Add(1)
		i.requestLatency.With("method", "AddTodo").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.AddTodo(in, ctx)

}

func (i instrumentingServices) GetAll(ctx *gin.Context) *todo.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "GetAll").Add(1)
		i.requestLatency.With("method", "GetAll").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.GetAll(ctx)

}
