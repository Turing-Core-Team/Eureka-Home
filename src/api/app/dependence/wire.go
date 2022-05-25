package dependence

import (
	"EurekaHome/internal/platform/constant"
	"EurekaHome/src/api/app/config"
	handlerGetOpportunities "EurekaHome/src/api/handler/getopportunities"
	sheetsLocal "EurekaHome/internal/platform/sheets"
	platformParams "EurekaHome/internal/platform/params"
)

type HandlerContainer struct {
	GetOpportunitiesHandler handlerGetOpportunities.Handler
}

func NewWire() HandlerContainer {
	conf := config.GetConfig()
	switch conf.Scope() {
	case constant.DevScope:
		sheetsClient = sheetsLocal.NewSheetsLocalClient()
	default:
		sheetsClient = conf.SheetsConfig().NewSheetsClient()
	}

	repositoryRead := getOpportunitiesRepositoryRead(sheetsClient)


	styleMapper := styleDefaultMapper.DefaultMapper{}
	styleCardConfig := styleCardConfiguration.NewProcessMiniCard(styleByStatementStatus, cardConfigMapper)

	accountService := account.NewService(
		cardRepositoryRead,
		cardRepositoryWrite,
		&repositoryAccount,
		conf.ProposalURL())

	consolidatorService := consolidator.NewService(
		repositoryLimit,
		&repositoryStatements)

	styleService := serviceStyle.NewService(
		styleCardConfig,
		styleMapper)

	orchestratorService := orchestrator.NewService(
		consolidatorService,
		styleService)

	return HandlerContainer{
		GetOpportunitiesHandler: newInformationHandler(
			getFullUseCaseGetInformation(accountService, orchestratorService, cardRepositoryWrite),
			getSimpleUseCaseGetInformation(accountService, orchestratorService, cardRepositoryWrite)),
	}
}

func newOpportunitiesHandler(fullUseCase, simpleUseCase handlerGetInformation.UseCase) handlerGetInformation.Handler {
	useCaseMap := make(map[bool]handlerGetInformation.UseCase)
	useCaseMap[true] = fullUseCase
	useCaseMap[false] = simpleUseCase

	return *handlerGetInformation.NewHandler(
		useCaseMap,
		mapperGetInformation.InformationMapper{},
		platformParams.NewParamValidation(getParamsValidationDefault()),
	)
}


func getSimpleUseCaseGetInformation(accountService simpleUseCaseGetInformation.AccountService,
	orchestratorService simpleUseCaseGetInformation.OrchestratorService,
	repositoryWrite simpleUseCaseGetInformation.RepositoryWrite) handlerGetInformation.UseCase {
	return simpleUseCaseGetInformation.NewGetCardInformation(
		accountService, orchestratorService, repositoryWrite,
	)
}


func getCardRepositoryRead(kvsCacheClient KVS.Client) account.CardRepositoryRead {
	return read.NewRepositoryCacheRead(kvsCacheClient, cardRepositoryMapper.Mapper{})
}

func getParamsValidationDefault() map[string]platformParams.ValidationParams {
	paramsMap := make(map[string]platformParams.ValidationParams)
	paramsMap[platformParams.FirstFilterValidator{}.KeyParam()] = platformParams.FirstFilterValidator{IsRequired: true}
	paramsMap[platformParams.SecondFilterValidator{}.KeyParam()] = platformParams.SecondFilterValidator{IsRequired: true}
	paramsMap[platformParams.ThirdFilterValidator{}.KeyParam()] = platformParams.ThirdFilterValidator{IsRequired: true}
	return paramsMap
}

