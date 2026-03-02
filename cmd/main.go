package main

import (
	"go-minstack-task/internal/authn"
	"go-minstack-task/internal/tasks"
	task_entities "go-minstack-task/internal/tasks/entities"
	"go-minstack-task/internal/users"
	user_entities "go-minstack-task/internal/users/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-minstack/auth"
	"github.com/go-minstack/core"
	mgin "github.com/go-minstack/gin"
	"github.com/go-minstack/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := core.New(mgin.Module(), sqlite.Module(), auth.Module())

	users.Register(app)
	authn.Register(app)
	tasks.Register(app)

	app.Invoke(migrate)
	app.Invoke(pingRoutes)
	app.Run()
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&user_entities.User{},
		&task_entities.Task{},
	)
}

func pingRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Pong"})
	})
}
