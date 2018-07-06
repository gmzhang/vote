package http

import (
	"github.com/labstack/echo"
	"vote/models/vote/usecase"
	"vote/models/vote"
	"time"
	"vote/utils"
	"fmt"
)

type VoteHttpHandler struct {
	VoteUsecase usecase.VoteUsecase
}

func (v VoteHttpHandler) Get(c echo.Context) error {
	votes := v.VoteUsecase.Get()
	return c.JSON(200, votes)
}

func (v VoteHttpHandler) Vote(c echo.Context) error {
	result := struct {
		ErrCode int `json:"errCode"`
	}{10000}

	postInfo := []vote.PostVote{}

	if err := c.Bind(&postInfo); err != nil {
		result.ErrCode = 10001
		return c.JSON(200, result)
	}

	//votedCookie, err := c.Cookie("voted")
	//if votedCookie.Value == "1" {
	//	result.ErrCode = 10003
	//	return c.JSON(200, result)
	//}

	visitorId := ""
	visitorIdCookie, err := c.Cookie("visitorId")
	if err != nil {
		//没有，创建
		visitorId = fmt.Sprintf("%d%s", time.Now().Unix(), utils.RandomString(10))
	} else {
		visitorId = visitorIdCookie.Value
	}
	err = v.VoteUsecase.PostVote(postInfo, visitorId)
	if err != nil {
		result.ErrCode = 10002
		return c.JSON(200, result)
	}
	//c.SetCookie(utils.CreateCookie("voted", "1", time.Now().Add(time.Hour*24*365), "/"))
	c.SetCookie(utils.CreateCookie("visitorId", visitorId, time.Now().Add(time.Hour*24*365), "/"))
	return c.JSON(200, result)
}

func NewVoteHttpHandler(g *echo.Group, usecase usecase.VoteUsecase) {
	handler := VoteHttpHandler{
		VoteUsecase: usecase,
	}
	g.GET("", handler.Get)
	g.POST("", handler.Vote)
}
