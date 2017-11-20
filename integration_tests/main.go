package main

import (
	"flag"
	"os"

	"github.com/dailymuse/gzap"
)

func main() {
	// Set the proper variables
	var graylogEnv = flag.String("graylog_env", "0", "Graylog environment variable")
	flag.Parse()

	if err := os.Setenv("GRAYLOG_ENV", *graylogEnv); err != nil {
		panic(err)
	}

	if err := gzap.InitLogger(); err != nil {
		panic(err)
	}

	gzap.Logger.Info("LOL")
}
