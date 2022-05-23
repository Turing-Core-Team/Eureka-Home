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
	action                   string          = "get_opportunities"
	actionExecuteUseCase     string          = "execute_use_case"
	actionValidateParameters string          = "validate_parameters"
	errorOpportunities       log.LogsMessage = "error in the creation handler"
	entityType               string          = "get_opportunities_by_filters"
	layer                    string          = "handler_information"
)

type UseCase interface {
	Execute(ctx context.Context, GetOpportunities query.GetOpportunity) (entity.Opportunity, error)
}

type Mapper interface {
	RequestToQuery(request contract.URLParams) query.GetOpportunity
	EntityToResponse(entity entity.Opportunity) contract.OpportunitiesResponse
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

func (h Handler) handler(ginCTX *gin.Context) error {

	requestParam := &contract.URLParams{}
	if err := h.validationParams.BindParamsAndValidation(requestParam, ginCTX.Params); err != nil {
		message := errorOpportunities.GetMessageWithTagParams(
			log.NewTagParams(layer, actionExecuteUseCase,
				log.Params{
					constant.Key:        fmt.Sprintf(`%s_%s_%s_%s`, requestParam.F1, requestParam.F2, requestParam.F3, requestParam.F4),
					constant.EntityType: entityType,
				}))
		ginCTX.JSON(ErrorResponse.BadRequest.Value(), ErrorResponse.Response{
			Status:  ErrorResponse.BadRequest.Value(),
			Message: message,
		})
		return nil
	}

	qry := h.mapper.RequestToQuery(*requestParam)

	opportunities, errorUseCase := h.useCase.Execute(ginCTX, qry)

	if errorUseCase != nil {
		message := errorOpportunities.GetMessageWithTagParams(
			log.NewTagParams(layer, actionExecuteUseCase,
				log.Params{
					constant.Key:        fmt.Sprintf(`%s_%s_%s_%s`, requestParam.F1, requestParam.F2, requestParam.F3, requestParam.F4),
					constant.EntityType: entityType,
				}))
		ginCTX.JSON(ErrorResponse.InternalError.Value(), ErrorResponse.Response{
			Status:  ErrorResponse.InternalError.Value(),
			Message: message,
		})
		return nil
	}

	ginCTX.JSON(http.StatusOK, h.mapper.EntityToResponse(opportunities))
	return nil
}
