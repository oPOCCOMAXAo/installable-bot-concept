package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/installable-bot-concept/pkg/services/auth"
	"github.com/opoccomaxao/installable-bot-concept/pkg/templates"
	"github.com/opoccomaxao/installable-bot-concept/pkg/utils/ginutils"
	"github.com/samber/do"
)

func Init(
	router gin.IRouter,
	i *do.Injector,
) error {
	authService, err := auth.Invoke(i)
	if err != nil {
		return err
	}

	service := New(
		authService,
	)

	router.Use(authService.Middleware(auth.MiddlewareParams{
		SetAdminPasswordPath: "/init",
	}))

	router.GET("/", ginutils.StaticRedirect("/dashboard"))

	router.GET("/init", ginutils.StaticTempl(templates.InitPage("/init", nil)))
	router.POST("/init", service.Init)

	router.GET("/dashboard", ginutils.StaticTempl(templates.DashboardPage()))

	return nil
}
