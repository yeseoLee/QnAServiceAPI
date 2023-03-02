package domain

import "time"

// Entity
type Question struct {
	Id        uint64    `json:"id"`
	WriterId  string    `json:"writerId"`
	Title     string    `json:"title"`
	Content   string    `json:"body"`
	Tags      []string  `json:"tags"`
	Images    []string  `json:"images"`
	IsAccept  bool      `json:"IsAccept"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deleteddAt"`
	// TODO: 좋아요 수, 조회 수 -> Redis & Batch Insert
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
	Id        uint64    `json:"id"`
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
	FindById(id uint64) (*Question, error)
	FindAll(limit, offset int) ([]*Question, error)
	//FindAllByTags(tags []string) ([]*Question, error)
	Create(questionInput *QuestionInput) (*Question, error)
	Update(id uint64, questionUpdate map[string]interface{}) (*Question, error)
	Delete(id uint64) error
}

type QuestionUseCase interface {
	Get(id uint64) (*QuestionOutput, error)
	GetAll(option *QuestionSearchOption) ([]*QuestionOutput, error)
	Create(questionInput *QuestionInput) (*QuestionOutput, error)
	Edit(id uint64, questionEdit map[string]interface{}) (*QuestionOutput, error)
	Accept(id uint64) error
	Delete(id uint64) error
}
