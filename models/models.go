package models

type Cocktail struct {
	ID          int
	Name        string
	Category    string
	Ingredients string
	Method      string
	Serving     string
	IsIBA       bool
	ImagePath   string
	MatchCount  int
}
type InventoryItem struct {
	ID      int
	Name    string
	InStock bool
}
