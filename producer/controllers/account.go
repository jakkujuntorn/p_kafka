package controllers

import (
	"producer/commands"
	"producer/services"

	fiber "github.com/gofiber/fiber/v2"
)

type I_AccountController interface {
	OpenAccount(c *fiber.Ctx) error
	DepositFund(c *fiber.Ctx) error
	WithdrawFund(c *fiber.Ctx) error
	CloseAccount(c *fiber.Ctx) error
}

type accountController struct {
	accountService services.I_AccountService
}

func NewAccountController(accountService services.I_AccountService) I_AccountController {
	return accountController{
		accountService: accountService,
	}
}

// CloseAccount implements AccountController
func (obj accountController) CloseAccount(c *fiber.Ctx) error {

	closeAccount := commands.CloseAccount_Command{}
	err := c.BodyParser(&closeAccount)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	err = obj.accountService.CloseAccount(closeAccount)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "CloseAccount Success",
	})

}

// DepositFund implements AccountController
func (obj accountController) DepositFund(c *fiber.Ctx) error {

	depositAccount := commands.DepositFund_Command{}
	err := c.BodyParser(&depositAccount)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	err = obj.accountService.DeposiFund(depositAccount)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Deposi Success",
	})

}

// OpenAccount implements AccountController
func (obj accountController) OpenAccount(c *fiber.Ctx) error {

	openAccount := commands.OpenAccount_Command{}
	err := c.BodyParser(&openAccount)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	id, err := obj.accountService.OpenAccount(openAccount)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err,
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"message": "Create success",
		"id":      id,
	})
}

// WithdrawFund implements AccountController
func (obj accountController) WithdrawFund(c *fiber.Ctx) error {

	withDrawAccount := commands.WithdrawFund_Command{}
	err := c.BodyParser(&withDrawAccount)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	err = obj.accountService.WithdrawFun(withDrawAccount)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Witdraw Success",
	})

}
