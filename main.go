package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)


type TODO struct {
    ID        int    `json:"id"`         // Unique identifier for the TODO
    Completed bool   `json:"completed"`  // Status of the TODO
    Body      string `json:"body"`       // Description of the TODO
}

func main(){
	app := fiber.New()

	todos := []TODO{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "hello world"})
	  })

	  app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &TODO{}

		if err := c.BodyParser(todo); err != nil{
			return err
		}

		if todo.Body == ""{
			return c.Status(400).JSON(fiber.Map{"error": "TODO body is required"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(200).JSON(todo)
	  })

	  app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos{
			if fmt.Sprint(todo.ID) == id{
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Todo not Found"})
	  })


	  app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos{
			if fmt.Sprint(todo.ID) == id{
				todos = append(todos[:i], todos[i+1:]... )
				return c.Status(200).JSON(fiber.Map{"success": true})
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Todo not Found"})
	  })

	  log.Fatal(app.Listen(":4000"))
}