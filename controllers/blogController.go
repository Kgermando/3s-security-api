package controllers
 
import (
	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/backend/common"
	"github.com/kgermando/backend/models"
)


func CreateBlog(c *fiber.Ctx) error {
	return common.CreateResource(c, &models.Blog{})
}

func GetBlogs(c *fiber.Ctx) error {
	return common.GetResources(c, &[]models.Blog{})
}

func GetBlog(c *fiber.Ctx) error {
	return common.GetResource(c, &models.Blog{})
}

func UpdateBlog(c *fiber.Ctx) error {
	
	return common.UpdateResource(c, &models.Blog{})
}

func DeleteBlog(c *fiber.Ctx) error {
	return common.DeleteResource(c, &models.Blog{})
}



