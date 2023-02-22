package comment

import (
	"database/sql"
	"qna/datasource"
	"qna/domain"
)

func NewCommentRepository(ds datasource.DataSource) (*CommentRepository, error) {
	db, err := ds.GetConnection()
	if err != nil {
		return nil, err
	}
	return &CommentRepository{DBEngine: db}, nil
}

type CommentRepository struct {
	DBEngine *sql.DB
}

func (r *CommentRepository) FindAllByPostId(id uint, limit int, offset int) ([]*domain.Comment, error) {
	return nil, nil
}

func (r *CommentRepository) Create(commentInput *domain.CommentInput) (*domain.Comment, error) {
	return nil, nil
}
func (r *CommentRepository) Delete(id uint) error {
	return nil
}
