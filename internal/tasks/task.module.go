package tasks

import (
	task_repositories "go-minstack-task/internal/tasks/repositories"

	"github.com/go-minstack/core"
)

func Register(app *core.App) {
	app.Provide(task_repositories.NewTaskRepository)
	app.Provide(NewTaskService)
	app.Provide(NewTaskController)
	app.Invoke(RegisterRoutes)
}
