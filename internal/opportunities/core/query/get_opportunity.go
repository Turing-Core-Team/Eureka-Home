package query

import "fmt"

const (
	end string = "A500"
)

type GetOpportunity struct {
	Sheet  string
	Column string
}

func (g GetOpportunity) Value() string{
	return fmt.Sprintf(
		`%s!%s:%s`,
		g.Sheet,
		g.Column,
		end,
	)
}