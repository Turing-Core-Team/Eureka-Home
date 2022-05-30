package mapper

import (
	"EurekaHome/internal/opportunities/core/query"
	"EurekaHome/internal/platform/sheets/structure/columns"
	"EurekaHome/src/api/handler/getopportunities/contract"
	"encoding/json"
	"errors"
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
	secondFilterSplit := strings.Split(request.SecondFilter, "-")

	getOpportunities := make([]query.GetOpportunity, 0)

	for i := range thirdFilterSplit {
		for j := range secondFilterSplit {
			column, err := columns.GetRange(isFirstPartition, secondFilterSplit[j])
			if err != nil && column != "" {
				fmt.Println("THERE IS NO VALID EQUIVALENCE FOR ", secondFilterSplit[i])
				// TODO report this error as ignored for the search
			} else {

				getOpportunities = append(getOpportunities, query.GetOpportunity{
					Sheet:  thirdFilterSplit[i],
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

func (om OpportunityMapper) EntityToResponse(entityOpp []string, fourthFilter string) ([]contract.OpportunitiesResponse, error){
	/*
		if fourthFilter != ""{
			fourthFilterSplit := strings.Split(fourthFilter, "-")
			fourthFilterSplit
		}

	*/
	response := make([]contract.OpportunitiesResponse, 0)

	for i := range entityOpp {

		out := &contract.OpportunitiesResponse{}
		err := json.Unmarshal([]byte(entityOpp[i]), out)
		if err != nil{
			return nil, errors.New("error unmarshal json opportunity from sheets")
		}
		entityResponse := contract.OpportunitiesResponse{
			Tags: out.Tags,
			Link: out.Link,
			Title: out.Title,
			Requirements: out.Requirements,
			Awards: out.Awards,
			Description: out.Description,
			PublicationDate: out.PublicationDate,
			UpdateDate: out.UpdateDate,
			DueDate: out.DueDate,
		}
		response = append(response, entityResponse)
	}
	return response, nil
}
