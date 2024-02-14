package eventplannerd

import (
	"github.com/gin-gonic/gin"
	"goproject.com/eventplanner-io/api/cmd/eventplannerd/routes"
	"goproject.com/eventplanner-io/api/internal/platform/database/migrations"
	"goproject.com/eventplanner-io/api/internal/platform/database/sqlite"
)

func Run() {
	// Agregar una estructura de DPI
	db, err := sqlite.Connect()
	if err != nil {
		panic(("Failed to connect to database"))
	}
	err = migrations.MigrateDB(db)
	if err != nil {
		panic(("Failed to use migrations."))
	}
	server := gin.Default()
	routes.SetupRoutes(server, db)
	server.Run(":8080")
}
