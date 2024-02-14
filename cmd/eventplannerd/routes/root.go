package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"goproject.com/eventplanner-io/api/cmd/eventplannerd/routes/handlers"
	"goproject.com/eventplanner-io/api/cmd/eventplannerd/routes/middlewares"
	"goproject.com/eventplanner-io/api/cmd/eventplannerd/routes/middlewares/newrelic"
	"goproject.com/eventplanner-io/api/internal/repositories"
	"goproject.com/eventplanner-io/api/internal/services"
	"gorm.io/gorm"
)

func SetupRoutes(server *gin.Engine, db *gorm.DB) {

	// API GROUP
	authenticated := server.Group("/")

	// MIDDLEWARES
	app, _ := newrelic.NewrelicConfig()
	server.Use(nrgin.Middleware(app))
	authenticated.Use(middlewares.Authenticate)

	// REPOSITORIES
	eventsRepository := repositories.NewEventsRepository(db)
	registrationsRepository := repositories.NewRegistrationsRepository(db)
	usersRepository := repositories.NewUsersRepository(db)

	// USE CASES
	eventsService := services.NewEventsService(eventsRepository)
	registrationsService := services.NewRegistrationsService(registrationsRepository)
	usersService := services.NewUsersService(usersRepository)

	// HANDLERS
	eventsHandler := handlers.NewEventsHandler(eventsService, usersService)
	registrationsHandler := handlers.NewRegistrationsHandler(registrationsService, eventsService, usersService)
	usersHandler := handlers.NewUsersHandler(usersService)

	// ENDPOINTS
	server.GET("/events", eventsHandler.GetEvents)
	server.GET("/events/:id", eventsHandler.GetEventByID)

	authenticated.POST("/events", eventsHandler.CreateEvent)
	authenticated.PUT("/events/:id", eventsHandler.UpdateEvent)
	authenticated.DELETE("/events/:id", eventsHandler.DeleteEvent)
	authenticated.POST("/events/:id/register", registrationsHandler.RegisterForEvent)
	authenticated.DELETE("/events/:id/register", registrationsHandler.CancelRegistration)

	server.POST("/signup", usersHandler.Signup)
	server.POST("/login", usersHandler.Login)
}
