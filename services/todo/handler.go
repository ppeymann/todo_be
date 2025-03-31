package todos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	todo "github.com/ppeymann/todo_be.git"
	"github.com/ppeymann/todo_be.git/models"
	"github.com/ppeymann/todo_be.git/server"
)

type handler struct {
	service models.TodoService
}

func NewHandler(svc models.TodoService, s *server.Server) models.TodoHandler {
	handler := &handler{
		service: svc,
	}

	group := s.Router.Group("/api/v1/todo")

	// Authentication Middleware
	group.Use(s.Authenticate())
	{
		group.POST("/", handler.AddTodo)
		group.GET("/", handler.GetAll)
		group.GET("/:id", handler.GetByID)
		group.DELETE("/:id", handler.DeleteTodo)
	}

	return handler
}

// AddTodo is handler for add new todo task http request.
//
// @BasePath			/api/v1/todo
// @Summary				add new todo task
// @Description			add new todo task with specified info and Account ID
// @Tags				todos
// @Accept				json
// @Produce				json
//
// @Param				input		body		models.TodoInput	true	"todo input"
// @Success				200			{object}	todo.BaseResult{result=models.TodoEntity}	"always returns status 200 but body contains error"
// @Router				/		[post]
// @Security			Bearer Authenticate
func (h *handler) AddTodo(ctx *gin.Context) {
	in := &models.TodoInput{}

	if err := ctx.ShouldBindBodyWithJSON(in); err != nil {
		ctx.JSON(http.StatusBadRequest, todo.BaseResult{
			Errors: []string{todo.ProvideRequiredJsonBody},
		})

		return
	}

	result := h.service.AddTodo(in, ctx)
	ctx.JSON(result.Status, result)
}

// GetAll is handler for get all todos http request.
//
// @BasePath			/api/v1/todo
// @Summary				get all todo
// @Description			get all todos with specified info and Account ID
// @Tags				todos
// @Accept				json
// @Produce				json
//
// @Success				200			{object}	todo.BaseResult{result=[]models.TodoEntity}	"always returns status 200 but body contains error"
// @Router				/		[get]
// @Security			Bearer Authenticate
func (h *handler) GetAll(ctx *gin.Context) {
	result := h.service.GetAll(ctx)
	ctx.JSON(result.Status, result)
}

// GetByID is handler for get one todo http request.
//
// @BasePath			/api/v1/todo
// @Summary				get one todo
// @Description			get one todo with specified info and Account ID
// @Tags				todos
// @Accept				json
// @Produce				json
//
// @Success				200			{object}	todo.BaseResult{result=models.TodoEntity}	"always returns status 200 but body contains error"
// @Router				/{id}		[get]
// @Security			Bearer Authenticate
func (h *handler) GetByID(ctx *gin.Context) {
	id, err := server.GetPathUint64(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &todo.BaseResult{
			Errors: []string{"Params Not Supported"},
		})

		return
	}

	result := h.service.GetByID(uint(id), ctx)
	ctx.JSON(result.Status, result)
}

// DeleteTodo is handler for delete todo http request.
//
// @BasePath			/api/v1/todo
// @Summary				delete todo
// @Description			delete todo with specified info and Account ID
// @Tags				todos
// @Accept				json
// @Produce				json
//
// @Success				200			{object}	todo.BaseResult{result=string}	"always returns status 200 but body contains error"
// @Router				/{id}		[delete]
// @Security			Bearer Authenticate
func (h *handler) DeleteTodo(ctx *gin.Context) {
	id, err := server.GetPathUint64(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &todo.BaseResult{
			Errors: []string{"Params Not Supported"},
		})

		return
	}

	result := h.service.DeleteTodo(uint(id), ctx)
	ctx.JSON(result.Status, result)
}
