package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/installable-bot-concept/pkg/templates"
	"github.com/opoccomaxao/installable-bot-concept/pkg/utils/ginutils"
	"github.com/opoccomaxao/installable-bot-concept/pkg/utils/hx"
)

type AuthRequest struct {
	Password string `form:"password" binding:"required,min=8,max=72"`
}

func (s *Service) Auth(ctx *gin.Context) {
	var req AuthRequest

	err := ctx.ShouldBind(&req)
	if err != nil {
		ginutils.RenderTempl(ctx, http.StatusBadRequest, templates.AuthPage(templates.AuthData{
			Path: "/auth",
			Errors: []string{
				"Invalid password",
			},
		}))

		return
	}

	err = s.auth.ValidateAdminPassword(ctx.Request.Context(), req.Password)
	if err != nil {
		ginutils.RenderTempl(ctx, http.StatusUnauthorized, templates.AuthPage(templates.AuthData{
			Path: "/auth",
			Errors: []string{
				"Invalid password",
			},
		}))

		return
	}

	err = s.auth.SetAuth(ctx)
	if err != nil {
		ginutils.RenderTempl(ctx, http.StatusUnauthorized, templates.AuthPage(templates.AuthData{
			Path: "/auth",
			Errors: []string{
				"Internal Server Error",
			},
		}))

		return
	}

	hx.Redirect(ctx, "/dashboard")
}
