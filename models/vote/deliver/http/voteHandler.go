package http

import (
	"github.com/labstack/echo"
	"vote/models/vote/usecase"
)

type VoteHttpHandler struct {
	VoteUsecase usecase.VoteUsecase
}

func (v VoteHttpHandler) Get(c echo.Context) error {
	votes := v.VoteUsecase.Get()
	return c.JSON(200, votes)
}

func NewVoteHttpHandler(g *echo.Group, usecase usecase.VoteUsecase) {
	handler := VoteHttpHandler{
		VoteUsecase: usecase,
	}
	g.GET("", handler.Get)
}
