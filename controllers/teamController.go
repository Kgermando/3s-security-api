package controllers


import (
	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/backend/common"
	"github.com/kgermando/backend/models"
)


func CreateTeam(c *fiber.Ctx) error {
	return common.CreateResource(c, &models.Team{})
}

func GetTeams(c *fiber.Ctx) error {
	return common.GetResources(c, &[]models.Team{})
}

func GetTeam(c *fiber.Ctx) error {
	return common.GetResource(c, &models.Team{})
}

func UpdateTeam(c *fiber.Ctx) error {
	
	return common.UpdateResource(c, &models.Team{})
}

func DeleteTeam(c *fiber.Ctx) error {
	return common.DeleteResource(c, &models.Team{})
}
