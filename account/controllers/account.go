package controllers

import (
	"account/services"

	"github.com/gofiber/fiber/v2"
)

type AccountController interface {
	GetAccounts(c *fiber.Ctx) error
}

type accountController struct {
	accountService services.AccountService
}

func NewAccountController(accountService services.AccountService) AccountController {
	return &accountController{accountService}
}

func (obj *accountController) GetAccounts(c *fiber.Ctx) error {

	bankAccounts, err := obj.accountService.GetAll()

	if err != nil {
		c.Status(fiber.StatusBadGateway)
		c.JSON(fiber.Map{
			"messsage": "GetAccounts Failed",
		})

		return err
	}

	c.Status(fiber.StatusCreated)
	c.JSON(bankAccounts)

	return nil
}
