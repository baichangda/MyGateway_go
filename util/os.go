package util

import (
	"os"
	"os/signal"
)

func ExitOnKill() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		for v := range c {
			Log.Infof("exit signal %+v", v)
			os.Exit(0)
		}
	}()
}
