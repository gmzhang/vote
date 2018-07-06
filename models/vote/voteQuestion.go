package vote

type VoteQuestion struct {
	ID     int    `db:"id"`
	VoteId int    `db:"vote_id" json:"-"`
	Name   string `db:"name"`
	Items []VoteQuestionItem
}

