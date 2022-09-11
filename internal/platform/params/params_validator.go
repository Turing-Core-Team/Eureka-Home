package params

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ParamValidation struct {
	ValidationParams map[string]ValidationParams
}

type ValidationParams interface {
	IsValid(value string) error
}

func NewParamValidation(validationParams map[string]ValidationParams) ParamValidation {
	return ParamValidation{ValidationParams: validationParams}
}

func (dp ParamValidation) BindParamsAndValidation(obj interface{}, params gin.Params) error {

	paramsMap := make(map[string]string)
	for key, valueParam := range dp.ValidationParams {
		param, _ := params.Get(key)
		if err := valueParam.IsValid(param); err != nil {
			return fmt.Errorf("error %s ", err.Error())
		}
		paramsMap[key] = param
	}


	jsonStr, err := json.Marshal(paramsMap)
	if err != nil {
		return fmt.Errorf("error Marshal map: %s", err.Error())
	}

	errUnmarshal := json.Unmarshal(jsonStr, &obj)
	if errUnmarshal != nil {
		return fmt.Errorf("error Unmarshal map: %s", errUnmarshal.Error())
	}
	return nil
}