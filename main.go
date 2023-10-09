package main

import (
	"MyGateway_go/util"
	"io"
	"time"
)

func main() {
	err1 := io.EOF
	go func() {
		util.Log.Errorf("%+v", err1)
	}()
	time.Sleep(time.Second)
}
