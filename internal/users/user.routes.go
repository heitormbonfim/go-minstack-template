package users

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, c *UserController) {
	g := r.Group("/api/users")
	g.GET("/", c.list)
	g.GET("/:id", c.get)
	g.POST("/", c.create)
	g.PUT("/:id", c.update)
	g.DELETE("/:id", c.delete)
}
