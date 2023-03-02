package question

import "qna/domain"

type questionUsecase struct {
	questionRepo domain.QuestionRepository
}

func NewQuestionUseCase(u domain.QuestionRepository) domain.QuestionUseCase {
	return &questionUsecase{
		questionRepo: u,
	}
}

func (u *questionUsecase) Get(id uint) (*domain.QuestionOutput, error) {
	q, err := u.questionRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	qo := u.transferOutput(q)
	return qo, nil
}

func (u *questionUsecase) GetAll(option *domain.QuestionSearchOption) ([]*domain.QuestionOutput, error) {
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

func (u *questionUsecase) Create(questionInput *domain.QuestionInput) (*domain.QuestionOutput, error) {
	q, err := u.questionRepo.Create(questionInput)
	if err != nil {
		return nil, err
	}
	qo := u.transferOutput(q)
	return qo, nil
}

func (u *questionUsecase) Edit(WriterId string, id uint, questionEdit map[string]interface{}) (*domain.QuestionOutput, error) {
	q, err := u.questionRepo.Update(id, questionEdit)
	if err != nil {
		return nil, err
	}
	qo := u.transferOutput(q)
	return qo, nil
}

func (u *questionUsecase) Accept(WriterId string, id uint) error {
	// TODO
	_, err := u.questionRepo.Update(id, map[string]interface{}{"IsAccept": true})
	return err
}

func (u *questionUsecase) Delete(WriterId string, id uint) error {
	return u.questionRepo.Delete(id)
}

func (u *questionUsecase) transferOutput(question *domain.Question) *domain.QuestionOutput {
	qo := &domain.QuestionOutput{}
	qo.Id = question.Id
	qo.Title = question.Title
	qo.Content = question.Content
	qo.WriterId = question.WriterId
	qo.Images = question.Images
	qo.UpdatedAt = question.UpdatedAt
	return qo
}
