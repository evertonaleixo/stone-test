package api

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"strings"

	"github.com/evertonaleixo/stone-test/models"
	"github.com/evertonaleixo/stone-test/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var accountCredentials models.AuthCredentialsInput

	if err := c.ShouldBindJSON(&accountCredentials); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	var account models.Account
	var cpf string = strings.Replace(accountCredentials.Cpf, ".", "", -1)
	cpf = strings.Replace(cpf, "-", "", -1)
	var secret string = accountCredentials.Secret

	hash := md5.Sum([]byte(secret))
	secret = hex.EncodeToString(hash[:])

	if err := services.DB.Where("cpf = ?", cpf).First(&account).Error; err != nil {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}

	if account.Secret != secret {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details 222")
		return
	}

	ts, err := services.CreateToken(uint64(account.ID))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	saveErr := services.CreateAuth(uint64(account.ID), ts)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
		return
	}
	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}
	c.JSON(http.StatusOK, tokens)
}

func Logout(c *gin.Context) {
	au, err := services.ExtractTokenMetadata(c.Request)

	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	deleted, delErr := services.DeleteAuth(au.AccessUuid)
	if delErr != nil || deleted == 0 { //if any goes wrong
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}
