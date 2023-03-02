package domain

import "time"

// Entity
type Question struct {
	Id        uint      `json:"id"`
	WriterId  string    `json:"writerId"`
	Title     string    `json:"title"`
	Content   string    `json:"body"`
	Tags      []string  `json:"tags"`
	Images    []string  `json:"images"`
	IsAccept  bool      `json:"IsAccept"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deleteddAt"`
	// TODO:
	// 좋아요 수
	// 조회 수
}

// DTO
type QuestionInput struct {
	WriterId string   `json:"writerId"`
	Title    string   `json:"title"`
	Content  string   `json:"body"`
	Tags     []string `json:"tags"`
	Images   []string `json:"images"`
}

type QuestionOutput struct {
	Id        uint      `json:"id"`
	WriterId  string    `json:"writer"`
	Title     string    `json:"title"`
	Content   string    `json:"body"`
	Tags      []string  `json:"tags"`
	Images    []string  `json:"images"`
	IsAccept  bool      `json:"IsAccept"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type QuestionSearchOption struct {
	Title    string `json:"title"`
	Content  string `json:"body"`
	WriterId string `json:"writerId"`
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
}

type QuestionOrderOption struct {
}

type QuestionRepository interface {
	FindById(id uint) (*Question, error)
	FindAll(limit, offset int) ([]*Question, error)
	//FindAllByTags(tags []string) ([]*Question, error)
	Create(questionInput *QuestionInput) (*Question, error)
	Update(id uint, questionUpdate map[string]interface{}) (*Question, error)
	Delete(id uint) error
}

type QuestionUseCase interface {
	Get(id uint) (*QuestionOutput, error)
	GetAll(option *QuestionSearchOption) ([]*QuestionOutput, error)
	Create(questionInput *QuestionInput) (*QuestionOutput, error)
	Edit(WriterId string, id uint, questionEdit map[string]interface{}) (*QuestionOutput, error)
	Accept(WriterId string, id uint) error
	Delete(WriterId string, id uint) error
}
