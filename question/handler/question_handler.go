package question

import (
	"net/http"
	"qna/domain"

	"github.com/labstack/echo/v4"
)

type QuestionHandler struct {
	QUseCase domain.QuestionUseCase
}

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
		e_question.POST("/:id", handler.AddQuestion)
		e_question.PATCH("/:id", handler.EditQuestion)
		e_question.DELETE("/:id", handler.DeleteQuestion)
	}
}

func (h *QuestionHandler) GetQuestions(c echo.Context) error {
	var req Req
	var res Res

	err := c.Bind(&req)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	return c.JSON(http.StatusOK, res)
}

func (h *QuestionHandler) GetQuestion(c echo.Context) error {
	/*
		func (a *ArticleHandler) GetArticle(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			id, _ := strconv.ParseInt(vars["id"], 10, 16)
			article, _ := a.AUseCase.GetByID(id)
			response, _ := json.Marshal(article)
			fmt.Fprint(w, string(response))
		}
	*/
	return c.String(http.StatusOK, "getquestion")
}

func (h *QuestionHandler) AddQuestion(c echo.Context) error {
	return c.String(http.StatusOK, "addquestion")
}

func (h *QuestionHandler) EditQuestion(c echo.Context) error {
	return c.String(http.StatusOK, "editquestion")
}

func (h *QuestionHandler) DeleteQuestion(c echo.Context) error {
	return c.String(http.StatusOK, "deletequestion")
}
