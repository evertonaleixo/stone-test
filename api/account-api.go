package api

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"strings"
	"time"

	"github.com/evertonaleixo/stone-test/models"
	"github.com/evertonaleixo/stone-test/services"
	"github.com/gin-gonic/gin"
)

func GetAcconts(c *gin.Context) {
	var accounts []models.Account
	services.DB.Find(&accounts)

	c.JSON(200, gin.H{"accounts": accounts})
}

func GetAccountById(id string) (models.Account, error) {
	var account models.Account

	if err := services.DB.Where("id = ?", id).First(&account).Error; err != nil {
		return models.Account{}, err
	}

	return account, nil
}

func GetAccont(c *gin.Context) {
	var account, err = GetAccountById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"account": account})
}

func GetAccontBalance(c *gin.Context) {
	var account, err = GetAccountById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": account.Balance})
}

func CreateAccount(c *gin.Context) {
	// Validate input
	var input models.CreateAccountInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create account
	var secret string = input.Secret
	hash := md5.Sum([]byte(secret))
	secret = hex.EncodeToString(hash[:])
	var cpf string = strings.Replace(input.Cpf, ".", "", -1)
	cpf = strings.Replace(cpf, "-", "", -1)

	account := models.Account{
		Name:    input.Name,
		Cpf:     cpf,
		Secret:  secret,
		Balance: 1000,
	}
	account.CreatedAt = time.Now()
	account.UpdatedAt = account.CreatedAt

	if insertErr := services.DB.Create(&account).Error; insertErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": insertErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"account": account})
}
