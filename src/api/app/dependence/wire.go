package dependence

import (
	useCaseOpportunities "EurekaHome/internal/opportunities/core/usecase/getopportunity"
	RepositoryRead "EurekaHome/internal/opportunities/infrastructure/repository/sheets/opportunity"
	RepositoryReadMapper "EurekaHome/internal/opportunities/infrastructure/repository/sheets/opportunity/mapper"
	"EurekaHome/internal/platform/constant"
	platformParams "EurekaHome/internal/platform/params"
	"EurekaHome/internal/platform/sheets"
	"EurekaHome/src/api/app/config"
	handlerGetOpportunities "EurekaHome/src/api/handler/getopportunities"
	mapperGetOpportunities "EurekaHome/src/api/handler/getopportunities/mapper"
)

type HandlerContainer struct {
	GetOpportunitiesHandler handlerGetOpportunities.Handler
}

func NewWire() HandlerContainer {
	var sheetsClients sheets.Client
	switch config.GetConfig().Scope() {
	case constant.DevScope:
		sheetsClients = sheets.NewSheetsClientMock()
	default:
		sheetsClients = config.GetConfig().OpportunitiesSheetsClient().NewSheetsClient()
	}

	mapperClient := RepositoryReadMapper.Mapper{}
	repositoryRead := RepositoryRead.NewRepositoryClient(sheetsClients, mapperClient)
	useCaseGetOpportunity := useCaseOpportunities.NewUseCase(repositoryRead)
	return HandlerContainer{
		GetOpportunitiesHandler: newWireGetOpportunitiesHandler(useCaseGetOpportunity),
	}
}

func newWireGetOpportunitiesHandler(useCase handlerGetOpportunities.UseCase) handlerGetOpportunities.Handler {

	return *handlerGetOpportunities.NewHandler(
		useCase,
		mapperGetOpportunities.OpportunityMapper{},
		platformParams.NewParamValidation(getParamsValidationDefault()),
	)
}

func getParamsValidationDefault() map[string]platformParams.ValidationParams {
	paramsMap := make(map[string]platformParams.ValidationParams)
	paramsMap[platformParams.FirstFilterValidator{}.KeyParam()] = platformParams.FirstFilterValidator{IsRequired: true}
	paramsMap[platformParams.SecondFilterValidator{}.KeyParam()] = platformParams.SecondFilterValidator{IsRequired: true}
	paramsMap[platformParams.ThirdFilterValidator{}.KeyParam()] = platformParams.ThirdFilterValidator{IsRequired: true}
	return paramsMap
}
