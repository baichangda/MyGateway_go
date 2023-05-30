package util

import (
	"go.uber.org/zap"
)

var Log = initLog()

func initLog() *zap.SugaredLogger {
	temp, err := zap.NewDevelopment()
	if err != nil {
		panic("log init fail:" + err.Error())
	}
	return temp.Sugar()
}
