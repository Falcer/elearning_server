package auth

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

var (
	databaseURL string
	r           Repository
	s           Service
)

func NewRouter(app *fiber.App) *fiber.App {
	// Get DATABASE_URL
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = ""
	}
	// Set UserRePository
	r = NewRepo(databaseURL)
	// Set UserService
	s = NewService(r)
	// Health Check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(&fiber.Map{
			"health": "UP",
		})
	})
	app.Post("/login", func(c *fiber.Ctx) error {
		var body = new(Login)
		if err := c.BodyParser(body); err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"message": "Something went wrong",
			})
		}
		res, err := s.Login(*body)
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"message": err.Error(),
			})
		}

		cookie := new(fiber.Cookie)
		cookie.Name = "refresh_token"
		cookie.Value = (*res).RefreshToken
		c.Cookie(cookie)

		return c.Status(200).JSON(&fiber.Map{
			"message": "Login successfully",
			"data":    *res,
		})

	})
	app.Post("/register", func(c *fiber.Ctx) error {
		var body = new(Register)
		if err := c.BodyParser(body); err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"message": "Something went wrong",
			})
		}
		res, err := s.Register(*body)
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"message": err.Error(),
			})
		}

		cookie := new(fiber.Cookie)
		cookie.Name = "refresh_token"
		cookie.Value = (*res).RefreshToken
		c.Cookie(cookie)

		return c.Status(200).JSON(&fiber.Map{
			"message": "Register successfully",
			"data":    *res,
		})

	})
	app.Post("/user/role", func(c *fiber.Ctx) error {
		userrole := new(UserRoleInput)
		if err := c.BodyParser(userrole); err != nil {
			return err
		}
		return c.Status(200).JSON(&fiber.Map{
			"message": "User role added is running 🔥",
			"data":    &userrole,
		})
	})
	app.Delete("/user/role", func(c *fiber.Ctx) error {
		userrole := new(UserRoleInput)
		if err := c.BodyParser(userrole); err != nil {
			return err
		}
		return c.Status(200).JSON(&fiber.Map{
			"message": "User role deleted is running 🔥",
			"data":    &userrole,
		})
	})

	// Role
	app.Get("/role", func(c *fiber.Ctx) error {
		res, err := s.GetRoles()
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"message": fmt.Sprintf("Something wrong : %s", err.Error()),
			})
		}
		return c.Status(200).JSON(&fiber.Map{
			"message": "Get all Role",
			"data":    *res,
		})
	})
	app.Get("/role/:id", func(c *fiber.Ctx) error {
		res, err := s.GetRoleByID(c.Params("id"))
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"message": fmt.Sprintf("Something wrong : %s", err.Error()),
			})
		}
		return c.Status(200).JSON(&fiber.Map{
			"message": fmt.Sprintf("Role id : %s", c.Params("id")),
			"data":    *res,
		})
	})
	app.Post("/role", func(c *fiber.Ctx) error {
		role := new(RoleInput)
		if err := c.BodyParser(role); err != nil {
			return err
		}
		res, err := s.AddRole(*role)
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"message": fmt.Sprintf("Something wrong : %s", err.Error()),
			})
		}
		return c.Status(200).JSON(&fiber.Map{
			"message": "Role created",
			"data":    *res,
		})
	})
	app.Put("/role", func(c *fiber.Ctx) error {
		role := new(RoleOutput)
		if err := c.BodyParser(role); err != nil {
			return err
		}
		res, err := s.UpdateRole(*role)
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"message": fmt.Sprintf("Something wrong : %s", err.Error()),
			})
		}
		return c.Status(200).JSON(&fiber.Map{
			"message": "Role updated",
			"data":    *res,
		})
	})
	app.Delete("/role/:id", func(c *fiber.Ctx) error {
		err := s.DeleteRoleByID(c.Params("id"))
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"message": fmt.Sprintf("Something wrong : %s", err.Error()),
			})
		}
		return c.Status(200).JSON(&fiber.Map{
			"message": fmt.Sprintf("Role id : %s deleted", c.Params("id")),
		})
	})
	return app
}
