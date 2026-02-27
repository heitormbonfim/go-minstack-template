package authn

import (
	authn_dto "go-minstack-task/internal/authn/dtos"
	user_dto "go-minstack-task/internal/users/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-minstack/web"
)

type AuthnController struct {
	service *AuthnService
}

func NewAuthnController(service *AuthnService) *AuthnController {
	return &AuthnController{service: service}
}

func (c *AuthnController) login(ctx *gin.Context) {
	var input user_dto.LoginDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewErrorDto(err))
		return
	}
	token, err := c.service.Login(input)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, web.NewErrorDto(err))
		return
	}
	ctx.JSON(http.StatusOK, authn_dto.TokenDto{Token: token})
}
