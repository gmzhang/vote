package vote

type VoteQuestion struct {
	ID     int    `db:"id" json:"-"`
	VoteId int    `db:"vote_id" json:"-"`
	Name   string `db:"name"`
	Items []VoteQuestionItem
}

