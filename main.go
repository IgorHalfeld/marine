package main

import (
	"os"

	"github.com/marine/app"
	"github.com/marine/config"
)

func main() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "firebase.json")
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":80")
}
