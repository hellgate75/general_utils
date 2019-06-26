package net

import (
	"errors"
	"fmt"
	"github.com/hellgate75/general_utils/net/common"
)

func NewServer(serverType common.ServerType) (common.Server, error) {
	return nil, errors.New(fmt.Sprintf("Feature <%v> Not implemented!!", serverType))
}
