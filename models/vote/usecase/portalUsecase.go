package usecase

import (
	repository "vote/models/vote/repository"
	"vote/models/vote"
)

type VoteUsecase interface {
	Get() []vote.Vote
	PostVote(postVotes []vote.PostVote, visitorId string) error
}

type voteUsecase struct {
	voteRepos repository.VoteRepository
}

func NewVoteUsecase(repos repository.VoteRepository) VoteUsecase {
	return &voteUsecase{repos}
}

func (v *voteUsecase) Get() []vote.Vote {
	return v.voteRepos.Get()
}

func (v *voteUsecase) PostVote(postVotes []vote.PostVote, visitorId string) error {
	return v.voteRepos.PostVote(postVotes, visitorId)
}