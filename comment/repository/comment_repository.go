package comment

import (
	"database/sql"
	"fmt"
	"qna/datasource"
	model "qna/domain"
	"qna/util"
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

func (r *CommentRepository) FindAllByPostId(id uint, limit int, offset int) ([]*model.Comment, error) {

	var cList []*model.Comment

	// Query
	rows, err := r.DBEngine.Query("SELECT id, questionId, AnswerId, content, writerId, createdAt FROM tbComment WHERE id=? LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Read
	for rows.Next() {
		var c model.Comment
		rows.Scan(&c.Id, &c.QuestionId, &c.AnswerId, &c.Content, &c.WriterId, &c.CreatedAt)
		cList = append(cList, &c)
	}

	return cList, nil
}

func (r *CommentRepository) Create(commentInput *model.CommentInput) (*model.Comment, error) {
	var comment *model.Comment
	now := util.DateTimeNow()

	// Query
	result, err := r.DBEngine.Exec("INSERT INTO tbComment (`questionId`, `answerId`, `Content`, `WriterId`, `CreatedAt`) VALUES (?, ?, ?, ?, ?)",
		commentInput.QuestionId, commentInput.AnswerId, commentInput.Content, commentInput.WriterId, now)
	if err != nil {
		return comment, err
	}

	// Check
	id, err := result.LastInsertId()
	comment.Id = uint(id)
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (r *CommentRepository) Delete(id uint) error {
	// Query
	result, err := r.DBEngine.Exec("DELETE FROM `tbQuestion` WHERE id = ?", id)
	if err != nil {
		return err
	}

	// Check
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return fmt.Errorf("expected to affect 1 row, affected %d", rows)
	}
	return nil
}
