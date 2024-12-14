package xfiber

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

type ResponseWithError interface {
	SetErrorMessage(message string)
	GetErrorMessage() string
}

func FiberResponse(ctx fiber.Ctx, res ResponseWithError) {
	body, err := json.Marshal(res)
	if err != nil {
		res.SetErrorMessage(err.Error())
	}

	if res.GetErrorMessage() != "" {
		ctx.Status(http.StatusInternalServerError)
	}

	ctx.Set("Content-Type", "application/json")

	_ = ctx.Send(body)
}
