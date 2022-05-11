package dependence

import (
	"EurekaHome/internal/platform/constant"
	"EurekaHome/src/api/app/config"
	handlerGetOpportunities "EurekaHome/src/api/handler/getopportunities"
)

type HandlerContainer struct {
	GetOpportunitiesHandler handlerGetOpportunities.Handler
}

func NewWire() HandlerContainer {
	conf := config.GetConfig()
	switch conf.Scope() {
	case constant.DevScope:
		sheetsClient = sheetsLocal.NewKVSLocalClient()
	default:
		sheetsClient = conf.SheetsConfig().NewSheetsClient()
	}


	repositoryRead := getOpportunitiesRepositoryRead(sheetsClient)

	cardConfigMapper := styleCardConfigMapper.Mapper{}

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

func newInformationHandler(fullUseCase, simpleUseCase handlerGetInformation.UseCase) handlerGetInformation.Handler {
	useCaseMap := make(map[bool]handlerGetInformation.UseCase)
	useCaseMap[true] = fullUseCase
	useCaseMap[false] = simpleUseCase

	return *handlerGetInformation.NewHandler(
		useCaseMap,
		mapperGetInformation.InformationMapper{},
		platformParams.NewParamValidation(getParamsValidationDefault()),
	)
}

func getFullUseCaseGetInformation(accountService fullUseCaseGetInformation.AccountService,
	orchestratorService fullUseCaseGetInformation.OrchestratorService,
	repositoryWrite fullUseCaseGetInformation.RepositoryWrite) handlerGetInformation.UseCase {
	return fullUseCaseGetInformation.NewGetCardInformation(
		accountService, orchestratorService, repositoryWrite,
	)
}

func getSimpleUseCaseGetInformation(accountService simpleUseCaseGetInformation.AccountService,
	orchestratorService simpleUseCaseGetInformation.OrchestratorService,
	repositoryWrite simpleUseCaseGetInformation.RepositoryWrite) handlerGetInformation.UseCase {
	return simpleUseCaseGetInformation.NewGetCardInformation(
		accountService, orchestratorService, repositoryWrite,
	)
}

func getLimitRestClient() limitsRestClient.ProcessorLimit {
	scope := appConfig.GetConfig().Scope()
	uri := appConfig.GetConfig().ProcessorLimitClient().BaseURL + appConfig.GetConfig().ProcessorLimitClient().LimitURI
	var httpClient rusty.Requester
	switch scope {
	case constant.DevScope:
		httpClient = dummyLimits.HTTPClient{}
	default:
		httpClient = appConfig.GetConfig().ProcessorLimitClient().GetHTTPClient()
	}

	limitClient := platformLimitsClient.NewLimitRestClient(
		httpClient,
		platformLimitsMapper.Mapper{},
		uri,
	)

	return limitsRestClient.NewProcessorLimit(
		limitClient,
		externalLimitsMapper.Mapper{},
	)
}

func getCardRepositoryRead(kvsCacheClient KVS.Client) account.CardRepositoryRead {
	return read.NewRepositoryCacheRead(kvsCacheClient, cardRepositoryMapper.Mapper{})
}

func getCardRepositoryWrite(kvsCacheClient KVS.Client) account.CardRepositoryWrite {
	return write.NewRepositoryCacheWrite(kvsCacheClient, cardRepositoryMapper.Mapper{})
}

func getAccountRestClient() accountRestClient.RestClient {
	scope := appConfig.GetConfig().Scope()
	uri := appConfig.GetConfig().ProcessorAccountClient().BaseURL + appConfig.GetConfig().ProcessorAccountClient().AccountURI
	var httpClient rusty.Requester
	switch scope {
	case constant.DevScope:
		httpClient = dummyAccount.HTTPClient{}
	default:
		httpClient = appConfig.GetConfig().ProcessorAccountClient().GetHTTPClient()
	}

	accountClient := platformAccountClient.NewAccountRestClient(
		httpClient,
		&platformAccountMapper.Mapper{},
		uri,
	)

	return *accountRestClient.NewRestClient(
		&externalAccountMapper.Mapper{},
		accountClient,
	)
}

func getStatementRestClient() statementsRestClient.ProcessorStatement {
	uri := appConfig.GetConfig().ProcessorStatementClient().BaseURL +
		appConfig.GetConfig().ProcessorStatementClient().StatementURI
	httpClient := dummyStatement.HTTPClient{}

	statementRestClient := platformStatementClient.NewStatementRestClient(
		&platformStatementMapper.Mapper{},
		httpClient,
		uri,
	)

	return *statementsRestClient.NewProcessorStatement(
		&statementRestClient,
		&externalStatementMapper.Mapper{},
	)
}

func getParamsValidationDefault() map[string]platformParams.ValidationParams {
	paramsMap := make(map[string]platformParams.ValidationParams)
	paramsMap[platformParams.SiteValidator{}.KeyParam()] = platformParams.SiteValidator{IsRequired: true}
	paramsMap[platformParams.UserValidator{}.KeyParam()] = platformParams.UserValidator{IsRequired: true}
	return paramsMap
}

func getStatementStatus() map[query.GetCardConfig]model.Style {
	return map[query.GetCardConfig]model.Style{
		query.NewGetCardConfig(constant.Active, constant.Open):            model.NewOpen(),
		query.NewGetCardConfig(constant.Active, constant.Close):           model.NewClosed(),
		query.NewGetCardConfig(constant.Active, constant.Overdue):         model.NewOverdue(),
		query.NewGetCardConfig(constant.Blocked, constant.Overdue):        model.NewBlocked(),
		query.NewGetCardConfig(constant.Active, constant.Default):         model.NewDefault(),
		query.NewGetCardConfig(constant.Blocked, constant.Default):        model.NewBlocked(),
		query.NewGetCardConfig(constant.DefaultAccount, constant.Default): model.NewDefault(),
	}
}