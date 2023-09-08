package main

import (
	"MyGateway_go/util"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io"
)

func init() {
	util.InitLog()
}

func main() {
	err1 := io.EOF
	err2 := errors.WithStack(err1)
	zap.S().Errorf("%+v", err2)
}
