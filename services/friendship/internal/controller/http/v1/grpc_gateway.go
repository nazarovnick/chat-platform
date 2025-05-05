package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"net/http"
)

func RegisterGrpcGatewayRoutes(router fiber.Router, gwMux http.Handler) {
	router.All("/*", func(c *fiber.Ctx) error {
		reqHandler := fasthttpadaptor.NewFastHTTPHandler(gwMux)
		reqHandler(c.Context())
		return nil
	})
}
