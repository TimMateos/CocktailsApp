package service

import (
	"CocktailsApp/models"
	"CocktailsApp/repository"
	"fmt"
	"strconv"
	"strings"
)

type InventoryService struct {
	Repo *repository.InventoryRepo
}

func (s *InventoryService) GetInventory() ([]models.InventoryItem, error) {
	return s.Repo.GetInventory()
}

func (s *InventoryService) AddNewItem(item models.InventoryItem) error {
	if strings.TrimSpace(item.Name) == "" {
		return fmt.Errorf("название ингредиента не может быть пустым")
	}

	item.Name = strings.TrimSpace(item.Name)

	return s.Repo.AddNewItem(item)
}

func (s *InventoryService) ToggleItem(id string) error {
	idInt, err := strconv.Atoi(id)
	if idInt <= 0 || err != nil {
		return fmt.Errorf("некорректный ID ингредиента: %s", id)
	}
	err = s.Repo.ToggleItem(idInt)
	if err != nil {
		return fmt.Errorf("ошибка базы данных: %w", err)
	}
	return nil
}

func (s *InventoryService) DeleteItem(id string) error {
	idInt, err := strconv.Atoi(id)
	if idInt <= 0 || err != nil {
		return fmt.Errorf("некорректный ID ингредиента: %s", id)
	}
	err = s.Repo.RemoveItem(idInt)
	if err != nil {
		return fmt.Errorf("ошибка базы данных: %w", err)
	}
	return nil
}

func (s *InventoryService) ToggleAllItems(state bool) error {
	return s.Repo.ToggleAllItems(state)
}
