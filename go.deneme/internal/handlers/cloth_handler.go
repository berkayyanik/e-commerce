package handlers

import (
	"database/sql"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.deneme/internal/models"
)

type ClothHandler struct {
	validator *validator.Validate
	db        *sql.DB
}

func (h ClothHandler) CreateCloth(c echo.Context) error {

	cloth := models.CreateCloth{}
	err := c.Bind(&cloth)
	if err != nil {
		return err
	}

	err = h.validator.Struct(cloth)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	sqlStatement := `INSERT INTO clothes (Name, Description, BrandName, Price, Stock, Size)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = h.db.Exec(sqlStatement, cloth.Name, cloth.Description, cloth.BrandName, cloth.Price, cloth.Stock, cloth.Size)

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func (h ClothHandler) GetCloth(c echo.Context) error {

	rows, err := h.db.Query("SELECT * FROM clothes")
	if err != nil {
		return err
	}

	clothes := []models.CreateCloth{}
	for rows.Next() {
		cloth := models.CreateCloth{}
		err := rows.Scan(&cloth.ID, &cloth.Name, &cloth.Description, &cloth.BrandName, &cloth.Price, &cloth.Stock, &cloth.Size)
		if err != nil {
			return err
		}

		clothes = append(clothes, cloth)
	}

	return c.JSON(http.StatusOK, clothes)
}

func (h ClothHandler) UpdateCloth(c echo.Context) error {

	id := c.Param("id")
	if id == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	updateCloth := models.UpdateCloth{}
	err := c.Bind(&updateCloth)
	if err != nil {
		return err
	}

	err = h.validator.Struct(updateCloth)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	sqlStatementUpdt := `UPDATE products SET Name = $1, Description = $2, BrandName = $3, Price = $4, Stock = $5, Size = $6 WHERE ID = $6;`
	_, err = h.db.Exec(sqlStatementUpdt, updateCloth.Name, updateCloth.Description, updateCloth.BrandName, updateCloth.Price, updateCloth.Stock, updateCloth.Size, updateCloth.ID)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func (h ClothHandler) DeleteCloth(c echo.Context) error {

	id := c.Param("id")
	if id == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	sqlStatementDel := `DELETE FROM clothes WHERE ID = $1;`
	_, err := h.db.Exec(sqlStatementDel, id)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusCreated)
}

func NewClothHandler(validator *validator.Validate, db *sql.DB) ClothHandler {
	return ClothHandler{
		validator: validator,
		db:        db,
	}

}
