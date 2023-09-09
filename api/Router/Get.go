package router

import (
	"fmt"
	"net/http"
    Utils "wapi/Utils"
	"github.com/gofiber/fiber/v2"
)

func Api(c *fiber.Ctx) error {
	sehir := c.Params("sehir")
	if sehir == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": 404,
		})
	}

	weatherData, err := Utils.Get(sehir)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Hava durumu alınamadı: %v", err),
		})
	}

	return c.JSON(weatherData)
}