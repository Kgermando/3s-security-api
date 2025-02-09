package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/backend/common"
	"github.com/kgermando/backend/models"
)


func CreateSliderHome(c *fiber.Ctx) error {
	return common.CreateResource(c, &models.SliderHome{})
}

func GetSliderHomes(c *fiber.Ctx) error {
	return common.GetResources(c, &[]models.SliderHome{})
}

func GetSliderHome(c *fiber.Ctx) error {
	return common.GetResource(c, &models.SliderHome{})
}

func UpdateSliderHome(c *fiber.Ctx) error {
	
	return common.UpdateResource(c, &models.SliderHome{})
}

func DeleteSliderHome(c *fiber.Ctx) error {
	return common.DeleteResource(c, &models.SliderHome{})
}

