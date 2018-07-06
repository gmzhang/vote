package vote

type Vote struct {
	ID           int       `db:"id" json:"-"`
	Name         string    `db:"name"`
	Questions []VoteQuestion
}

