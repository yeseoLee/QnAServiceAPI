package domain

// Entity
type Comment struct {
	Id         uint   `json:"id"`
	QuestionId uint   `json:"questionId"`
	AnswerId   uint   `json:"answerId"`
	Content    string `json:"content"`
	WriterId   string `json:"writerId"`
	CreatedAt  string `json:"createdAt"`
	// TODO:
	// 좋아요 수
}

// DTO
type CommentInput struct {
	QuestionId uint   `json:"questionId"`
	AnswerId   uint   `json:"answerId"`
	Content    string `json:"content"`
	WriterId   string `json:"writerId"`
}

type CommentOutput struct {
	QuestionId uint   `json:"questionId"`
	AnswerId   uint   `json:"answerId"`
	Content    string `json:"content"`
	WriterId   string `json:"writerId"`
	CreatedAt  string `json:"createdAt"`
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
