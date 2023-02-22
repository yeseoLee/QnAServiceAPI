package comment

import (
	handler "qna/comment/handler"
	repository "qna/comment/repository"
	usecase "qna/comment/usecase"
	"qna/datasource"

	"github.com/labstack/echo/v4"
)

func RegistCommentRoute(ds datasource.DataSource, e *echo.Echo) {
	// comment
	cr, err := repository.NewCommentRepository(ds)
	if err != nil {
		panic(err)
	}
	cu := usecase.NewCommentUseCase(cr)
	handler.NewCommentHandler(e, cu)
}
