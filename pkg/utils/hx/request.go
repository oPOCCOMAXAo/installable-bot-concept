package hx

import "github.com/gin-gonic/gin"

func Is(
	ctx *gin.Context,
) bool {
	h := ctx.GetHeader("Hx-Request")

	return len(h) > 0
}
