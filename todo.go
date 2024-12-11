package todogoapp

import "time"

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

/*
Add method
	-> add the task
	-> done is set to false
	-> created time
	-> completed time yet to give
*/