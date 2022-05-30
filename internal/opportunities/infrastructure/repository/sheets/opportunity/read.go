package opportunity

import (
	ErrorUseCase "EurekaHome/internal/opportunities/core/error"
	"EurekaHome/internal/platform/constant"
	"EurekaHome/internal/platform/log"
	sheet "EurekaHome/internal/platform/sheets"
	"context"
	"fmt"
)

const (
	action                   string          = "execute_use_case"
	errorReadRepository      log.LogsMessage = "error in the use case, when read repository"
	errorMapperModelToDomain log.LogsMessage = "error in the mapper, when read repository"
	entityType               string          = "read_repository"
	pathCredentials          string          = "internal/platform/sheets/environment/credentials.json"
	spreadsheetID            string          = "1y8rAuDeYfcT7jx6O1KbC_DwO68SeWeoXkeUdYVMCV_0"
	layer                    string          = "use_case_get_opportunities"
)


type RepositoryClient struct {
	client sheet.Client
}

func NewRepositoryClient(client sheet.Client) *RepositoryClient {
	return &RepositoryClient{client: client}
}

func (rc RepositoryClient) GetByQuery(ctx context.Context, queryValue string) ([]string, error) {

	item, errReadClient := rc.client.Read(ctx, pathCredentials, spreadsheetID, queryValue)

	if errReadClient != nil {
		message := errorReadRepository.GetMessageWithTagParams(
			log.NewTagParams(layer, action,
				log.Params{
					constant.Key: fmt.Sprintf(
						`%s_%s_%s`,
						pathCredentials,
						spreadsheetID,
						queryValue,
					),
					constant.EntityType: entityType,
				}))
		return nil, ErrorUseCase.FailedQueryValue{
			Message: message,
			Err:     errReadClient,
		}
	}

	return item, nil
}
