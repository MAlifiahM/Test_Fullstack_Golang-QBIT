package comments

import "case2/internal/domain"

type CommentService struct {
	repo domain.CommentRepository
}

func NewCommentService(repo domain.CommentRepository) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) GetComments() ([]domain.Comment, error) {
	return s.repo.GetAll()
}

func (s *CommentService) CountTotalComments() (int, error) {
	comments, err := s.repo.GetAll()
	if err != nil {
		return 0, err
	}
	return s.repo.CountComments(comments), nil
}
