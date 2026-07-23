package service

import (
	"CocktailsApp/models"
	"CocktailsApp/repository"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type CocktailService struct {
	Repo *repository.CocktailRepo
}

func (s *CocktailService) GetCatalog(search, category string, useMyBar bool) ([]models.Cocktail, error) {
	cocktails, err := s.Repo.GetCocktails(search, category)
	if err != nil {
		return nil, err
	}
	if useMyBar {
		shelf, _ := s.Repo.GetActiveInventory()
		for i := range shelf {
			shelf[i] = strings.ToLower(shelf[i])
		}

		var matchedCocktails []models.Cocktail
		for _, c := range cocktails {
			count := 0
			ingText := strings.ToLower(c.Ingredients)
			for _, item := range shelf {
				if strings.Contains(ingText, item) {
					count++
				}
			}
			if count > 0 {
				c.MatchCount = count
				matchedCocktails = append(matchedCocktails, c)
			}
		}
		sort.Slice(matchedCocktails, func(i, j int) bool {
			return matchedCocktails[i].MatchCount > matchedCocktails[j].MatchCount
		})
		return matchedCocktails, nil
	}
	return cocktails, nil
}

func (s *CocktailService) SaveFile(file io.Reader, filename string) (string, error) {

	name := fmt.Sprintf("%d_%s", time.Now().Unix(), filename)
	savePath := filepath.Join("uploads", name)

	dst, err := os.Create(savePath)
	if err != nil {
		return "", err
	}

	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}
	webPath := "/uploads/" + name
	return webPath, nil
}

func (s *CocktailService) CreateCocktail(c models.Cocktail) error {

	if strings.TrimSpace(c.Name) == "" {
		return fmt.Errorf("название коктейля не может быть пустым")
	}
	if strings.TrimSpace(c.Ingredients) == "" {
		return fmt.Errorf("коктейль не может состоять из воздуха")
	}
	if strings.TrimSpace(c.Method) == "" {
		return fmt.Errorf("ингрединты коктейля нужно хотя бы смешать")
	}

	c.Name = strings.TrimSpace(c.Name)
	c.Ingredients = strings.TrimSpace(c.Ingredients)
	c.Method = strings.TrimSpace(c.Method)
	c.Serving = strings.TrimSpace(c.Serving)

	err := s.Repo.AddNewCocktail(c)
	if err != nil {
		return fmt.Errorf("ошибка при сохранении коктейля в репозитории: %w", err)
	}
	return nil
}

func (s *CocktailService) DeleteCocktail(id string) error {
	idInt, err := strconv.Atoi(id)
	if idInt <= 0 || err != nil {
		return fmt.Errorf("некорректный ID коктейля: %s", id)
	}
	err = s.Repo.DeleteCocktail(idInt)
	if err != nil {
		return fmt.Errorf("ошибка базы данных: %w", err)
	}
	return nil
}

func (s *CocktailService) UpdateCocktail(c models.Cocktail) error {
	if c.ID <= 0 {
		fmt.Errorf("некорректный ID")
	}
	if strings.TrimSpace(c.Name) == "" {
		fmt.Errorf("навзвание коктейля не может быть пустым")
	}
	if strings.TrimSpace(c.Ingredients) == "" {
		return fmt.Errorf("коктейль не может состоять из воздуха")
	}

	c.Name = strings.TrimSpace(c.Name)
	c.Ingredients = strings.TrimSpace(c.Ingredients)
	c.Method = strings.TrimSpace(c.Method)
	c.Serving = strings.TrimSpace(c.Serving)

	return s.Repo.UpdateCocktail(c)
}

func (s *CocktailService) GetCocktailByID(id string) (models.Cocktail, error) {
	idInt, err := strconv.Atoi(id)
	if idInt <= 0 || err != nil {
		return models.Cocktail{}, fmt.Errorf("неккоректный ID: %s", id)
	}
	return s.Repo.GetCocktailByID(idInt)
}
