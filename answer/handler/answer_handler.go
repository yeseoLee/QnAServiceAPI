package answer

import (
	"log"
	"net/http"
	"qna/domain"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AnswerHandler struct {
	AUseCase domain.AnswerUseCase
}

// TODO: 객체 계층 분리
type Req struct{}
type Res struct{}

func NewAnswerHandler(e *echo.Echo, us domain.AnswerUseCase) {
	handler := AnswerHandler{
		AUseCase: us,
	}
	e_answer := e.Group("/answers")
	{
		e_answer.GET("/:questionId", handler.GetAnswers)
		e_answer.POST("/:id", handler.AddAnswer)
		e_answer.PATCH("/:id", handler.EditAnswer)
		e_answer.DELETE("/:id", handler.DeleteAnswer)
	}
}

func (h *AnswerHandler) GetAnswers(c echo.Context) error {
	var req *domain.AnswerSearchOption
	var res []*domain.AnswerOutput

	err := c.Bind(&req)
	if err != nil {
		log.Print(err)
		return c.String(http.StatusBadRequest, "bad request")
	}

	res, err = h.AUseCase.GetAll(req)
	if err != nil {
		log.Print(err)
		return c.String(http.StatusInternalServerError, "InternalServerError")
	}

	return c.JSON(http.StatusOK, res)
}

func (h *AnswerHandler) AddAnswer(c echo.Context) error {
	var req *domain.AnswerInput
	var res *domain.AnswerOutput

	err := c.Bind(&req)
	if err != nil {
		log.Print(err)
		return c.String(http.StatusBadRequest, "bad request")
	}

	res, err = h.AUseCase.Create(req)
	if err != nil {
		log.Print(err)
		return c.String(http.StatusInternalServerError, "InternalServerError")
	}

	return c.JSON(http.StatusOK, res)
}

func (h *AnswerHandler) EditAnswer(c echo.Context) error {
	var req *domain.AnswerInput
	var res *domain.AnswerOutput

	idString := c.FormValue("id")
	idUint, _ := strconv.ParseUint(idString, 10, 16)

	err := c.Bind(&req)
	if err != nil {
		log.Print(err)
		return c.String(http.StatusBadRequest, "bad request")
	}

	res, err = h.AUseCase.Edit(idUint, map[string]interface{}{})
	if err != nil {
		log.Print(err)
		return c.String(http.StatusInternalServerError, "InternalServerError")
	}

	return c.JSON(http.StatusOK, res)
}
func (h *AnswerHandler) DeleteAnswer(c echo.Context) error {

	idString := c.FormValue("id")
	idUint, _ := strconv.ParseUint(idString, 10, 16)

	err := h.AUseCase.Delete(idUint)
	if err != nil {
		log.Print(err)
		return c.String(http.StatusInternalServerError, "InternalServerError")
	}

	return c.String(http.StatusOK, "DeleteAnswer")
}
