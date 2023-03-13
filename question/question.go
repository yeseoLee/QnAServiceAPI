package question

import (
	"log"
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
		log.Fatal(err)
	}
	qu := usecase.NewQuestionUseCase(qr)
	handler.NewQuestionHandler(e, qu)

}
