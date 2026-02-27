package authn

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, c *AuthnController) {
	g := r.Group("/api/auth")
	g.POST("/login", c.login)
}
