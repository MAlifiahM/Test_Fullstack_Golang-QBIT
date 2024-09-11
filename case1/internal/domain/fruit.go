package domain

type FruitType string

const (
	FruitTypeImport FruitType = "IMPORT"
	FruitTypeLocal  FruitType = "LOCAL"
)

type Fruit struct {
	ID    int       `json:"fruitId"`
	Name  string    `json:"fruitName"`
	Type  FruitType `json:"fruitType"`
	Stock int       `json:"stock"`
}

type ResponseFruitWTotalStock struct {
	Fruits     []Fruit `json:"fruits"`
	TotalStock int     `json:"total_stock"`
}

type FruitRepository interface {
	GetAll() ([]Fruit, error)
}
