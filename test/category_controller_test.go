package test

import (
	"kevinhend97/belajar-golang-restful-api/app"
	"kevinhend97/belajar-golang-restful-api/controller"
	"kevinhend97/belajar-golang-restful-api/middleware"
	"kevinhend97/belajar-golang-restful-api/repositories"
	"kevinhend97/belajar-golang-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func setupRouter() http.Handler {
	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repositories.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}
