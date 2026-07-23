package handlers

import (
	"CocktailsApp/models"
	"CocktailsApp/service"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	Service *service.CocktailService
}

func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	searchQuery := r.URL.Query().Get("search")
	categoryQuery := r.URL.Query().Get("category")
	myBarFilter := r.URL.Query().Get("mybar") == "true"

	cocktails, err := h.Service.GetCatalog(searchQuery, categoryQuery, myBarFilter)
	if err != nil {
		http.Error(w, "Внутренняя ошибка", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Ошибка загрузки шаблона", http.StatusInternalServerError)
		log.Println("Ошибка шаблона:", err)
		return
	}
	err = tmpl.Execute(w, cocktails)
	if err != nil {
		log.Println("Ошибка шаблона:", err)
	}
}

func (h *Handler) AddCocktailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/form.html")
		if err != nil {
			http.Error(w, "Ошибка загрузки шаблона", http.StatusInternalServerError)
			log.Println("Ошибка шаблона: ", err)
			return
		}
		tmpl.Execute(w, nil)
		return
	}
	if r.Method == http.MethodPost {
		r.ParseMultipartForm(10 << 20)

		file, header, err := r.FormFile("image")
		var imagePath string

		if err == nil {
			defer file.Close()
			imagePath, err = h.Service.SaveFile(file, header.Filename)

			if err != nil {
				log.Println("Ошибка сохранения данных", err)
			}
		} else if err != http.ErrMissingFile {
			log.Println("Ошибка чтения потока данных", err)
		}

		name := r.FormValue("name")
		category := r.FormValue("category")
		ingredients := r.FormValue("ingredients")
		method := r.FormValue("method")
		serving := r.FormValue("serving")
		isIBA := r.FormValue("is_iba") == "on"

		cocktail := models.Cocktail{
			Name:        name,
			Category:    category,
			Ingredients: ingredients,
			Method:      method,
			Serving:     serving,
			IsIBA:       isIBA,
			ImagePath:   imagePath,
		}

		err = h.Service.CreateCocktail(cocktail)
		if err != nil {
			log.Println("Ошибка создания коктейля:", err)
			http.Error(w, "Ошибка при сохранении: "+err.Error(), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func (h *Handler) DeleteCocktailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")

	if id == "" {
		http.Error(w, "ID не указан", http.StatusBadRequest)
		return
	}
	err := h.Service.DeleteCocktail(id)
	if err != nil {
		log.Println("Ошибка удаления коктейля:", err)
		http.Error(w, "Ошибка при удалении: "+err.Error(), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Handler) EditCocktailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "ID не указан", http.StatusBadRequest)
		}
		cocktail, err := h.Service.GetCocktailByID(id)
		if err != nil {
			log.Println("Ошибка поиска коктейля:", err)
			http.Error(w, "Коктейль не найден", http.StatusNotFound)
			return
		}

		tmpl, err := template.ParseFiles("templates/edit.html")
		if err != nil {
			http.Error(w, "Ошибка загрузки шаблона", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, cocktail)
		return
	}

	if r.Method == http.MethodPost {
		r.ParseMultipartForm(10 << 20)

		idStr := r.FormValue("id")
		idInt, _ := strconv.Atoi(idStr)
		imagePath := r.FormValue("current_image")

		file, header, err := r.FormFile("image")

		if err == nil {
			defer file.Close()
			newImagePath, saveErr := h.Service.SaveFile(file, header.Filename)
			if saveErr != nil {
				log.Println("Ошибка сохранения новой картинки:", saveErr)
			} else {

				imagePath = newImagePath
			}
		}

		updatedCocktail := models.Cocktail{
			ID:          idInt,
			Name:        r.FormValue("name"),
			Category:    r.FormValue("category"),
			Ingredients: r.FormValue("ingredients"),
			Method:      r.FormValue("method"),
			Serving:     r.FormValue("serving"),
			IsIBA:       r.FormValue("is_iba") == "on",
			ImagePath:   imagePath,
		}

		updateErr := h.Service.UpdateCocktail(updatedCocktail)
		if updateErr != nil {
			log.Println("Ошибка обновления коктейля:", updateErr)
			http.Error(w, "Ошибка при обновлении: "+updateErr.Error(), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
