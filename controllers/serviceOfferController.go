package controllers
 
import (
	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/backend/common"
	"github.com/kgermando/backend/models"
)


func CreateServiceOffer(c *fiber.Ctx) error {
	return common.CreateResource(c, &models.ServiceOffer{})
}

func GetServiceOffers(c *fiber.Ctx) error {
	return common.GetResources(c, &[]models.ServiceOffer{})
}

func GetServiceOffer(c *fiber.Ctx) error {
	return common.GetResource(c, &models.ServiceOffer{})
}

func UpdateServiceOffer(c *fiber.Ctx) error {
	
	return common.UpdateResource(c, &models.ServiceOffer{})
}

func DeleteServiceOffer(c *fiber.Ctx) error {
	return common.DeleteResource(c, &models.ServiceOffer{})
}



