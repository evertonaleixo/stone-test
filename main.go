package main

import (
	"github.com/evertonaleixo/stone-test/api"
	"github.com/evertonaleixo/stone-test/middlewares"
	"github.com/evertonaleixo/stone-test/services"
	"github.com/gin-gonic/gin"
)

func main() {
	services.ConnectDatabase()

	r := gin.Default()

	r.GET("/accounts", api.GetAcconts)
	r.GET("/accounts/:id", api.GetAccont)
	r.GET("/accounts/:id/balance", api.GetAccontBalance)
	r.POST("/accounts", api.CreateAccount)

	r.GET("/transfers", middlewares.TokenAuthMiddleware(), api.GetTransfers)
	r.POST("/transfers", middlewares.TokenAuthMiddleware(), api.DoTransfer)

	r.POST("/login", api.Login)
	r.POST("/logout", middlewares.TokenAuthMiddleware(), api.Logout)

	r.Run("0.0.0.0:8080")
}
