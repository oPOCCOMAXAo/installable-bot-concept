package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/installable-bot-concept/assets"
	"github.com/opoccomaxao/installable-bot-concept/pkg/services/auth"
	"github.com/opoccomaxao/installable-bot-concept/pkg/services/demo"
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

	demoService, err := demo.Invoke(i)
	if err != nil {
		return err
	}

	service := New(
		authService,
		demoService,
	)

	unauthedGroup := router.Group("/")

	unauthedGroup.StaticFS("/assets", http.FS(&assets.FS))

	unauthedGroup.GET("/auth", ginutils.StaticTempl(templates.AuthPage(templates.AuthData{
		Path: "/auth",
	})))

	unauthedGroup.POST("/auth", service.Auth)
	unauthedGroup.POST("/logout", service.Logout)
	unauthedGroup.POST("/expire", service.Expire)

	authedGroup := router.Group("/",
		authService.Middleware(auth.MiddlewareParams{
			SetAdminPasswordPath: "/init",
			AuthPath:             "/auth",
		}),
	)

	authedGroup.GET("/", ginutils.StaticRedirect("/dashboard"))

	authedGroup.GET("/init", ginutils.StaticTempl(templates.InitPage(templates.InitData{
		Path: "/init",
	})))
	authedGroup.POST("/init", service.Init)

	authedGroup.GET("/dashboard", service.Dashboard)
	authedGroup.POST("/increment", service.Increment)
	authedGroup.POST("/key", service.Key)

	return nil
}
