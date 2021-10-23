package web_presenter

import (
	"github.com/gofiber/fiber/v2"
)

type WebPresenter struct {
	port string
}

func NewWebPresenter(port string) WebPresenter {
	return WebPresenter{
		port: port,
	}
}

func (p *WebPresenter) StartWebPresenter() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello on GoTus")
	})

	app.Listen(":" + p.port)
}
