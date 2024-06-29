package ginutils

import "github.com/gin-gonic/gin"

func GetCookieOrEmpty(
	ctx *gin.Context,
	name string,
) string {
	res, _ := ctx.Cookie(name)

	return res
}
