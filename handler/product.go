package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/brenddonanjos/echo_gorm/database"
	"github.com/brenddonanjos/echo_gorm/model"
	"github.com/labstack/echo/v4"
)

//GetProduct funtion to get all products from database
func GetProduct(c echo.Context) error {
	db := database.StartGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	products := model.Products{}

	if err := db.Find(&products).Error; err != nil {
		return echo.NewHTTPError(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, products)
}

//ShowProduct function to get specific product based on id
func ShowProduct(c echo.Context) error {
	db := database.StartGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	p := new(model.Product)
	id := c.Param("id")
	db.First(&p, id)
	return c.JSON(http.StatusOK, p)
}

// SetProduct function to save a new product on database
func SetProduct(c echo.Context) error {
	db := database.StartGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	p := new(model.Product)
	if err := c.Bind(&p); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	//create register on database
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	db.Omit("DeletedAt").Create(&p)

	return c.JSON(http.StatusCreated, p)

}

// UpdateProdutct function to update specific product
func UpdateProdutct(c echo.Context) error {
	db := database.StartGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	p := new(model.Product)
	id, _ := strconv.Atoi(c.Param("id"))
	db.First(&p, id)                   //find on db and set on p(Product)
	if err := c.Bind(&p); err != nil { //bind the body params
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	db.Omit("DeletedAt").Save(&p) //save new product values

	return c.JSON(http.StatusCreated, p)
}

//DeleteProduct function to inactivate a specific produtc on database
func DeleteProduct(c echo.Context) error {
	db := database.StartGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	p := new(model.Product)
	id, _ := strconv.Atoi(c.Param("id"))

	db.Delete(&p, id)

	return c.String(http.StatusOK, "Produto: "+p.Name+" inativado")
}
