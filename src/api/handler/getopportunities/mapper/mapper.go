package mapper

import (
	"EurekaHome/internal/opportunities/core/entity"
	"EurekaHome/internal/opportunities/core/query"
	"EurekaHome/internal/platform/sheets/structure/columns"
	"EurekaHome/src/api/handler/getopportunities/contract"
	"fmt"
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

	for i := range thirdFilterSplit {
		column, err := columns.GetRange(isFirstPartition, thirdFilterSplit[i])
		if err != nil {
			fmt.Println("THERE IS NO VALID EQUIVALENCE FOR ", thirdFilterSplit[i])
			// TODO report this error as ignored for the search
		}else{
			getOpportunities = append(getOpportunities, query.GetOpportunity{
				Sheet:  request.FirstFilter,
				Column: column,
			})
		}
	}

	if len(getOpportunities) == 0 {
		return getOpportunities, fmt.Errorf("empty response")
	}
	return getOpportunities, nil
}

func (om OpportunityMapper) EntityToResponse(entity []entity.Opportunity) []contract.OpportunitiesResponse {

}
