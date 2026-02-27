package users

import (
	user_repositories "go-minstack-task/internal/users/repositories"

	"github.com/go-minstack/core"
)

func Register(app *core.App) {
	app.Provide(user_repositories.NewUserRepository)
	app.Provide(NewUserService)
	app.Provide(NewUserController)
	app.Invoke(RegisterRoutes)
}
