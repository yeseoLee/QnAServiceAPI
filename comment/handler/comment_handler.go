package comment

import (
	"net/http"
	"qna/domain"

	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	CUseCase domain.CommentUseCase
}

type Req struct{}
type Res struct{}

func NewCommentHandler(e *echo.Echo, us domain.CommentUseCase) {
	handler := CommentHandler{
		CUseCase: us,
	}
	e_comment := e.Group("/comments")
	{
		e_comment.GET("", handler.GetComments)
		e_comment.POST("/:id", handler.AddComment)
		e_comment.DELETE("/:id", handler.DeleteComment)
	}
}

func (h *CommentHandler) GetComments(c echo.Context) error {
	return c.String(http.StatusOK, "GetComments")
}

func (h *CommentHandler) AddComment(c echo.Context) error {
	return c.String(http.StatusOK, "AddComment")
}

func (h *CommentHandler) DeleteComment(c echo.Context) error {
	return c.String(http.StatusOK, "DeleteComment")
}
