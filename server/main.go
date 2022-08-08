package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/name/pjName/server/api"
	"github.com/name/pjName/server/command"
	"github.com/name/pjName/server/pkg/configutil"
	"github.com/name/pjName/server/pkg/dotenvutil"
)

func main() {

	// Asia/Tokyo
	time.Local = time.FixedZone("Local", 9*60*60)

	// dotenv
	err := dotenvutil.Init()
	if err != nil {
		log.Fatalf("dotenv error: %+v\n", err)
	}

	// server init
	err = configutil.ServerInit()
	if err != nil {
		log.Fatalf("server error: %+v\n", err)
	}

	// db init
	err = configutil.DBInit()
	if err != nil {
		log.Fatalf("db error: %+v\n", err)
	}
	defer configutil.DBClose()

	// router
	r := gin.Default()
	api.SetHandlers(r)

	// command
	if len(os.Args) >= 2 {
		name := os.Args[1]
		err = command.Run(name)
		if err != nil {
			log.Fatalf("%+v\n", err)
		}
		return
	}

}
