package question

import (
	"log"
	"net/http"
	"qna/domain"
	"strconv"

	"github.com/labstack/echo/v4"
)

type QuestionHandler struct {
	QUseCase domain.QuestionUseCase
}

// TODO: 객체 계층 분리
type Req struct{}
type Res struct{}

func NewQuestionHandler(e *echo.Echo, us domain.QuestionUseCase) {
	handler := QuestionHandler{
		QUseCase: us,
	}
	e_question := e.Group("/questions")
	{
		e_question.GET("", handler.GetQuestions)
		e_question.GET("/:id", handler.GetQuestion)
		e_question.POST("", handler.AddQuestion)
		e_question.PATCH("/:id", handler.EditQuestion)
		e_question.DELETE("/:id", handler.DeleteQuestion)
	}
}

func (h *QuestionHandler) GetQuestions(c echo.Context) error {
	var req *domain.QuestionSearchOption
	var res []*domain.QuestionOutput

	err := c.Bind(&req)
	if err != nil {
		log.Print(err)
		return c.String(http.StatusBadRequest, "bad request")
	}

	res, err = h.QUseCase.GetAll(req)
	if err != nil {
		log.Print(err)
		return c.String(http.StatusInternalServerError, "InternalServerError")
	}

	return c.JSON(http.StatusOK, res)
}

func (h *QuestionHandler) GetQuestion(c echo.Context) error {
	var res *domain.QuestionOutput

	idString := c.FormValue("id")
	idUint, _ := strconv.ParseUint(idString, 10, 16)

	res, err := h.QUseCase.Get(idUint)
	if err != nil {
		log.Print(err)
		return c.String(http.StatusInternalServerError, "InternalServerError")
	}

	return c.JSON(http.StatusOK, res)
}

func (h *QuestionHandler) AddQuestion(c echo.Context) error {
	var req *domain.QuestionInput
	var res *domain.QuestionOutput

	err := c.Bind(&req)
	if err != nil {
		log.Print(err)
		return c.String(http.StatusBadRequest, "bad request")
	}

	res, err = h.QUseCase.Create(req)
	if err != nil {
		log.Print(err)
		return c.String(http.StatusInternalServerError, "InternalServerError")
	}
	return c.JSON(http.StatusOK, res)
}

func (h *QuestionHandler) EditQuestion(c echo.Context) error {
	var req *domain.QuestionInput
	var res *domain.QuestionOutput

	idString := c.FormValue("id")
	idUint, _ := strconv.ParseUint(idString, 10, 16)

	err := c.Bind(&req)
	if err != nil {
		log.Print(err)
		return c.String(http.StatusBadRequest, "bad request")
	}

	res, err = h.QUseCase.Edit(idUint, map[string]interface{}{})
	if err != nil {
		log.Print(err)
		return c.String(http.StatusInternalServerError, "InternalServerError")
	}
	return c.JSON(http.StatusOK, res)
}

func (h *QuestionHandler) DeleteQuestion(c echo.Context) error {
	idString := c.FormValue("id")
	idUint, _ := strconv.ParseUint(idString, 10, 16)

	err := h.QUseCase.Delete(idUint)
	if err != nil {
		log.Print(err)
		return c.String(http.StatusInternalServerError, "InternalServerError")
	}

	return c.String(http.StatusOK, "deletequestion")
}
