package service

import (
	"dcat/models"
)

func Login(name, pwd string) *models.AccountInfo {
	account := models.NewAccountInfo()
	account.FindByAccountName(name)
	return account
}
