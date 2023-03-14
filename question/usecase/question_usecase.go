package question

import (
	"qna/domain"
	"qna/util"
	"strings"
)

type questionUsecase struct {
	questionRepo domain.QuestionRepository
}

func NewQuestionUseCase(u domain.QuestionRepository) domain.QuestionUseCase {
	return &questionUsecase{
		questionRepo: u,
	}
}

func (u *questionUsecase) Get(id uint64) (*domain.QuestionOutput, error) {
	q, err := u.questionRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	qo := u.transferOutput(q)
	return qo, nil
}

func (u *questionUsecase) GetAll(option *domain.QuestionSearchOption) ([]*domain.QuestionOutput, error) {
	// TODO: 검색 옵션 적용
	qList, err := u.questionRepo.FindAll(option.Limit, option.Offset)
	if err != nil {
		return nil, err
	}
	var qoList []*domain.QuestionOutput
	for _, v := range qList {
		qo := u.transferOutput(v)
		qoList = append(qoList, qo)
	}
	return qoList, nil
}

func (u *questionUsecase) Create(questionInput *domain.QuestionInput) (uint64, error) {
	qdto := u.transferDAO(questionInput)
	qId, err := u.questionRepo.Create(qdto)
	if err != nil {
		return 0, err
	}
	return qId, nil
}

func (u *questionUsecase) Edit(id uint64, questionEdit map[string]interface{}) error {
	err := u.questionRepo.Update(id, questionEdit)
	if err != nil {
		return err
	}
	return nil
}

func (u *questionUsecase) Accept(id uint64) error {
	// TODO: 채택 로직 개선
	err := u.questionRepo.Update(id, map[string]interface{}{"IsAccept": true})
	return err
}

func (u *questionUsecase) Delete(id uint64) error {
	return u.questionRepo.Delete(id)
}

func (u *questionUsecase) transferDAO(question *domain.QuestionInput) *domain.QuestionDAO {
	dao := &domain.QuestionDAO{}
	dao.WriterId = question.WriterId
	dao.Title = question.Title
	dao.Content = question.Content
	dao.Tags = strings.Join(question.Tags, ",")
	dao.Images = strings.Join(question.Images, ",")
	return dao
}

func (u *questionUsecase) transferOutput(dao *domain.QuestionDAO) *domain.QuestionOutput {
	qo := &domain.QuestionOutput{}
	qo.Id = dao.Id
	qo.WriterId = dao.WriterId
	qo.Title = dao.Title
	qo.Content = dao.Content
	qo.Tags = strings.Split(dao.Tags, ",")
	qo.Images = strings.Split(dao.Images, ",")
	qo.IsAccept = util.Uint8ToBool(dao.IsAccept)
	qo.UpdatedAt = util.DateTimeStringToTime(dao.UpdatedAt)
	return qo
}
