package handlers

import (
	"database/sql"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.deneme/internal/models"
)

type ProductHandler struct {
	validator *validator.Validate
	db        *sql.DB
}

func (h ProductHandler) CreateProduct(c echo.Context) error {

	product := models.CreateProduct{}
	err := c.Bind(&product)
	if err != nil {
		return err
	}

	err = h.validator.Struct(product)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	sqlStatement := `INSERT INTO products (Name, Description, BrandName, Price, Stock)
		VALUES ($1, $2, $3, $4, $5)`
	_, err = h.db.Exec(sqlStatement, product.Name, product.Description, product.BrandName, product.Price, product.Stock)

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func (h ProductHandler) GetProduct(c echo.Context) error {

	rows, err := h.db.Query("SELECT * FROM products")
	if err != nil {
		return err
	}

	products := []models.CreateProduct{}
	for rows.Next() {
		product := models.CreateProduct{}
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.BrandName, &product.Price, &product.Stock)
		if err != nil {
			return err
		}

		products = append(products, product)
	}

	return c.JSON(http.StatusOK, products)
}

func (h ProductHandler) UpdateProduct(c echo.Context) error {

	id := c.Param("id")
	if id == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	updateProduct := models.UpdateProduct{}
	err := c.Bind(&updateProduct)
	if err != nil {
		return err
	}

	err = h.validator.Struct(updateProduct)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	sqlStatementUpdt := `UPDATE products SET Name = $1, Description = $2, BrandName = $3, Price = $4, Stock = $5 WHERE ID = $6;`
	_, err = h.db.Exec(sqlStatementUpdt, updateProduct.Name, updateProduct.Description, updateProduct.BrandName, updateProduct.Price, updateProduct.Stock, updateProduct.ID)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func (h ProductHandler) DeleteProduct(c echo.Context) error {

	id := c.Param("id")
	if id == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	sqlStatementDel := `DELETE FROM products WHERE ID = $1;`
	_, err := h.db.Exec(sqlStatementDel, id)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func NewProductHandler(validator *validator.Validate, db *sql.DB) ProductHandler {
	return ProductHandler{
		validator: validator,
		db:        db,
	}

}
