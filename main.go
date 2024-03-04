package main

import (
	"kevinhend97/belajar-golang-restful-api/app"
	"kevinhend97/belajar-golang-restful-api/controller"
	"kevinhend97/belajar-golang-restful-api/helper"
	"kevinhend97/belajar-golang-restful-api/middleware"
	"kevinhend97/belajar-golang-restful-api/repositories"
	"kevinhend97/belajar-golang-restful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repositories.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr: "localhost:3000",
		// Handler: router,  // Jika Tidak Menggunakan  middleware
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()

	helper.PanicIfError(err)
}
