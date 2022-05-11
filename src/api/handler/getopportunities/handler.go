package getopportunities

import (
	"EurekaHome/src/api/handler/getopportunities/contract"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	action                   string = "get_opportunities"
	actionExecuteUseCase     string = "execute_use_case"
	actionValidateParameters string = "validate_parameters"
	errorOpportunities       string = "error in the creation handler"
	entityType               string = "get_opportunities_by_description_and_tags"
	layer                    string = "handler_information"
)

type UseCase interface {
	Execute(ctx context.Context, GetOpportunities query.GetMiniCard) (entity.MiniCard, error)
}

type Mapper interface {
	RequestToQuery(request contract.URLParams) query.GetMiniCard
	EntityToResponse(entity entity.MiniCard) contract.InformationResponse
	QueryStringToBool(queryString string) bool
}

type ValidationParams interface {
	BindParamsAndValidation(obj interface{}, params gin.Params) error
}

type Handler struct {
	useCaseSelector  map[bool]UseCase
	mapper           Mapper
	validationParams ValidationParams
}

func NewHandler(
	useCase map[bool]UseCase,
	mapper Mapper,
	validationParams ValidationParams,
) *Handler {
	return &Handler{
		useCaseSelector:  useCase,
		mapper:           mapper,
		validationParams: validationParams,
	}
}

func (h Handler) Handler(c *gin.Context) {
	commonsHandlers.ErrorWrapper(h.handler, c)
}

func (h Handler) handler(ginCTX *gin.Context) *commonsApiErrors.APIError {
	ctx := commonsContext.RequestContext(ginCTX)
	defer commonsContext.Recorder(ctx).Segment(newrelic.GetSegmentName(layer, action)).End()

	isFullParam := ginCTX.DefaultQuery(isFull, "")
	requestParam := &contract.URLParams{}

	if err := h.validationParams.BindParamsAndValidation(requestParam, ginCTX.Params); err != nil {
		message := errorOpportunities.GetMessageWithTagParams(
			log.NewTagParams(commonsContext.UUID(ctx), layer, actionValidateParameters,
				log.Params{
					constant.EntityType: entityType,
				}))
		commonsContext.Logger(ctx).Error(message, err)
		return commonsApiErrors.NewBadRequest(err.Error())
	}

	qry := h.mapper.RequestToQuery(*requestParam)
	isFullBool := h.mapper.QueryStringToBool(isFullParam)

	information, errorUseCase := h.useCaseSelector[isFullBool].Execute(ctx, qry)

	if errorUseCase != nil {
		switch errorUseCase.(type) {
		case errorOrchestrator.PartialContent:
			ginCTX.JSON(http.StatusPartialContent, h.mapper.EntityToResponse(information))
			return nil
		default:
			message := errorOpportunities.GetMessageWithTagParams(
				log.NewTagParams(commonsContext.UUID(ctx), layer, actionExecuteUseCase,
					log.Params{
						constant.Key:        fmt.Sprintf(`%s_%v`, requestParam.SiteID, requestParam.UserID),
						constant.EntityType: entityType,
					}))
			commonsContext.Logger(ctx).Error(message, errorUseCase)
			return handler.MapperAPIError(errorUseCase)
		}
	}

	ginCTX.JSON(http.StatusOK, h.mapper.EntityToResponse(information))
	return nil
}
