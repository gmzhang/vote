package vote

type VoteQuestionItem struct {
	ID         int    `db:"id"`
	VoteId     int    `db:"vote_id" json:"-"`
	QuestionId int    `db:"question_id" json:"-"`
	Name       string `db:"name"`
	Percent    string
	Sum        int
}
