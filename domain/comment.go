package domain

import "time"

// Entity
type Comment struct {
	Id         uint      `json:"id"`
	QuestionId uint      `json:"questionId"`
	AnswerId   uint      `json:"answerId"`
	WriterId   string    `json:"writerId"`
	Content    string    `json:"content"`
	CreatedAt  string    `json:"createdAt"`
	DeletedAt  time.Time `json:"deleteddAt"`
	// TODO: 좋아요 수 -> Redis & Batch Insert
}

// DTO
type CommentInput struct {
	QuestionId uint   `json:"questionId"`
	AnswerId   uint   `json:"answerId"`
	WriterId   string `json:"writerId"`
	Content    string `json:"content"`
}

type CommentOutput struct {
	QuestionId uint      `json:"questionId"`
	AnswerId   uint      `json:"answerId"`
	WriterId   string    `json:"writerId"`
	Content    string    `json:"content"`
	CreatedAt  string    `json:"createdAt"`
	DeletedAt  time.Time `json:"deleteddAt"`
}

type CommentRepository interface {
	FindAllByPostId(id uint, limit int, offset int) ([]*Comment, error)
	Create(commentInput *CommentInput) (*Comment, error)
	Delete(id uint) error
}

type CommentUseCase interface {
	GetAll(limit, offset int) ([]*CommentOutput, error)
	Create(commentInput *CommentInput) (*CommentOutput, error)
	Delete(id uint) error
}
