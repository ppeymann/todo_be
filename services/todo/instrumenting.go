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

func (i instrumentingServices) GetByID(id uint, ctx *gin.Context) *todo.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "GetByID").Add(1)
		i.requestLatency.With("method", "GetByID").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.GetByID(id, ctx)

}

func (i instrumentingServices) DeleteTodo(id uint, ctx *gin.Context) *todo.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "DeleteTodo").Add(1)
		i.requestLatency.With("method", "DeleteTodo").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.DeleteTodo(id, ctx)

}

func (i instrumentingServices) UpdateTodo(in *models.TodoInput, id uint, ctx *gin.Context) *todo.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "UpdateTodo").Add(1)
		i.requestLatency.With("method", "UpdateTodo").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.UpdateTodo(in, id, ctx)

}

func (i instrumentingServices) ChangeStatus(status string, id uint, ctx *gin.Context) *todo.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "ChangeStatus").Add(1)
		i.requestLatency.With("method", "ChangeStatus").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.ChangeStatus(status, id, ctx)

}
