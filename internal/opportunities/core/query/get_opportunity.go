package query

import "fmt"

const (
	start string = "3"
	end string = "500"
)

type GetOpportunity struct {
	Sheet  string
	Column string
}

func (g GetOpportunity) Value() string{
	return fmt.Sprintf(
		`%s!%s:%s`,
		g.Sheet,
		g.Column+start,
		g.Column+end,
	)
}