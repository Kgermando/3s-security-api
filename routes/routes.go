package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kgermando/backend/controllers/auth"
	"github.com/kgermando/backend/controllers/users"
	"github.com/kgermando/backend/controllers"
	"github.com/kgermando/backend/middlewares"
)

func Setup(app *fiber.App) {

	api := app.Group("/api", logger.New())

	// Authentification controller
	au := api.Group("/auth")
	au.Post("/register", auth.Register)
	au.Post("/login", auth.Login)
	au.Post("/forgot-password", auth.Forgot)
	au.Post("/reset/:token", auth.ResetPassword)
 

	// Routes for blog
	blog := api.Group("/blogs")
	blog.Get("/all", controllers.GetBlogs)
	blog.Get("/get/:id", controllers.GetBlog)
	blog.Post("/create", controllers.CreateBlog)
	blog.Put("/update/:id", controllers.UpdateBlog)
	blog.Delete("/delete/:id", controllers.DeleteBlog)

	// Routes for service-offer
	serviceOffer := api.Group("/service-offers")
	serviceOffer.Get("/all", controllers.GetServiceOffers)
	serviceOffer.Get("/get/:id", controllers.GetServiceOffer)
	serviceOffer.Post("/create", controllers.CreateServiceOffer)
	serviceOffer.Put("/update/:id", controllers.UpdateServiceOffer)
	serviceOffer.Delete("/delete/:id", controllers.DeleteServiceOffer)

	// Routes for SliderHome
	sliderHome := api.Group("/slider-homes")
	sliderHome.Get("/all", controllers.GetSliderHomes)
	sliderHome.Get("/get/:id", controllers.GetSliderHome)
	sliderHome.Post("/create", controllers.CreateSliderHome)
	sliderHome.Put("/update/:id", controllers.UpdateSliderHome)
	sliderHome.Delete("/delete/:id", controllers.DeleteSliderHome)

	// Routes for Team
	team := api.Group("/teams")
	team.Get("/all", controllers.GetTeams)
	team.Get("/get/:id", controllers.GetTeam)
	team.Post("/create", controllers.CreateTeam)
	team.Put("/update/:id", controllers.UpdateTeam)
	team.Delete("/delete/:id", controllers.DeleteTeam)


	// Secure routes
	app.Use(middlewares.IsAuthenticated)

	au.Get("/user", auth.AuthUser)
	au.Put("/profil/info", auth.UpdateInfo)
	au.Put("/change-password", auth.ChangePassword)
	au.Post("/logout", auth.Logout)

	// User controller
	u := api.Group("/users")
	u.Get("/all", users.GetAllUsers)
	u.Get("/all/paginate", users.GetPaginatedUsers)
	u.Get("/get/:id", users.GetUser)
	u.Post("/create", users.CreateUser)
	u.Put("/update/:id", users.UpdateUser)
	u.Delete("/delete/:id", users.DeleteUser)
 

}
