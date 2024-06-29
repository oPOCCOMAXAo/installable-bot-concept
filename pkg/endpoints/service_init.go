package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/installable-bot-concept/pkg/templates"
	"github.com/opoccomaxao/installable-bot-concept/pkg/utils/ginutils"
)

type InitRequest struct {
	Password string `form:"password" binding:"required,min=8,max=72"`
}

func (s *Service) Init(ctx *gin.Context) {
	var req InitRequest

	err := ctx.ShouldBind(&req)
	if err != nil {
		ginutils.RenderTempl(ctx, http.StatusBadRequest, templates.InitPage(
			"/init",
			[]string{
				"Password is required",
				"Password must be at least 8 characters long",
				"Password must be at most 72 characters long",
			},
		))

		return
	}

	err = s.auth.SetAdminPassword(ctx.Request.Context(), req.Password)
	if err != nil {
		ginutils.RenderTempl(ctx, http.StatusInternalServerError, templates.InitPage(
			"/init",
			[]string{
				"Failed to set admin password",
			},
		))

		return
	}

	ctx.Redirect(http.StatusFound, "/dashboard")
}
