package answer

import (
	handler "qna/answer/handler"
	repository "qna/answer/repository"
	usecase "qna/answer/usecase"
	"qna/datasource"

	"github.com/labstack/echo/v4"
)

func RegistAnswerRoute(ds datasource.DataSource, e *echo.Echo) {
	// Answer
	ar, err := repository.NewAnswerRepository(ds)
	if err != nil {
		panic(err)
	}
	au := usecase.NewAnswerUseCase(ar)
	handler.NewAnswerHandler(e, au)
}
