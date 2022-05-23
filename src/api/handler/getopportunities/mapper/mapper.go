package mapper

import (
	"EurekaHome/internal/opportunities/core/query"
	"EurekaHome/internal/platform/sheets/structure/columns"
	"EurekaHome/src/api/handler/getopportunities/contract"
)

type OpportunityMapper struct{}

func (om OpportunityMapper) RequestToQuery(request contract.URLParams) query.GetOpportunity{
	queryOpportunities := query.GetOpportunity{
		Sheet: request.F1,
		Column: query.Column{
			Start: columns.
				End:
		}
	}

	return queryOpportunities
}

