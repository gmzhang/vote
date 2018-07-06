package web

import (

	pkgVoteHandler "vote/models/vote/deliver/http"
	pkgVoteRepos "vote/models/vote/repository"
	pkgVoteUsecase "vote/models/vote/usecase"

	"vote/utils"
	"sync"

	"github.com/labstack/echo"
)

func NewControl(e *echo.Echo) *Control {
	ctl := Control{}
	ctl.e = e
	ctl.singleRunnerLock = sync.Map{}

	voteRepos := pkgVoteRepos.NewVoteRepositoryFromDB(utils.DB)
	portalUsecase := pkgVoteUsecase.NewVoteUsecase(voteRepos)
	ctl.voteUsecase = portalUsecase

	return &ctl
}

func (c *Control) Dispatch(g *echo.Group) {
	pkgVoteHandler.NewVoteHttpHandler(g, c.voteUsecase)

}

// ActivityCtl control request and route
type Control struct {
	e                *echo.Echo
	singleRunnerLock sync.Map

	voteUsecase pkgVoteUsecase.VoteUsecase
}
