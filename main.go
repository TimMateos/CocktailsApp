package main

import (
	"CocktailsApp/db"
	"CocktailsApp/handlers"
	"CocktailsApp/repository"
	"CocktailsApp/service"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	database := db.InitDB()
	defer database.Close()

	cocktailRepo := &repository.CocktailRepo{DB: database}
	cocktailService := &service.CocktailService{Repo: cocktailRepo}
	cocktailHandler := &handlers.Handler{Service: cocktailService}

	inventoryRepo := &repository.InventoryRepo{DB: database}
	inventoryService := &service.InventoryService{Repo: inventoryRepo}
	inventoryHandler := &handlers.InventoryHandler{Service: inventoryService}

	os.MkdirAll("uploads", os.ModePerm)
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))

	http.HandleFunc("/", cocktailHandler.HomeHandler)
	http.HandleFunc("/add", cocktailHandler.AddCocktailHandler)
	http.HandleFunc("/edit", cocktailHandler.EditCocktailHandler)
	http.HandleFunc("/delete", cocktailHandler.DeleteCocktailHandler)

	http.HandleFunc("/inventory", inventoryHandler.InventoryHandler)
	http.HandleFunc("/inventory/add", inventoryHandler.AddItemHandler)
	http.HandleFunc("/inventory/toggle", inventoryHandler.ToggleHandler)
	http.HandleFunc("/inventory/delete", inventoryHandler.DeleteItemHandler)
	http.HandleFunc("/inventory/toggle-all", inventoryHandler.ToggleAllItemsHandler)

	port := ":8080"
	fmt.Printf("[B.T.C TERMINAL] Архитектура собрана успешно. Ядро запущено на порту %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Ошибка при запуске веб-сервера: ", err)
	}
}
