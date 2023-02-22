package answer

import (
	"net/http"
	"qna/domain"

	"github.com/labstack/echo/v4"
)

type AnswerHandler struct {
	AUseCase domain.AnswerUseCase
}

type Req struct{}
type Res struct{}

func NewAnswerHandler(e *echo.Echo, us domain.AnswerUseCase) {
	handler := AnswerHandler{
		AUseCase: us,
	}
	e_answer := e.Group("/answers")
	{
		e_answer.GET("", handler.GetAnswers)
		e_answer.POST("/:id", handler.AddAnswer)
		e_answer.PATCH("/:id", handler.EditAnswer)
		e_answer.DELETE("/:id", handler.DeleteAnswer)
	}
}

func (h *AnswerHandler) GetAnswers(c echo.Context) error {
	return c.String(http.StatusOK, "GetAnswers")
}
func (h *AnswerHandler) AddAnswer(c echo.Context) error {
	return c.String(http.StatusOK, "AddAnswer")
}
func (h *AnswerHandler) EditAnswer(c echo.Context) error {
	return c.String(http.StatusOK, "EditAnswer")
}
func (h *AnswerHandler) DeleteAnswer(c echo.Context) error {
	return c.String(http.StatusOK, "DeleteAnswer")
}
