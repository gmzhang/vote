package repository

import (
	"encoding/json"
	"fmt"
	"vote/utils"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"vote/models/vote"
	"mudutv/web/models/portal"
)

func NewVoteRepositoryFromDB(conn *sqlx.DB) VoteRepository {
	return &mysqlVoteRepos{conn}
}

type mysqlVoteRepos struct {
	Conn *sqlx.DB
}

func (m *mysqlVoteRepos) Get() []vote.Vote {
	sql := "select id, name from vote"

	Votes := []vote.Vote{}
	err := m.Conn.Select(&Votes, sql)

	if err != nil {
		utils.Logger.WithField("err", err).Errorln("Get Vote Error")
		return nil
	}
	return Votes
}

