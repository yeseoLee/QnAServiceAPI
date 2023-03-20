package domain

import "time"

// Entity
type Answer struct {
	Id         uint64    `json:"id"`
	QuestionId uint64    `json:"questionId"`
	WriterId   string    `json:"writerId"`
	Content    string    `json:"content"`
	Images     []string  `json:"images"`
	IsAccepted bool      `json:"isAccepted"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	DeletedAt  time.Time `json:"deletedAt"`
	// TODO: 좋아요 수, 조회 수 -> Redis & Batch Insert
}

// DAO
type AnswerDAO struct {
	Id         uint64
	QuestionId uint64
	WriterId   string
	Content    string
	Images     string
	IsAccepted uint8
	CreatedAt  string
	UpdatedAt  string
	DeletedAt  string
}

// DTO
type AnswerInput struct {
	QuestionId uint64   `json:"questionId"`
	WriterId   string   `json:"writerId"`
	Content    string   `json:"content"`
	Images     []string `json:"images"`
}

type AnswerOutput struct {
	Id         uint64    `json:"id"`
	QuestionId uint64    `json:"questionId"`
	WriterId   string    `json:"writerId"`
	Content    string    `json:"content"`
	Images     []string  `json:"images"`
	IsAccepted bool      `json:"isAccepted"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type AnswerSearchOption struct {
	QuestionId uint64 `json:"questionId"`
	WriterId   string `json:"writerId"`
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
}

type AnswerOrderOption struct{}

type AnswerRepository interface {
	FindAllByQuestionId(id uint64, limit int, offset int) ([]*AnswerDAO, error)
	//FindAllByWriterId(writerId string, limit int, offset int) ([]*Answer, error)
	Create(answer *AnswerDAO) (uint64, error)
	Update(id uint64, answerUpdate map[string]interface{}) error
	Delete(id uint64) error
}

type AnswerUseCase interface {
	GetAll(option *AnswerSearchOption) ([]*AnswerOutput, error)
	Create(answerInput *AnswerInput) (uint64, error)
	Edit(id uint64, answerUpdate map[string]interface{}) error
	Accept(id uint64) error
	Delete(id uint64) error
}
