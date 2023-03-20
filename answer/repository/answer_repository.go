package answer

import (
	"database/sql"
	"fmt"
	"strings"

	"qna/datasource"
	model "qna/domain"
	"qna/util"
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

func (r *AnswerRepository) FindAllByQuestionId(id uint64, limit int, offset int) ([]*model.AnswerDAO, error) {

	var aList []*model.AnswerDAO

	// Query
	rows, err := r.DBEngine.Query("SELECT id, questionId, content, writerId, images, isAccepted, createdAt, updatedAt FROM tbAnswer WHERE questionId = ? LIMIT ? OFFSET ?", id, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Read
	for rows.Next() {
		var a model.AnswerDAO
		rows.Scan(&a.Id, &a.QuestionId, &a.Content, &a.WriterId, &a.Images, &a.IsAccepted, &a.CreatedAt, &a.UpdatedAt)
		aList = append(aList, &a)
	}

	return aList, nil
}

// func (r *AnswerRepository) FindAllByWriterId(writerId string, limit int, offset int) ([]*model.Answer, error) {}

func (r *AnswerRepository) Create(answer *model.AnswerDAO) (uint64, error) {
	now := util.DateTimeNow()

	// Query
	result, err := r.DBEngine.Exec("INSERT INTO tbAnswer (`questionId`,`content`, `writerId`, `images`,`createdAt` ,`updatedAt`) VALUES (?,?,?,?,?,?)",
		answer.QuestionId, answer.Content, answer.WriterId, answer.Images, now, now)
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

func (r *AnswerRepository) Update(id uint64, answerUpdate map[string]interface{}) error {
	// TODO: map key-value check & make query logic
	now := util.DateTimeNow()

	// type casting
	images, ok := answerUpdate["Images"].([]string)
	if !ok {
		return fmt.Errorf("unexpected parameter, wants: Images []string")
	}
	images_str := strings.Join(images, ",")

	// Query
	result, err := r.DBEngine.Exec("UPDATE tbAnswer SET content = ?, images = ?, UpdatedAt = ? WHERE id = ?",
		answerUpdate["Content"], images_str, now, id)
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

func (r *AnswerRepository) Delete(id uint64) error {
	// Query
	result, err := r.DBEngine.Exec("DELETE FROM `tbAnswer` WHERE id = ?", id)
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
