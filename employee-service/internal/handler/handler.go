package handler

import (
	"Insurance/internal/service"
	"fmt"
	"github.com/valyala/fasthttp"
	"resenje.org/logging"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-Type", "application/json")
	logging.Info(fmt.Sprintf("%s %s", string(ctx.Path()), string(ctx.Method())))
	switch string(ctx.Path()) {
	case "/sing-in/hash":
		h.GenerateHashSingInHandler(ctx)
	default:
		ctx.Error("Page not found", fasthttp.StatusNotFound)
	}
}
