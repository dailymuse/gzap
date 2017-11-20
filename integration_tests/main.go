package main

import (
	"github.com/dailymuse/gzap"
)

func main() {
	if err := gzap.InitLogger(); err != nil {
		panic(err)
	}

	gzap.Logger.Info("LOL")
}
