package domain

type Comment struct {
	CommentID      int       `json:"commentId"`
	CommentContent string    `json:"commentContent"`
	Replies        []Comment `json:"replies,omitempty"`
}

type CommentRepository interface {
	GetAll() ([]Comment, error)
	CountComments(comments []Comment) int
}
