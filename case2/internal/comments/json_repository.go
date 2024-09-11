package comments

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"case2/internal/domain"
)

type JSONCommentRepository struct {
	filepath string
}

func NewJSONCommentRepository(filepath string) *JSONCommentRepository {
	return &JSONCommentRepository{filepath: filepath}
}

func (repo *JSONCommentRepository) GetAll() ([]domain.Comment, error) {
	file, err := os.Open(repo.filepath)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	byteValue, _ := io.ReadAll(file)

	var comments []domain.Comment
	err = json.Unmarshal(byteValue, &comments)
	if err != nil {
		return nil, err
	}

	if len(comments) == 0 {
		return nil, errors.New("no comments found")
	}

	return comments, nil
}

func (repo *JSONCommentRepository) CountComments(comments []domain.Comment) int {
	totalComments := 0

	var countReplies func([]domain.Comment) int
	countReplies = func(comments []domain.Comment) int {
		total := 0
		for _, comment := range comments {
			total += 1
			if len(comment.Replies) > 0 {
				total += countReplies(comment.Replies)
			}
		}
		return total
	}

	totalComments = countReplies(comments)
	return totalComments
}
