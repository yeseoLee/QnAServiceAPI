package answer

import (
	"qna/domain"
	"qna/util"
	"strings"
)

type answerUsecase struct {
	answerRepo domain.AnswerRepository
}

func NewAnswerUseCase(u domain.AnswerRepository) domain.AnswerUseCase {
	return &answerUsecase{
		answerRepo: u,
	}
}

func (u *answerUsecase) GetAll(option *domain.AnswerSearchOption) ([]*domain.AnswerOutput, error) {
	qList, err := u.answerRepo.FindAllByQuestionId(option.QuestionId, option.Limit, option.Offset)
	if err != nil {
		return nil, err
	}
	var qoList []*domain.AnswerOutput
	for _, v := range qList {
		qo := u.transferOutput(v)
		qoList = append(qoList, qo)
	}
	return qoList, nil
}

func (u *answerUsecase) Create(answerInput *domain.AnswerInput) (uint64, error) {
	adto := u.transferDAO(answerInput)
	aId, err := u.answerRepo.Create(adto)
	if err != nil {
		return 0, err
	}
	return aId, nil
}

func (u *answerUsecase) Edit(id uint64, answerEdit map[string]interface{}) error {
	err := u.answerRepo.Update(id, answerEdit)
	if err != nil {
		return err
	}
	return nil
}

func (u *answerUsecase) Accept(id uint64) error {
	// TODO: 채택 로직 개선
	err := u.answerRepo.Update(id, map[string]interface{}{"IsAccept": true})
	return err
}

func (u *answerUsecase) Delete(id uint64) error {
	return u.answerRepo.Delete(id)
}

func (u *answerUsecase) transferDAO(answer *domain.AnswerInput) *domain.AnswerDAO {
	dao := &domain.AnswerDAO{}
	dao.QuestionId = answer.QuestionId
	dao.WriterId = answer.WriterId
	dao.Content = answer.Content
	dao.Images = strings.Join(answer.Images, ",")
	return dao
}

func (u *answerUsecase) transferOutput(dao *domain.AnswerDAO) *domain.AnswerOutput {
	ao := &domain.AnswerOutput{}
	ao.Id = dao.Id
	ao.QuestionId = dao.QuestionId
	ao.Content = dao.Content
	ao.WriterId = dao.WriterId
	ao.Images = strings.Split(dao.Images, ",")
	ao.IsAccepted = util.Uint8ToBool(dao.IsAccepted)
	ao.UpdatedAt = util.DateTimeStringToTime(dao.UpdatedAt)
	return ao
}
