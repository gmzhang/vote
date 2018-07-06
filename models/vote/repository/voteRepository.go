package repository

import (
	"vote/models/vote"
)

type VoteRepository interface {
	Get() []vote.Vote
	PostVote(postVotes []vote.PostVote, visitorId string) error
}
