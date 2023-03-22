package comment

import (
	"qna/domain"
	"qna/util"
)

type commentUsecase struct {
	commentRepo domain.CommentRepository
}

func NewCommentUseCase(u domain.CommentRepository) domain.CommentUseCase {
	return &commentUsecase{
		commentRepo: u,
	}
}

func (u *commentUsecase) GetAll(limit, offset int) ([]*domain.CommentOutput, error) {
	return nil, nil
}
func (u *commentUsecase) Create(commentInput *domain.CommentInput) (*domain.CommentOutput, error) {
	return nil, nil
}
func (u *commentUsecase) Delete(id uint64) error {
	return nil
}

func (u *commentUsecase) transferDAO(comment *domain.CommentInput) *domain.CommentDAO {
	dao := &domain.CommentDAO{}
	dao.QuestionId = comment.QuestionId
	dao.AnswerId = comment.AnswerId
	dao.WriterId = comment.WriterId
	dao.Content = comment.Content
	return dao
}

func (u *commentUsecase) transferOutput(dao *domain.CommentDAO) *domain.CommentOutput {
	co := &domain.CommentOutput{}
	co.Id = dao.Id
	co.QuestionId = dao.QuestionId
	co.AnswerId = dao.AnswerId
	co.WriterId = dao.WriterId
	co.Content = dao.Content
	co.CreatedAt = util.DateTimeStringToTime(dao.CreatedAt)
	return co
}
