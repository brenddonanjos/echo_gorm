package main

import (
	"net/http"

	"github.com/brenddonanjos/echo_gorm/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	//echo framework instance
	e := echo.New()

	// Middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{ //enable cors
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	//Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Olá mundo !")
	})
	e.GET("/products", handler.GetProduct)
	e.GET("/products/:id", handler.ShowProduct)
	e.POST("/products", handler.SetProduct)
	e.PUT("/products/:id", handler.UpdateProdutct)
	e.DELETE("/products/:id", handler.DeleteProduct)

	e.Logger.Fatal(e.Start(":3000"))
}
