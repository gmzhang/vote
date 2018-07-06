package vote

type Vote struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	Questions []VoteQuestion
}

type PostVote struct {
	VoteID     int `json:"vote_id"`
	QuestionID int `json:"question_id"`
	ItemID     int `json:"item_id"`
}
