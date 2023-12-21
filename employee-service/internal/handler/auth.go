package handler

import (
	"Insurance/internal/model"
	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
	"resenje.org/logging"
	"time"
)

func (h *Handler) GenerateHashSingInHandler(ctx *fasthttp.RequestCtx) (res *model.GenerateHashHandlerRes, message string, code int) {
	defer func() {
		start := time.Now()
		resFinal := model.Response{Data: res, Code: code, Message: message}
		if code != fasthttp.StatusOK {
			logging.Info(message)
		}
		jsonRes, _ := json.Marshal(resFinal)
		ctx.Response.SetStatusCode(resFinal.Code)
		ctx.Response.SetBody(jsonRes)
		logging.Info(time.Since(start))
	}()
	start := time.Now()
	if !ctx.IsPost() {
		return nil, "handler: check method: wrong method", fasthttp.StatusMethodNotAllowed
	}

	var req *model.GenerateHashHandlerReq

	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		return nil, "handler: unmarshal request: " + err.Error() + ": wrong format of input data", fasthttp.StatusUnprocessableEntity
	}

	if err := req.Validate(); err != nil {
		return nil, err.Message, err.Code
	}

	req.IP = ctx.RemoteIP().String()

	res, err := h.services.GenerateHashSingInService(req)
	if err != nil {
		return nil, err.Message, err.Code
	}
	logging.Info(time.Since(start))
	return res, "OK", fasthttp.StatusOK
}
