package todo

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
