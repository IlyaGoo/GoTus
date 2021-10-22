package web_presenter

import (
	"github.com/gofiber/fiber/v2"
)

type WebPresenter struct {
	Port    string
	Address string
}

func (p *WebPresenter) StartWebPresenter() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello on GoTus")
	})

	p.Address = "localhost:" + p.Port

	app.Listen(p.Address)
}
