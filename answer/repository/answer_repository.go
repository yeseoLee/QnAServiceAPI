package answer

import (
	"database/sql"

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

func (r *AnswerRepository) FindAllByQuestionId(id uint) ([]*model.Answer, error) {
	return nil, nil
}

// func (r *AnswerRepository) FindAllByWriterId(writerId string, limit int, offset int) ([]*model.Answer, error) {
// 	return nil, nil
// }

func (r *AnswerRepository) Create(answerInput *model.AnswerInput) (*model.Answer, error) {
	return nil, nil
}

func (r *AnswerRepository) Update(id uint, answerUpdate map[string]interface{}) (*model.Answer, error) {
	return nil, nil
}

func (r *AnswerRepository) Delete(id uint) error {
	return nil
}
