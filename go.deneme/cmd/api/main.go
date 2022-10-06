package main

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"go.deneme/internal/database"
	"go.deneme/internal/handlers"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "berkay.yanik"
	dbname = "calhounio_demo"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func main() {
	db, err := database.SetupPsql(host, port, user, dbname)
	if err != nil {
		log.Fatal(err)
	}

	err = database.RunMigration(db)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	productHandler := handlers.NewProductHandler(validate, db)

	clothHandler := handlers.NewClothHandler(validate, db)

	e.POST("/product", productHandler.CreateProduct)

	e.GET("/product", productHandler.GetProduct)

	e.PUT("/product/:id", productHandler.UpdateProduct)

	e.DELETE("/product/:id", productHandler.DeleteProduct)

	e.POST("/cloth", clothHandler.CreateCloth)

	e.GET("/cloth", clothHandler.GetCloth)

	e.PUT("/cloth/:id", clothHandler.UpdateCloth)

	e.DELETE("/cloth/:id", clothHandler.DeleteCloth)

	e.Logger.Fatal(e.Start(":1323"))

}
