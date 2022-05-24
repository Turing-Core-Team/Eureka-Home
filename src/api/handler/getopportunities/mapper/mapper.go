package mapper

import (
	"EurekaHome/internal/opportunities/core/entity"
	"EurekaHome/internal/opportunities/core/query"
	"EurekaHome/internal/platform/sheets/structure/columns"
	"EurekaHome/src/api/handler/getopportunities/contract"
	errorMapper "EurekaHome/src/api/handler/getopportunities/error"
	"strings"
)

const (
	person   string = "personas"
	projects string = "proyectos"
)

type OpportunityMapper struct{}

func (om OpportunityMapper) RequestToQuery(request contract.URLParams) ([]query.GetOpportunity, error) {
	var isFirstPartition bool
	switch request.FirstFilter {
	case person:
		isFirstPartition = true
	case projects:
		isFirstPartition = false
	}

	thirdFilterSplit := strings.Split(request.ThirdFilter, "-")

	getOpportunities := make([]query.GetOpportunity, len(thirdFilterSplit))
	isErrorGetRange := false
	errorsOpportunities //TODO  mirar partial content

	for i := range thirdFilterSplit {
		column, err := columns.GetRange(isFirstPartition, thirdFilterSplit[i])
		if err != nil {
			isErrorGetRange = true
			errorsOpportunities = append(errorsOpportunities, errorMapper.NewUnknownCategory("", thirdFilterSplit[i]))
		}
		getOpportunities = append(getOpportunities, query.GetOpportunity{
			Sheet:  request.FirstFilter,
			Column: column,
		})
	}
	return getOpportunities
}

func (om OpportunityMapper) EntityToResponse(entity []entity.Opportunity) []contract.OpportunitiesResponse {

}
