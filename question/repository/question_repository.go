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

func (r *QuestionRepository) FindById(id uint64) (*model.Question, error) {
	var q model.Question
	var isAccept uint8
	var tags_str string
	var images_str string

	// Query
	row := r.DBEngine.QueryRow("SELECT id, title, content, writerId, tags, images, isAccept, createdAt, updatedAt FROM tbQuestion WHERE id = ?", id)

	// Read
	err := row.Scan(&q.Id, &q.Title, &q.Content, &q.WriterId, tags_str, images_str, isAccept, q.CreatedAt, q.UpdatedAt)
	switch {
	case err == sql.ErrNoRows:
		return nil, errors.New("no data")
	case err != nil:
		return nil, errors.New("query error")
	}

	// type casting
	q.Tags = strings.Split(tags_str, ",")
	q.Images = strings.Split(images_str, ",")
	if isAccept == 1 {
		q.IsAccept = true
	} else {
		q.IsAccept = false
	}
	return &q, nil
}

func (r *QuestionRepository) FindAll(limit, offset int) ([]*model.Question, error) {
	var qList []*model.Question
	var tags_str string
	var images_str string

	// Query
	rows, err := r.DBEngine.Query("SELECT id, title, content, writerId, tags, images, createdAt, updatedAt FROM tbQuestion LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Read
	for rows.Next() {
		var q model.Question
		rows.Scan(&q.Id, &q.Title, &q.Content, &q.WriterId, tags_str, images_str, q.CreatedAt, q.UpdatedAt)
		q.Tags = strings.Split(tags_str, ",")
		q.Images = strings.Split(images_str, ",")
		qList = append(qList, &q)
	}

	return qList, nil
}

func (r *QuestionRepository) Create(questionInput *model.QuestionInput) (uint64, error) {
	now := util.DateTimeNow()

	// type casting
	tags_str := strings.Join(questionInput.Tags, ",")
	images_str := strings.Join(questionInput.Images, ",")

	// Query
	result, err := r.DBEngine.Exec("INSERT INTO tbQuestion (`Title`,`Content`,`WriterId`, `Tags`, `Images`,`CreatedAt`,`UpdatedAt`) VALUES (?, ?, ?, ?, ?, ?, ?)",
		questionInput.Title, questionInput.Content, questionInput.WriterId, tags_str, images_str, now, now)
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

func (r *QuestionRepository) Update(id uint64, questionUpdate map[string]interface{}) (*model.Question, error) {
	// TODO: map key-value check & make query logic
	var question *model.Question
	now := util.DateTimeNow()

	// type casting
	tags, ok := questionUpdate["Tags"].([]string)
	if !ok {
		return question, fmt.Errorf("unexpected parameter, wants: Tags []string")
	}
	images, ok := questionUpdate["Images"].([]string)
	if !ok {
		return question, fmt.Errorf("unexpected parameter, wants: Images []string")
	}
	tags_str := strings.Join(tags, ",")
	images_str := strings.Join(images, ",")

	fmt.Printf("%+v\n", questionUpdate)
	fmt.Printf("%s, %s\n", images_str, tags_str)

	// Query
	result, err := r.DBEngine.Exec("UPDATE tbQuestion SET Title = ?, Content = ?, Tags =?, Images = ?, updatedAt = ? WHERE id = ?",
		questionUpdate["Title"], questionUpdate["Content"], images_str, tags_str, now, id)
	if err != nil {
		return question, err
	}

	// Check
	rows, err := result.RowsAffected()
	if err != nil {
		return question, err
	}
	if rows != 1 {
		return question, fmt.Errorf("expected to affect 1 row, affected %d", rows)
	}
	return question, nil
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
