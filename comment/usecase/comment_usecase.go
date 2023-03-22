package comment

import "qna/domain"

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
