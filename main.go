package main

import (
	"MyGateway_go/util"
	"github.com/pkg/errors"
	"io"
)

func main() {
	err1 := io.EOF
	err2 := errors.WithStack(err1)
	util.Log.Errorf("%+v", err2)
}
