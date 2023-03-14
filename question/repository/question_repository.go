package question

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"qna/datasource"
	model "qna/domain"
	"qna/util"
)

func NewQuestionRepository(ds datasource.DataSource) (*QuestionRepository, error) {
	db, err := ds.GetConnection()
	if err != nil {
		return nil, err
	}
	return &QuestionRepository{DBEngine: db}, nil
}

type QuestionRepository struct {
	DBEngine *sql.DB
}

func (r *QuestionRepository) FindById(id uint64) (*model.QuestionDAO, error) {
	var q model.QuestionDAO
	// Query
	row := r.DBEngine.QueryRow("SELECT id, title, content, writerId, tags, images, isAccept, createdAt, updatedAt FROM tbQuestion WHERE id = ?", id)

	// Read
	err := row.Scan(&q.Id, &q.Title, &q.Content, &q.WriterId, &q.Tags, &q.Images, &q.IsAccept, &q.CreatedAt, &q.UpdatedAt)
	switch {
	case err == sql.ErrNoRows:
		return nil, errors.New("no data")
	case err != nil:
		return nil, fmt.Errorf("query error: %w", err)
	}
	return &q, nil
}

func (r *QuestionRepository) FindAll(limit, offset int) ([]*model.QuestionDAO, error) {
	var qList []*model.QuestionDAO

	// Query
	rows, err := r.DBEngine.Query("SELECT id, title, content, writerId, tags, images, isAccept, createdAt, updatedAt FROM tbQuestion LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Read
	for rows.Next() {
		var q model.QuestionDAO
		rows.Scan(&q.Id, &q.Title, &q.Content, &q.WriterId, &q.Tags, &q.Images, &q.IsAccept, &q.CreatedAt, &q.UpdatedAt)
		qList = append(qList, &q)
	}

	return qList, nil
}

func (r *QuestionRepository) Create(question *model.QuestionDAO) (uint64, error) {
	now := util.DateTimeNow()

	// Query
	result, err := r.DBEngine.Exec("INSERT INTO tbQuestion (`writerId`, `title`,`content`, `tags`, `images`,`createdAt`,`updatedAt`) VALUES (?, ?, ?, ?, ?, ?, ?)",
		question.WriterId, question.Title, question.Content, question.Tags, question.Images, now, now)
	if err != nil {
		return 0, err
	}

	// Check
	id, err := result.LastInsertId()
	if err != nil {
		return uint64(id), err
	}
	return uint64(id), nil
}

func (r *QuestionRepository) Update(id uint64, questionUpdate map[string]interface{}) error {
	// TODO: map key-value check & make query logic
	now := util.DateTimeNow()

	// type casting
	tags, ok := questionUpdate["Tags"].([]string)
	if !ok {
		return fmt.Errorf("unexpected parameter, wants: Tags []string")
	}
	images, ok := questionUpdate["Images"].([]string)
	if !ok {
		return fmt.Errorf("unexpected parameter, wants: Images []string")
	}
	tags_str := strings.Join(tags, ",")
	images_str := strings.Join(images, ",")

	// Query
	result, err := r.DBEngine.Exec("UPDATE tbQuestion SET title = ?, content = ?, tags =?, images = ?, updatedAt = ? WHERE id = ?",
		questionUpdate["Title"], questionUpdate["Content"], images_str, tags_str, now, id)
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

func (r *QuestionRepository) Delete(id uint64) error {
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
