package main

import (
	"awesomeProject/accounts"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	accountsHandler := accounts.New()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/account/create", accountsHandler.CreateAccount)
	e.GET("/account", accountsHandler.GetAccount)
	e.DELETE("/account", accountsHandler.DeleteAccount)
	e.PATCH("/account/name", accountsHandler.ChangeAccountName)
	e.PATCH("/account/balance", accountsHandler.ChangeAccountBalance)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
