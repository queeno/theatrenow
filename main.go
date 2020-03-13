package main

import (
	"os"
	"theatrenow/pkg/app"
	"theatrenow/pkg/tools"
)

func main() {
	err := app.Run()

	if err != nil {
	  tools.Logger().Error(err.Error())
	  os.Exit(1)
	}
}