package controllers

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
)

type Transaction struct {
	ProfileName string `json:"profileName"`
	ID          string `json:"id"`
}

type Request struct {
	Title string `json:"title"`
}

var sampleTransactions = []*Transaction{
	{
		ProfileName: "Sample001",
		ID:          "0001",
	},
	{
		ProfileName: "Sample002",
		ID:          "0002",
	},
}

func GetTransactions(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"transactions": sampleTransactions,
		},
	})
}

func CreatTransactions(c *fiber.Ctx) error {
	var body Request
	err := c.BodyParser(&body)
	// if error
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	transaction := &Transaction{
		ProfileName: "Sample003",
		ID:          "003",
	}

	sampleTransactions = append(sampleTransactions, transaction)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": transaction,
		},
	})
}
