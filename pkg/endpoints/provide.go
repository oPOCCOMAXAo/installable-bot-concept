package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/opoccomaxao/installable-bot-concept/pkg/server"
	"github.com/samber/do"
)

type route func(gin.IRouter, *do.Injector) error

func Provide(
	i *do.Injector,
) {
	do.ProvideNamed[bool](i, "endpoints", func(i *do.Injector) (bool, error) {
		router, err := server.InvokeRouter(i)
		if err != nil {
			return false, err
		}

		for _, r := range []route{
			Init,
		} {
			err := r(router, i)
			if err != nil {
				return false, err
			}
		}

		return true, nil
	})

}

func Invoke(i *do.Injector) error {
	_, err := do.InvokeNamed[bool](i, "endpoints")

	return err
}
