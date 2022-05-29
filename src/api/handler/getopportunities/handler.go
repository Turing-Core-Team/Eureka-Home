package getopportunities

import (
	"EurekaHome/internal/opportunities/core/entity"
	"EurekaHome/internal/opportunities/core/query"
	"EurekaHome/internal/platform/constant"
	"EurekaHome/internal/platform/log"
	ErrorResponse "EurekaHome/src/api/handler"
	"EurekaHome/src/api/handler/getopportunities/contract"

	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	keyMessageError          string          = "error_in_use_case_get_data"
	actionExecuteUseCase     string          = "execute_use_case"
	actionValidateParameters string          = "validate_parameters"
	errorOpportunities       log.LogsMessage = "error in the creation handler"
	entityType               string          = "get_opportunities_by_filters"
	layer                    string          = "handler_opportunities"
)

type UseCase interface {
	Execute(ctx context.Context, GetOpportunities query.GetOpportunity) ([]entity.Opportunity, error)
}

type Mapper interface {
	RequestToQuery(request contract.URLParams) ([]query.GetOpportunity, error)
	EntityToResponse(entity []entity.Opportunity, fourthFilter string) []contract.OpportunitiesResponse
}

type ValidationParams interface {
	BindParamsAndValidation(obj interface{}, params gin.Params) error
}

type Handler struct {
	useCase          UseCase
	mapper           Mapper
	validationParams ValidationParams
}

func NewHandler(
	useCase UseCase,
	mapper Mapper,
	validationParams ValidationParams,
) *Handler {
	return &Handler{
		useCase:          useCase,
		mapper:           mapper,
		validationParams: validationParams,
	}
}

func (h Handler) Handler(ginCTX *gin.Context) {

	requestParam := &contract.URLParams{}
	if err := h.validationParams.BindParamsAndValidation(requestParam, ginCTX.Params); err != nil {
		message := errorOpportunities.GetMessageWithTagParams(
			log.NewTagParams(layer, actionExecuteUseCase,
				log.Params{
					constant.Key: fmt.Sprintf(
						`%s_%s_%s_%s`,
						requestParam.FirstFilter,
						requestParam.SecondFilter,
						requestParam.ThirdFilter,
						requestParam.FourthFilter,
					),
					constant.EntityType: entityType,
				}))
		ginCTX.JSON(http.StatusBadRequest, ErrorResponse.Response{
			Status:  http.StatusBadRequest,
			Message: message,
		})
	}

	fullQuery, mapperError := h.mapper.RequestToQuery(*requestParam)
	if mapperError != nil {
		message := errorOpportunities.GetMessageWithTagParams(
			log.NewTagParams(layer, actionExecuteUseCase,
				log.Params{
					constant.Key: fmt.Sprintf(
						`%s_%s_%s_%s`,
						requestParam.FirstFilter,
						requestParam.SecondFilter,
						requestParam.ThirdFilter,
						requestParam.FourthFilter,
					),
					constant.EntityType: entityType,
				}))
		ginCTX.JSON(http.StatusBadRequest, ErrorResponse.Response{
			Status:  http.StatusBadRequest,
			Message: message,
		})
	}

	opportunities := make([]entity.Opportunity, len(fullQuery))
	isErrorUseCase := false
	messageKey := keyMessageError

	for i := range fullQuery {
		opportunitiesUseCase, errorUseCase := h.useCase.Execute(ginCTX, fullQuery[i])

		if errorUseCase != nil {
			isErrorUseCase = true
			messageKey = fmt.Sprintf(
				`%s_%s_%s`,
				messageKey,
				fullQuery[i].Sheet,
				fullQuery[i].Column,
			)
			continue
		}

		for j := range opportunitiesUseCase{
			opportunities = append(opportunities, opportunitiesUseCase[j])
		}
	}

	if isErrorUseCase {
		message := errorOpportunities.GetMessageWithTagParams(
			log.NewTagParams(layer, actionExecuteUseCase,
				log.Params{
					constant.Key:        messageKey,
					constant.EntityType: entityType,
				}))
		ginCTX.JSON(http.StatusInternalServerError, ErrorResponse.Response{
			Status:  http.StatusInternalServerError,
			Message: message,
		})
	}

	ginCTX.JSON(http.StatusOK, h.mapper.EntityToResponse(opportunities, requestParam.FourthFilter))
}
