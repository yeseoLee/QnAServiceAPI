package domain

import "time"

// Entity
type Answer struct {
	Id         uint      `json:"id"`
	QuestionId uint      `json:"questionId"`
	Content    string    `json:"content"`
	WriterId   string    `json:"writerId"`
	Images     []string  `json:"images"`
	IsAccepted bool      `json:"IsAccepted"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	DeletedAt  time.Time `json:"deleteddAt"`
	// TODO:
	// 좋아요 수
	// 조회 수
}

// DTO
type AnswerInput struct {
	QuestionId uint     `json:"questionId"`
	Content    string   `json:"content"`
	WriterId   string   `json:"writerId"`
	Images     []string `json:"images"`
}

type AnswerOutput struct {
	Id         uint      `json:"id"`
	QuestionId uint      `json:"questionId"`
	Content    string    `json:"content"`
	WriterId   string    `json:"writerId"`
	Images     []string  `json:"images"`
	IsAccepted bool      `json:"IsAccepted"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type AnswerSearchOption struct {
	QuestionId uint   `json:"questionId"`
	WriterId   string `json:"writerId"`
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
}

type AnswerOrderOption struct{}

type AnswerRepository interface {
	FindAllByQuestionId(id uint, limit int, offset int) ([]*Answer, error)
	//FindAllByWriterId(writerId string, limit int, offset int) ([]*Answer, error)
	Create(answerInput *AnswerInput) (*Answer, error)
	Update(id uint, answerUpdate map[string]interface{}) (*Answer, error)
	Delete(id uint) error
}

type AnswerUseCase interface {
	GetAll(option *AnswerSearchOption) ([]*AnswerOutput, error)
	Create(answerInput *AnswerInput) (*AnswerOutput, error)
	Edit(WriterId string, id uint, answerUpdate map[string]interface{}) (*AnswerOutput, error)
	Accept(QuestionWriterId string, id uint) error
	Delete(WriterId string, id uint) error
}
