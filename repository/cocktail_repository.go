package repository

import (
	"CocktailsApp/models"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type CocktailRepo struct {
	DB *sql.DB
}

func (r *CocktailRepo) GetCocktails(searchQuery, categoryQuery string) ([]models.Cocktail, error) {
	query := "SELECT id, name, category, ingredients, method, serving, is_iba, image_path FROM cocktails WHERE 1=1"
	var args []interface{}
	argId := 1

	if searchQuery != "" {
		query += fmt.Sprintf(" AND (name ILIKE $%d OR ingredients ILIKE $%d)", argId, argId+1)
		args = append(args, "%"+searchQuery+"%", "%"+searchQuery+"%")
		argId += 2
	}

	if categoryQuery != "" && categoryQuery != "ALL" {
		query += fmt.Sprintf(" AND category = $%d", argId)
		args = append(args, categoryQuery)
		argId++
	}
	query += " ORDER BY created_at DESC"
	rows, err := r.DB.Query(query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cocktails []models.Cocktail

	for rows.Next() {
		var c models.Cocktail
		var isIbaNull sql.NullBool
		var imagePathNull sql.NullString

		err := rows.Scan(&c.ID, &c.Name, &c.Category, &c.Ingredients, &c.Method, &c.Serving, &isIbaNull, &imagePathNull)
		if err != nil {
			log.Fatal("Ошибка при считывании строки: ", err)
			continue
		}
		if isIbaNull.Valid {
			c.IsIBA = isIbaNull.Bool
		}
		if imagePathNull.Valid {
			c.ImagePath = imagePathNull.String
		}

		cocktails = append(cocktails, c)
	}
	return cocktails, nil
}

func (r *CocktailRepo) GetActiveInventory() ([]string, error) {
	rows, err := r.DB.Query("SELECT name FROM inventory WHERE in_stock = TRUE")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shelf []string
	for rows.Next() {
		var item string
		rows.Scan(&item)
		shelf = append(shelf, strings.ToLower(item))
	}
	return shelf, nil
}

func (r *CocktailRepo) AddNewCocktail(c models.Cocktail) error {
	insertQuery := `
			INSERT INTO cocktails(name, category, ingredients, method, serving, is_iba, image_path)
			VALUES ($1, $2, $3, $4, $5, $6, $7)
			`
	_, err := r.DB.Exec(insertQuery, c.Name, c.Category, c.Ingredients, c.Method, c.Serving, c.IsIBA, c.ImagePath)
	if err != nil {
		return err
	}
	return nil
}

func (r *CocktailRepo) DeleteCocktail(id int) error {
	_, err := r.DB.Exec("DELETE FROM cocktails WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *CocktailRepo) GetCocktailByID(id int) (models.Cocktail, error) {
	var c models.Cocktail
	var isIbaNull sql.NullBool
	var imagePathNull sql.NullString

	query := `SELECT id, name, category, ingredients, method, serving, is_iba, image_path 
				FROM cocktails WHERE id = $1`
	err := r.DB.QueryRow(query, id).Scan(
		&c.ID, &c.Name, &c.Category, &c.Ingredients, &c.Method, &c.Serving, &isIbaNull, &imagePathNull,
	)

	if err != nil {
		return c, err
	}
	if isIbaNull.Valid {
		c.IsIBA = isIbaNull.Bool
	}
	if imagePathNull.Valid {
		c.ImagePath = imagePathNull.String
	}
	return c, nil
}

func (r *CocktailRepo) UpdateCocktail(c models.Cocktail) error {
	updateQuery := `
                UPDATE cocktails
				SET name=$1, category=$2, ingredients=$3, method=$4, serving=$5, is_iba=$6, image_path=$7
				WHERE id=$8
				`
	_, err := r.DB.Exec(updateQuery, c.Name, c.Category, c.Ingredients, c.Method, c.Serving, c.IsIBA, c.ImagePath, c.ID)
	return err
}
