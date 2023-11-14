package controllers

import (
	"apiauth/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Show(c *fiber.Ctx) error {

	username := c.Params("username")

	var auth models.Auth

	if err := models.DB.Model(&auth).Where("username = ?", username).First(&auth).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Username tidak ditemukan.",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Server sedang mengalami gangguan.",
		})
	}

	return c.JSON(auth)
}

func Create(c *fiber.Ctx) error {

	var auth models.Auth
	if err := c.BodyParser(&auth); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := models.DB.Create(&auth).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "User berhasil dibuat",
	})
}

func Update(c *fiber.Ctx) error {

	username := c.Params("username")

	var auth models.Auth

	if err := c.BodyParser(&auth); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if models.DB.Where("username = ?", username).Updates(&auth).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Username tidak dapat ditemukan.",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil diupdate.",
	})
}

func Delete(c *fiber.Ctx) error {

	username := c.Params("username")
	var auth models.Auth

	if models.DB.Where("username = ?", username).Delete(&auth).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "USername tidak dapat dihapus.",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil menghapus user.",
	})
}