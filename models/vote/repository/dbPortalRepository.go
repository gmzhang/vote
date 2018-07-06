package repository

import (
	"vote/utils"

	"github.com/jmoiron/sqlx"
	"vote/models/vote"
	"fmt"
)

func NewVoteRepositoryFromDB(conn *sqlx.DB) VoteRepository {
	return &mysqlVoteRepos{conn}
}

type mysqlVoteRepos struct {
	Conn *sqlx.DB
}

func (m *mysqlVoteRepos) Get() []vote.Vote {
	sql := "select id, name from vote"

	votes := []vote.Vote{}
	err := m.Conn.Select(&votes, sql)
	if err != nil {
		utils.Logger.WithField("err", err).Errorln("Get Vote Error")
		return nil
	}

	for voteKey, voteVal := range votes {

		voteRecordCountSql := "SELECT count(*) FROM vote_record WHERE vote_id = ?"
		voteRecordCount := utils.RowCount(voteRecordCountSql, voteVal.ID) //总投票数

		questionsSql := "SELECT id, vote_id, name FROM vote_question  WHERE vote_id = ?"
		questions := []vote.VoteQuestion{}
		err = m.Conn.Select(&questions, questionsSql, voteVal.ID)
		if err != nil {
			utils.Logger.WithField("err", err).Errorln("Get Vote Question Error")
			return nil
		}

		for questionKey, questionVal := range questions {
			itemSql := "SELECT * FROM vote_question_item where vote_id = ? and question_id = ?"
			items := []vote.VoteQuestionItem{}
			err = m.Conn.Select(&items, itemSql, voteVal.ID, questionVal.ID)
			if err != nil {
				utils.Logger.WithField("err", err).Errorln("Get Vote Question Item Error")
				return nil
			}

			for itemKey, itemVal := range items {
				itemCountSql := "SELECT count(*) FROM vote_record WHERE vote_id = ? and question_id = ? and item_id = ?"
				itemCount := utils.RowCount(itemCountSql, voteVal.ID, questionVal.ID, itemVal.ID)

				if voteRecordCount == 0 {
					items[itemKey].Percent = "0%";
					items[itemKey].Sum = 0;
				} else {
					itemPercent := float64(itemCount) / float64(voteRecordCount) * 100
					items[itemKey].Sum = itemCount
					items[itemKey].Percent = fmt.Sprintf("%.2f", itemPercent) + "%"
				}
			}
			questions[questionKey].Items = items
		}

		votes[voteKey].Questions = questions
	}

	return votes
}

func (m *mysqlVoteRepos) PostVote(postVotes []vote.PostVote, visitorId string) error {

	for _, postVal := range postVotes {

		sqlRecord := "INSERT into vote_record (vote_id,question_id,item_id,visitor_id)" +
			" VALUES (:vote_id,:question_id,:item_id,:visitor_id)"
		data := map[string]interface{}{
			"vote_id":     postVal.VoteID,
			"question_id": postVal.QuestionID,
			"item_id":     postVal.ItemID,
			"visitor_id":  visitorId,
		}
		_, err := utils.DB.NamedExec(sqlRecord, data)
		if err != nil {
			utils.Logger.WithError(err).Println("insert vote_record error")
			return err
		}
	}
	return nil
}
