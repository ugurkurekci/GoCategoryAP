package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/ugurkurekci/GoCategoryAPI/docs"

	"github.com/ugurkurekci/GoCategoryAPI/controller"
	"github.com/ugurkurekci/GoCategoryAPI/repository"
	"github.com/ugurkurekci/GoCategoryAPI/service"
)

func main() {
	connString := "server=UGURKUREKCI\\SQLEXPRESS;database=CategoryAPI;trusted_connection=true"

	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close()

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryController := controller.NewCategoryController(categoryService)

	http.HandleFunc("/categories", categoryController.CreateCategory)
	http.HandleFunc("/categories/{id}", categoryController.GetCategory)
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
