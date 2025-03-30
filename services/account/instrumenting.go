package account

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
	next           models.AccountService
}

func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, service models.AccountService) models.AccountService {
	return &instrumentingServices{
		requestCounter: counter,
		requestLatency: latency,
		next:           service,
	}
}

// SignUp implements services.AccountServices.
func (i *instrumentingServices) SignUp(input *models.SignUpInput, ctx *gin.Context) *todo.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "SignUp").Add(1)
		i.requestLatency.With("method", "SignUp").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.SignUp(input, ctx)
}

// SignIn implements services.AccountServices.
func (i *instrumentingServices) SignIn(input *models.LoginInput, ctx *gin.Context) *todo.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "SignIn").Add(1)
		i.requestLatency.With("method", "SignIn").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.SignIn(input, ctx)
}

// ChangePassword implements services.AccountServices.
func (i *instrumentingServices) ChangePassword(in *models.ChangePasswordInput, ctx *gin.Context) *todo.BaseResult {
	defer func(begin time.Time) {
		i.requestCounter.With("method", "ChangePassword").Add(1)
		i.requestLatency.With("method", "ChangePassword").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.ChangePassword(in, ctx)
}
