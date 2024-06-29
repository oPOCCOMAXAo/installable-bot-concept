package hx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Redirect(
	ctx *gin.Context,
	path string,
) {
	if Is(ctx) {
		ctx.Header("Hx-Redirect", path)
	} else {
		ctx.Redirect(http.StatusFound, path)
	}

	ctx.Abort()
}
