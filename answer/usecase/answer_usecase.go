package answer

import "qna/domain"

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

func (u *answerUsecase) Create(answerInput *domain.AnswerInput) (*domain.AnswerOutput, error) {
	q, err := u.answerRepo.Create(answerInput)
	if err != nil {
		return nil, err
	}
	qo := u.transferOutput(q)
	return qo, nil
}

func (u *answerUsecase) Edit(id uint64, answerEdit map[string]interface{}) (*domain.AnswerOutput, error) {
	q, err := u.answerRepo.Update(id, answerEdit)
	if err != nil {
		return nil, err
	}
	qo := u.transferOutput(q)
	return qo, nil
}

func (u *answerUsecase) Accept(id uint64) error {
	// TODO: 채택 로직 개선
	_, err := u.answerRepo.Update(id, map[string]interface{}{"IsAccept": true})
	return err
}

func (u *answerUsecase) Delete(id uint64) error {
	return u.answerRepo.Delete(id)
}

func (u *answerUsecase) transferOutput(answer *domain.Answer) *domain.AnswerOutput {
	ao := &domain.AnswerOutput{}
	ao.Id = answer.Id
	ao.Content = answer.Content
	ao.WriterId = answer.WriterId
	ao.Images = answer.Images
	return ao
}
