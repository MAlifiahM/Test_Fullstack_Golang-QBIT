package fruits

import (
	"case1/internal/domain"
	"strings"
)

type FruitService struct {
	repo domain.FruitRepository
}

func NewFruitService(repo domain.FruitRepository) *FruitService {
	return &FruitService{repo: repo}
}

func (s *FruitService) GetFruits() ([]string, error) {
	fruits, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	var fruitsName []string

	for _, fruit := range fruits {
		fruitsName = append(fruitsName, fruit.Name)
	}

	return fruitsName, nil
}

func (s *FruitService) GetFruitsByType() (map[domain.FruitType]domain.ResponseFruitWTotalStock, error) {
	fruits, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	result := map[domain.FruitType]domain.ResponseFruitWTotalStock{
		domain.FruitTypeImport: {
			Fruits:     []domain.Fruit{},
			TotalStock: 0,
		},
		domain.FruitTypeLocal: {
			Fruits:     []domain.Fruit{},
			TotalStock: 0,
		},
	}

	for _, fruit := range fruits {
		result[fruit.Type] = domain.ResponseFruitWTotalStock{
			Fruits:     append(result[fruit.Type].Fruits, fruit),
			TotalStock: result[fruit.Type].TotalStock + fruit.Stock,
		}
	}

	return result, nil
}

func (s *FruitService) GetUniqueFruits() ([]string, error) {
	fruits, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	uniqueFruits := make(map[string]bool)
	fruitNames := []string{}

	for _, fruit := range fruits {
		normalizedFruitName := strings.ToLower(fruit.Name)
		if !uniqueFruits[normalizedFruitName] {
			uniqueFruits[normalizedFruitName] = true
			fruitNames = append(fruitNames, fruit.Name)
		}
	}

	return fruitNames, nil
}
