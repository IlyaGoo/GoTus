package web_presenter

import (
	"github.com/gofiber/fiber/v2"
)

type WebPresenter struct {
	Port string
}

func (p *WebPresenter) StartWebPresenter() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello on GoTus")
	})

	app.Listen(":" + p.Port)
}
