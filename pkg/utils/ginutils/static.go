package ginutils

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func StaticRedirect(url string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, url)
	}
}

func StaticTempl(component templ.Component) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		RenderTempl(ctx, http.StatusOK, component)
	}
}
