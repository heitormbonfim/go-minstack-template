package main

import (
	"go-minstack/internal/users"
	user_entities "go-minstack/internal/users/entities"
	"go-minstack/internal/users/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-minstack/core"
	mgin "github.com/go-minstack/gin"
	"github.com/go-minstack/logger"
	"github.com/go-minstack/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := core.New(mgin.Module(), sqlite.Module(), logger.Module())

	// users domain
	app.Provide(repositories.NewUserRepository)
	app.Provide(users.NewUserService)
	app.Provide(users.NewUserController)
	app.Invoke(users.RegisterRoutes)
	app.Invoke(migrate)

	// healthcheck
	app.Invoke(pingRoutes)

	app.Run()
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(&user_entities.User{})
}

func pingRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Pong"})
	})
}
