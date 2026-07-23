package repository

import (
	"CocktailsApp/models"
	"database/sql"
)

type InventoryRepo struct {
	DB *sql.DB
}

func (r *InventoryRepo) GetInventory() ([]models.InventoryItem, error) {
	rows, err := r.DB.Query("SELECT id, name, in_stock FROM inventory ORDER BY name ASC")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []models.InventoryItem

	for rows.Next() {
		var item models.InventoryItem
		err = rows.Scan(&item.ID, &item.Name, &item.InStock)
		if err != nil {
			continue
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *InventoryRepo) AddNewItem(item models.InventoryItem) error {
	_, err := r.DB.Exec("INSERT INTO inventory (name) VALUES ($1) ON CONFLICT (name) DO NOTHING", item.Name)
	return err
}

func (r *InventoryRepo) ToggleItem(id int) error {
	_, err := r.DB.Exec("UPDATE inventory SET in_stock = NOT in_stock WHERE id = $1", id)
	return err
}

func (r *InventoryRepo) RemoveItem(id int) error {
	_, err := r.DB.Exec("DELETE FROM inventory WHERE id = $1", id)
	return err
}

func (r *InventoryRepo) ToggleAllItems(state bool) error {
	_, err := r.DB.Exec("UPDATE inventory SET in_stock = $1", state)
	return err
}
