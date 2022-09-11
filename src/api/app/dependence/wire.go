package dependence

import (
	useCaseOpportunities "EurekaHome/internal/opportunities/core/usecase/getopportunity"
	RepositoryRead "EurekaHome/internal/opportunities/infrastructure/repository/sheets/opportunity"
	platformParams "EurekaHome/internal/platform/params"
	"EurekaHome/internal/platform/sheets"
	handlerGetOpportunities "EurekaHome/src/api/handler/getopportunities"
	mapperGetOpportunities "EurekaHome/src/api/handler/getopportunities/mapper"
	handlerPing "EurekaHome/src/api/handler/ping"
	"fmt"
)

type HandlerContainer struct {
	GetOpportunitiesHandler handlerGetOpportunities.Handler
	PingHandler handlerPing.Handler
}

func NewWire() HandlerContainer {
	sheetsClients := sheets.Client{}
	repositoryRead := RepositoryRead.NewRepositoryClient(sheetsClients)
	useCaseGetOpportunity := useCaseOpportunities.NewUseCase(repositoryRead)
	return HandlerContainer{
		PingHandler: newWirePingHandler(),
		GetOpportunitiesHandler: newWireGetOpportunitiesHandler(useCaseGetOpportunity),
	}
}

func newWirePingHandler() handlerPing.Handler{
	return *handlerPing.NewHandler()
}

func newWireGetOpportunitiesHandler(useCase handlerGetOpportunities.UseCase) handlerGetOpportunities.Handler {
	fmt.Println("WIRE")
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
