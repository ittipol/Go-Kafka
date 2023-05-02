package controllers

import (
	"log"
	"producer/commands"
	"producer/services"

	"github.com/gofiber/fiber/v2"
)

type AccountController interface {
	OpenAccount(c *fiber.Ctx) error
	DepositFund(c *fiber.Ctx) error
	WithdrawFund(c *fiber.Ctx) error
	CloseAccount(c *fiber.Ctx) error
}

type accountController struct {
	accountService services.AccountService
}

func NewAccountController(accountService services.AccountService) AccountController {
	return &accountController{accountService}
}

func (obj *accountController) OpenAccount(c *fiber.Ctx) error {

	command := commands.OpenAccountCommand{}

	err := c.BodyParser(&command)

	if err != nil {
		log.Printf("Error: %v \n", err)

		c.Status(fiber.StatusBadGateway)
		c.JSON(fiber.Map{
			"messsage": "OpenAccount Failed",
		})

		return err
	}

	id, err := obj.accountService.OpenAccount(command)

	if err != nil {
		log.Printf("Error: %v \n", err)

		c.Status(fiber.StatusBadGateway)
		c.JSON(fiber.Map{
			"messsage": "OpenAccount Failed",
		})

		return err
	}

	log.Printf("Return ID: %v \n", id)

	c.Status(fiber.StatusCreated)
	c.JSON(fiber.Map{
		"messsage": "OpenAccount Success",
	})

	return nil
}

func (obj *accountController) DepositFund(c *fiber.Ctx) error {

	command := commands.DepositFundCommand{}

	err := c.BodyParser(&command)

	if err != nil {
		log.Printf("Error: %v \n", err)

		c.Status(fiber.StatusBadGateway)
		c.JSON(fiber.Map{
			"messsage": "DepositFund Failed",
		})

		return err
	}

	err = obj.accountService.DepositFund(command)

	if err != nil {
		log.Printf("Error: %v \n", err)

		c.Status(fiber.StatusBadGateway)
		c.JSON(fiber.Map{
			"messsage": "DepositFund Failed",
		})

		return err
	}

	c.Status(fiber.StatusOK)
	c.JSON(fiber.Map{
		"messsage": "DepositFund Success",
	})

	return nil
}

func (obj *accountController) WithdrawFund(c *fiber.Ctx) error {

	command := commands.WithdrawFundCommnd{}

	err := c.BodyParser(&command)

	if err != nil {
		log.Printf("Error: %v \n", err)

		c.Status(fiber.StatusBadGateway)
		c.JSON(fiber.Map{
			"messsage": "WithdrawFund Failed",
		})

		return err
	}

	err = obj.accountService.WithdrawFund(command)

	if err != nil {
		log.Printf("Error: %v \n", err)

		c.Status(fiber.StatusBadGateway)
		c.JSON(fiber.Map{
			"messsage": "WithdrawFund Failed",
		})

		return err
	}

	c.Status(fiber.StatusOK)
	c.JSON(fiber.Map{
		"messsage": "WithdrawFund Success",
	})

	return nil
}

func (obj *accountController) CloseAccount(c *fiber.Ctx) error {

	command := commands.CloseAccountCommand{}

	err := c.BodyParser(&command)

	if err != nil {
		log.Printf("Error: %v \n", err)

		c.Status(fiber.StatusBadGateway)
		c.JSON(fiber.Map{
			"messsage": "CloseAccount Failed",
		})

		return err
	}

	err = obj.accountService.CloseAccount(command)

	if err != nil {
		log.Printf("Error: %v \n", err)

		c.Status(fiber.StatusBadGateway)
		c.JSON(fiber.Map{
			"messsage": "CloseAccount Failed",
		})

		return err
	}

	c.Status(fiber.StatusOK)
	c.JSON(fiber.Map{
		"messsage": "CloseAccount Success",
	})

	return nil
}
