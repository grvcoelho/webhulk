package handlers

import (
	"time"

	iris "gopkg.in/kataras/iris.v8"
)

func NewHealth() *HealthHandler {
	return &HealthHandler{}
}

type HealthHandler struct{}

type HealthGetResponse struct {
	Time   string `json:"time"`
	Status string `json:"status"`
}

func (h *HealthHandler) Get(ctx iris.Context) {
	now := time.Now().Format(time.RFC3339)
	r := HealthGetResponse{
		Time:   now,
		Status: "ok",
	}

	time.Sleep(2000 * time.Millisecond)

	ctx.JSON(r)
}
