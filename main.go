package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/aluno", func(c *fiber.Ctx) error {
		return c.SendString("Lista de alunos:")
	})

	app.Get("/aluno/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("Aluno id =" + id)
	})

	app.Post("/aluno", func(c *fiber.Ctx) error {
		json := string(c.Request().Body())
		return c.SendString(("Inserir aluno =" + json))
	})

	app.Put("/aluno/ :id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		json := string(c.Request().Body())
		return c.SendString("Alterar aluno id= " + id + "para" + json)
	})

	app.Delete("/aluno/ :id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString("Excluir aluno id =" + id)
	})

	app.Listen(" :8080")

}
