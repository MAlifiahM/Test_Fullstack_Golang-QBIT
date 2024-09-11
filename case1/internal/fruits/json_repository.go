package fruits

import (
	"case1/internal/domain"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

type JSONFruitRepository struct {
	filepath string
}

func NewJSONFruitRepository(filepath string) *JSONFruitRepository {
	return &JSONFruitRepository{filepath: filepath}
}

func (repo *JSONFruitRepository) GetAll() ([]domain.Fruit, error) {
	file, err := os.Open(repo.filepath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	byteValue, _ := io.ReadAll(file)

	var fruits []domain.Fruit
	err = json.Unmarshal(byteValue, &fruits)
	if err != nil {
		return nil, err
	}

	if len(fruits) == 0 {
		return nil, errors.New("no fruits found")
	}

	return fruits, nil
}
