package util

import (
	"go.uber.org/zap"
	"os"
	"os/signal"
)

func ExitOnKill(fnOnExit func() error) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		for v := range c {
			zap.S().Infof("exit signal %+v", v)
			if fnOnExit != nil {
				err := fnOnExit()
				if err != nil {
					zap.S().Errorf("%+v", err)
				}
			}
			os.Exit(0)
		}
	}()
}
