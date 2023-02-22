package domain

// Entity
type Answer struct {
	Id         uint     `json:"id"`
	QuestionId uint     `json:"questionId"`
	Content    string   `json:"content"`
	WriterId   string   `json:"writerId"`
	Images     []string `json:"images"`
	IsAccepted bool     `json:"IsAccepted"`
	CreateTime string   `json:"createTime"`
	UpdateTime string   `json:"updateTime"`
	// ---------
	// 좋아요 수
	// 조회 수
}

// DTO
type AnswerInput struct {
	QuestionId uint     `json:"questionId"`
	Content    string   `json:"content"`
	WriterId   string   `json:"writerId"`
	Images     []string `json:"images"`
	IsAccepted bool     `json:"IsAccepted"`
}

type AnswerOutput struct {
	Id         uint     `json:"id"`
	QuestionId uint     `json:"questionId"`
	Content    string   `json:"content"`
	WriterId   string   `json:"writerId"`
	Images     []string `json:"images"`
	IsAccepted bool     `json:"IsAccepted"`
	CreateTime string   `json:"createTime"`
	UpdateTime string   `json:"updateTime"`
}

type AnswersearchOption struct {
	QuestionId uint   `json:"questionId"`
	WriterId   string `json:"writerId"`
}

type AnswerRepository interface {
	FindAllByQuestionId(id uint) ([]*Answer, error)
	//FindAllByWriterId(writerId string, limit int, offset int) ([]*Answer, error)
	Create(answerInput *AnswerInput) (*Answer, error)
	Update(id uint, answerUpdate map[string]interface{}) (*Answer, error)
	Delete(id uint) error
}

type AnswerUseCase interface {
	GetAll(option *AnswersearchOption) ([]*AnswerOutput, error)
	Create(answerInput *AnswerInput) (*AnswerOutput, error)
	Edit(WriterId string, id uint, answerUpdate map[string]interface{}) (*AnswerOutput, error)
	Accept(QuestionWriterId string, id uint) error
	Delete(WriterId string, id uint) error
}
