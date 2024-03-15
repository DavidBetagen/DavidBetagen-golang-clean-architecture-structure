package route

import (
	"golang-clean-architecture-structure/internal/infrastructure/helper"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func defaultRoute(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "✅✅✅✅",
		})
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "✅ I'm still alive. 😎",
		})
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Pong! 🏓 👻👻👻",
		})
	})

	app.Get("/app.log", func(c *fiber.Ctx) error {
		content, err := os.ReadFile("log/app.log")
		if err != nil {
			return c.Status(500).SendString("Error reading the file")
		}

		return c.SendString(string(content))
	})

	app.Get("/logger.log", func(c *fiber.Ctx) error {
		content, err := os.ReadFile("log/logger.log")
		if err != nil {
			return c.Status(500).SendString("Error reading the file")
		}

		return c.SendString(string(content))
	})

	app.Get("/metrics", monitor.New())

	app.Get("/download-log", func(c *fiber.Ctx) error {
		// Path to the folder you want to download
		folderPath := "log"
		// Path to the zip file to be created
		zipFileName := "log.zip"

		// Create a zip file
		if err := helper.ZipFolder(folderPath, zipFileName); err != nil {
			return err
		}

		// Read the zip file
		data, err := os.ReadFile(zipFileName)
		if err != nil {
			return err
		}

		// Delete the zip file after reading it
		defer os.Remove(zipFileName)

		// Set response headers
		c.Set(fiber.HeaderContentType, "application/zip")
		c.Set(fiber.HeaderContentDisposition, "attachment; filename="+zipFileName)

		// Send the zip file as response
		return c.Send(data)
	})

}
