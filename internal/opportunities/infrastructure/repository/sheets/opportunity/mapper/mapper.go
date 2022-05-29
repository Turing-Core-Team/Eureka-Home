package mapper

import (
	"EurekaHome/internal/opportunities/core/entity"
	"EurekaHome/internal/opportunities/infrastructure/repository/sheets/opportunity/model"
	"encoding/json"
	"errors"
)

type Mapper struct{}

func (m Mapper) ModelToDomain(item []string) ([]entity.Opportunity, error){
	response := make([]entity.Opportunity, len(item))
	for i := range item{
		out := model.OpportunityModel{}
		err := json.Unmarshal([]byte(item[i]), out)
		if err != nil{
			return nil, errors.New("error unmarshal json opportunity from sheets")
		}
		entityOpp := entity.Opportunity{
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
		response = append(response, entityOpp)
	}

	return response, nil
}
