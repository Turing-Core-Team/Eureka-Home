package log

import (
	"EurekaHome/internal/platform/constant"
	"fmt"
	"sort"
	"strings"
)


type LogsMessage string

type TagParams struct {
	layer   string
	action  string
	params  Params
}
type Params map[string]interface{}

func NewTagParams( layer, action string, otherParams Params) TagParams {
	return TagParams{ layer: layer, action: action, params: otherParams}
}

func (lm LogsMessage) GetMessage() string {
	return string(lm)
}

func (lm LogsMessage) GetMessageWithTagParams(tagParams TagParams) string {
	msg := concatStrWithTag(lm.GetMessage(), tagParams)

	return msg
}

func concatStrWithTag(str string, tagParams TagParams) string {
	keys := make([]string, 0)
	sb := strings.Builder{}

	sb.WriteString(str)
	sb.WriteString(fmt.Sprintf(" %v:%v", constant.Layer, tagParams.layer))
	sb.WriteString(fmt.Sprintf(" %v:%v", constant.Action, tagParams.action))

	for k := range tagParams.params {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		sb.WriteString(fmt.Sprintf(" %v:%v", k, tagParams.params[k]))
	}

	return sb.String()
}