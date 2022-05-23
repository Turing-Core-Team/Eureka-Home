package mapper

import (
	"EurekaHome/internal/opportunities/core/query"
	"EurekaHome/internal/platform/sheets/structure/columns"
	"EurekaHome/src/api/handler/getopportunities/contract"
)

const (
	personas string = "personas"
	proyectos string = "proyectos"
	end string = "500"
)

type OpportunityMapper struct{}

func (om OpportunityMapper) RequestToQuery(request contract.URLParams) query.GetOpportunity {
	isFirstPartition := true
	switch request.F1{
	case personas:
		isFirstPartition = true
	case proyectos:
		isFirstPartition = false
	}

	queryOpportunities := query.GetOpportunity{
		Sheet: request.F1,
		Column: columns.GetRange(isFirstPartition),
	}

	return queryOpportunities
}
