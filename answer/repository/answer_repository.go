package answer

import (
	"database/sql"
	"fmt"

	"qna/datasource"
	model "qna/domain"
)

func NewAnswerRepository(ds datasource.DataSource) (*AnswerRepository, error) {
	db, err := ds.GetConnection()
	if err != nil {
		return nil, err
	}
	return &AnswerRepository{DBEngine: db}, nil
}

type AnswerRepository struct {
	DBEngine *sql.DB
}

func (r *AnswerRepository) FindAllByQuestionId(id uint64, limit int, offset int) ([]*model.Answer, error) {

	var aList []*model.Answer

	// Query
	rows, err := r.DBEngine.Query("SELECT id, questionId, content, writerId, images, createdAt,updatedAt FROM tbAnswer WHERE questionId = ? LIMIT ? OFFSET ?", id, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Read
	for rows.Next() {
		var a model.Answer
		rows.Scan(&a.Id, &a.QuestionId, &a.Content, &a.WriterId, &a.Images, &a.CreatedAt, &a.UpdatedAt)
		aList = append(aList, &a)
	}

	return aList, nil
}

// func (r *AnswerRepository) FindAllByWriterId(writerId string, limit int, offset int) ([]*model.Answer, error) {}

func (r *AnswerRepository) Create(answerInput *model.AnswerInput) (*model.Answer, error) {
	var answer *model.Answer

	// Query
	result, err := r.DBEngine.Exec("INSERT INTO tbAnswer (`questionId`,`content`, `writerId`, `images`,`createdAt`) VALUES (?,?,?,?,now())",
		answerInput.QuestionId, answerInput.Content, answerInput.WriterId, answerInput.Images)
	if err != nil {
		return nil, err
	}

	// Check
	id, err := result.LastInsertId()
	answer.Id = uint64(id)
	if err != nil {
		return answer, err
	}
	return answer, nil
}

func (r *AnswerRepository) Update(id uint64, answerUpdate map[string]interface{}) (*model.Answer, error) {
	// TODO: map key-value check & make query logic
	var answer *model.Answer

	// Query
	result, err := r.DBEngine.Exec("UPDATE tbAnswer SET content = ?, images = ?, UpdatedAt = now() WHERE id = ?",
		answerUpdate["Content"], answerUpdate["Images"], id)
	if err != nil {
		return answer, err
	}

	// Check
	rows, err := result.RowsAffected()
	if err != nil {
		return answer, err
	}
	if rows != 1 {
		return answer, fmt.Errorf("expected to affect 1 row, affected %d", rows)
	}
	return answer, nil
}

func (r *AnswerRepository) Delete(id uint64) error {
	// Query
	result, err := r.DBEngine.Exec("DELETE tbAnswer WHERE id = ?", id)
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
