package domain

import "time"

// Entity
type Comment struct {
	Id         uint64    `json:"id"`
	QuestionId uint64    `json:"questionId"`
	AnswerId   uint64    `json:"answerId"`
	WriterId   string    `json:"writerId"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"createdAt"`
	DeletedAt  time.Time `json:"deletedAt"`
	// TODO: 좋아요 수 -> Redis & Batch Insert
}

// DAO
type CommentDAO struct {
	Id         uint64
	QuestionId uint64
	AnswerId   uint64
	WriterId   string
	Content    string
	CreatedAt  string
	DeletedAt  string
}

// DTO
type CommentInput struct {
	QuestionId uint64 `json:"questionId"`
	AnswerId   uint64 `json:"answerId"`
	WriterId   string `json:"writerId"`
	Content    string `json:"content"`
}

type CommentOutput struct {
	Id         uint64    `json:"id"`
	QuestionId uint64    `json:"questionId"`
	AnswerId   uint64    `json:"answerId"`
	WriterId   string    `json:"writerId"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"createdAt"`
	DeletedAt  time.Time `json:"deletedAt"`
}

type CommentRepository interface {
	FindAllByPostId(id uint64, limit int, offset int) ([]*Comment, error)
	Create(commentInput *CommentInput) (*Comment, error)
	Delete(id uint64) error
}

type CommentUseCase interface {
	GetAll(limit, offset int) ([]*CommentOutput, error)
	Create(commentInput *CommentInput) (*CommentOutput, error)
	Delete(id uint64) error
}
