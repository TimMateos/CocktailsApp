package handlers

import (
	"CocktailsApp/models"
	"CocktailsApp/service"
	"html/template"
	"log"
	"net/http"
)

type InventoryHandler struct {
	Service *service.InventoryService
}

func (h *InventoryHandler) InventoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	items, err := h.Service.GetInventory()
	if err != nil {
		log.Println("[SYSTEM] Ошибка получения инвентаря:", err)
		http.Error(w, "Внутренняя ошибка базы данных склада", http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles("templates/inventory.html")
	if err != nil {
		log.Println("[SYSTEM] Ошибка загрузки шаблона инвентаря:", err)
		http.Error(w, "Ошибка загрузки кибер-интерфейса", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, items)
	if err != nil {
		log.Println("[SYSTEM] Ошибка рендера шаблона:", err)
	}
}

func (h *InventoryHandler) AddItemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	name := r.FormValue("name")

	item := models.InventoryItem{
		Name: name,
	}

	err := h.Service.AddNewItem(item)
	if err != nil {
		log.Println("Ошибка создания ингредиента:", err)
		http.Error(w, "Ошибка при добавлении: "+err.Error(), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/inventory", http.StatusSeeOther)
}

func (h *InventoryHandler) ToggleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	id := r.FormValue("id")
	if id == "" {
		http.Error(w, "ID не указан", http.StatusBadRequest)
		return
	}
	err := h.Service.ToggleItem(id)
	if err != nil {
		log.Println("[SYSTEM] Ошибка переключения статуса ингредиента:", err)
	}

	http.Redirect(w, r, "/inventory", http.StatusSeeOther)
}

func (h *InventoryHandler) DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	id := r.FormValue("id")
	err := h.Service.DeleteItem(id)
	if err != nil {
		log.Println("Ошибка удаления ингредиента:", err)
		http.Error(w, "Ошибка при удалении: "+err.Error(), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/inventory", http.StatusSeeOther)
}

func (h *InventoryHandler) ToggleAllItemsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	state := r.FormValue("state") == "true"
	err := h.Service.ToggleAllItems(state)
	if err != nil {
		http.Error(w, "Ошибка при изменении статуса: "+err.Error(), http.StatusBadRequest)
	}
	http.Redirect(w, r, "/inventory", http.StatusSeeOther)
}
