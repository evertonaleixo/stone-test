package api

import (
	"net/http"
	"strconv"

	"github.com/evertonaleixo/stone-test/models"
	"github.com/evertonaleixo/stone-test/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTransfers(c *gin.Context) {
	var transfers []models.Transfer

	us, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	services.DB.Where("accountoriginid = ?", us.(models.Account).ID).Find(&transfers)
	c.JSON(200, gin.H{"transfers": transfers})
}

func truncate(some float32) float32 {
	return float32(int(some*100)) / 100
}

func DoTransfer(c *gin.Context) {
	// VALIDATIONS
	var input models.DoTransferInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	us := user.(models.Account)

	input.Amount = truncate(input.Amount)
	if input.Amount <= 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "the transference need to be positive"})
		return
	}

	if us.Balance < input.Amount {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "insuficient credit"})
		return
	}
	destinationAccount, err := GetAccountById(input.AccountDestinationId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "destination account does not exist"})
		return
	}

	if us.ID == destinationAccount.ID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you cannot transfer to yourself"})
		return
	}
	// END-VALIDATIONS
	var transfer models.Transfer

	transfer.AccountDestinationId = input.AccountDestinationId
	transfer.AccountOriginId = strconv.Itoa(int(us.ID))
	transfer.Amount = input.Amount

	err = services.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&destinationAccount).Update("balance", truncate(destinationAccount.Balance+input.Amount)).Error; err != nil {
			// return any error will rollback
			return err
		}

		if err := tx.Model(&us).Update("balance", truncate(us.Balance-input.Amount)).Error; err != nil {
			// return any error will rollback
			return err
		}

		if err := tx.Create(&transfer).Error; err != nil {
			// return any error will rollback
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "transaction fails"})
		return
	}

	c.JSON(200, gin.H{"transfer": transfer})
}
