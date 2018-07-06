package usecase

import (
	repository "vote/models/vote/repository"
	"vote/models/vote"
)

type VoteUsecase interface {
	Get() []vote.Vote
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
