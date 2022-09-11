package log_test

import (
	"EurekaHome/internal/platform/constant"
	"EurekaHome/internal/platform/log"
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

const (
	action               string          = "action"
	entityType           string          = "Account"
	errorCardInformation log.LogsMessage = "error consulting get opportunities."
	key                  string          = "eureka"
	layer                string          = "use_case"
)

func TestGetMessageWhenExecuteShouldReturnString(t *testing.T) {
	expectedMessage := "error consulting get opportunities."

	message := errorCardInformation.GetMessage()

	assert.Equal(t, expectedMessage, message)
}

func TestGetMessageWhenExecuteShouldReturnTagWhitParams(t *testing.T) {
	expectedMsj := "error consulting get opportunities.  layer:%s action:%s entity_type:%s key:%s"
	msj := errorCardInformation.GetMessageWithTagParams(
		log.NewTagParams(layer, action,
			log.Params{
				constant.EntityType: entityType,
				constant.Key:        key,
			}))

	assert.Equal(t, msj, fmt.Sprintf(
		expectedMsj,
		layer,
		action,
		entityType,
		key),
	)
}
