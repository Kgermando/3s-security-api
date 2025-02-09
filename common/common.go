package common

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/backend/services"
)

// CreateResource creates a new resource
func CreateResource(c *fiber.Ctx, resource interface{}) error {
	if err := c.BodyParser(resource); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := services.CreateResource(resource); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(resource)
}

// GetResources retrieves a paginated list of resources
func GetResources(c *fiber.Ctx, resources interface{}) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize", "10"))
	if err := services.GetResources(page, pageSize, resources); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(resources)
}

// GetResource retrieves a single resource by ID
func GetResource(c *fiber.Ctx, resource interface{}) error {
	id := c.Params("id")
	if err := services.GetResource(id, resource); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(resource)
}

// UpdateResource updates an existing resource by ID
func UpdateResource(c *fiber.Ctx, resource interface{}) error {
	id := c.Params("id")
	if err := c.BodyParser(resource); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := services.UpdateResource(id, resource); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(resource)
}

// DeleteResource deletes a resource by ID
func DeleteResource(c *fiber.Ctx, resource interface{}) error {
	id := c.Params("id")
	if err := services.DeleteResource(id, resource); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Resource deleted"})
}
