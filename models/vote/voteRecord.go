package vote

type VoteRecord struct {
	ID         int    `db:"id"`
	VoteId     int    `db:"vote_id"`
	QuestionId int    `db:"question_id"`
	ItemId     int    `db:"item_id"`
	VisitorId  string `db:"visitor_id"`
	VoteTime   string `db:"vote_time"`
}
