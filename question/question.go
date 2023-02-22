package question

import (
	"qna/datasource"
	handler "qna/question/handler"
	repository "qna/question/repository"
	usecase "qna/question/usecase"

	"github.com/labstack/echo/v4"
)

func RegistQuestionRoute(ds datasource.DataSource, e *echo.Echo) {
	// Question
	qr, err := repository.NewQuestionRepository(ds)
	if err != nil {
		panic(err)
	}
	qu := usecase.NewQuestionUseCase(qr)
	handler.NewQuestionHandler(e, qu)

}
