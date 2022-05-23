package query

type GetOpportunity struct {
	Sheet  string
	Column []Column
}

type Column struct {
	Start string
	End   string
}
