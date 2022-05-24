package getopportunity

import (
	"EurekaHome/internal/opportunities/core/entity"
	ErrorUseCase "EurekaHome/internal/opportunities/core/error"
	"EurekaHome/internal/opportunities/core/query"
	"EurekaHome/internal/platform/constant"
	"EurekaHome/internal/platform/log"
	"context"
	"fmt"
)

const (
	action              string          = "execute_use_case"
	errorReadRepository log.LogsMessage = "error in the use case, when read repository"
	entityType          string          = "read_repository"
	layer               string          = "use_case_get_opportunities"
)

type RepositoryRead interface {
	GetByQuery(ctx context.Context, queryValue string) (entity.Opportunity, error)
}

type UseCase struct {
	repositoryRead RepositoryRead
}

func NewUseCase(repositoryRead RepositoryRead) UseCase {
	return UseCase{repositoryRead: repositoryRead}
}

func (uc UseCase) Execute(ctx context.Context, getOpportunityQuery query.GetOpportunity) (entity.Opportunity, error) {

	opportunity, err := uc.repositoryRead.GetByQuery(ctx, getOpportunityQuery.Value())

	if err != nil {
		message := errorReadRepository.GetMessageWithTagParams(
			log.NewTagParams(layer, action,
				log.Params{
					constant.Key: fmt.Sprintf(
						`%s_%s_%s`,
						getOpportunityQuery.Sheet,
						getOpportunityQuery.Column,
						getOpportunityQuery.Value(),
					),
					constant.EntityType: entityType,
				}))
		return entity.Opportunity{}, ErrorUseCase.FailedQueryValue{
			Message: message,
			Err: err,
		}
	}

	return opportunity, nil
}
