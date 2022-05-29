package sheets

import (
	"context"
	"errors"
	"sync"
)

type ClientLocal struct {
	store sync.Map
}

var (
	singleInstance ClientLocal
	once           sync.Once
)

func NewSheetsClientMock() ClientLocal {
	once.Do(func() {
		singleInstance = ClientLocal{}
	})

	return singleInstance
}

func (c *ClientLocal) Read(_ context.Context, path, spreadsheetId, readRange string) ([][]interface{}, error) {

	if v, ok := c.store.Load(readRange); ok {
		item := v.([][]interface{})
		return item, nil
	}

	return nil, errors.New("local sheets item not found")
}
