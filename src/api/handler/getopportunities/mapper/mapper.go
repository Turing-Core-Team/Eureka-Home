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

	getOpportunities := make([]query.GetOpportunity, len(thirdFilterSplit)-1)

	for i := range thirdFilterSplit {
		secondFilterSplit := strings.Split(request.SecondFilter, "-")
		for j := range secondFilterSplit {
			column, err := columns.GetRange(isFirstPartition, secondFilterSplit[j])
			if err != nil && column != ""{
				fmt.Println("THERE IS NO VALID EQUIVALENCE FOR ", secondFilterSplit[i])
				// TODO report this error as ignored for the search
			} else {

				getOpportunities = append(getOpportunities, query.GetOpportunity{
					Sheet:  request.FirstFilter,
					Column: column,
				})
			}
		}
	}

	if len(getOpportunities) == 0 {
		return getOpportunities, fmt.Errorf("empty response")
	}
	return getOpportunities, nil
}

func (om OpportunityMapper) EntityToResponse(entityOpp []entity.Opportunity, fourthFilter string) []contract.OpportunitiesResponse {
	/*
		if fourthFilter != ""{
			fourthFilterSplit := strings.Split(fourthFilter, "-")
			fourthFilterSplit
		}

	*/
	response := make([]contract.OpportunitiesResponse, len(entityOpp))

	for i := range entityOpp {
		response = append(response, contract.OpportunitiesResponse{
			Tags:            entityOpp[i].Tags,
			Link:            entityOpp[i].Link,
			Title:           entityOpp[i].Title,
			Requirements:    entityOpp[i].Requirements,
			Awards:          entityOpp[i].Awards,
			Description:     entityOpp[i].Description,
			PublicationDate: entityOpp[i].PublicationDate,
			UpdateDate:      entityOpp[i].UpdateDate,
			DueDate:         entityOpp[i].DueDate,
		})
	}
	return response
}
