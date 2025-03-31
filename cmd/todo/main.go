package main

import (
	"fmt"
	"log"
	"os"
	"time"

	kitLog "github.com/go-kit/log"
	todo "github.com/ppeymann/todo_be.git"
	"github.com/ppeymann/todo_be.git/cmd/todo/pkg"
	"github.com/ppeymann/todo_be.git/env"
	"github.com/ppeymann/todo_be.git/server"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	now := time.Now()

	base := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).Unix()
	start := time.Date(now.Year(), now.Month(), now.Day(), 7, 35, 0, 0, time.UTC).Unix()
	end := time.Date(now.Year(), now.Month(), now.Day(), 23, 30, 0, 0, time.UTC).Unix()

	fmt.Println("date:", base, "start:", start, "end:", end)

	// initialization configuration from environment variable
	config, err := todo.NewConfiguration("")
	if err != nil {
		log.Fatal(err)
		return
	}

	env := env.GetStringDefault("DSN", "")

	// connect to Db server
	db, err := gorm.Open(pg.Open(env), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatal(err)
		return
	}

	// configuring logger
	var logger kitLog.Logger
	logger = kitLog.NewJSONLogger(kitLog.NewSyncWriter(os.Stderr))
	logger = kitLog.With(logger, "ts", kitLog.DefaultTimestampUTC)

	// Service Logger
	sl := kitLog.With(logger, "component", "http")

	// Server instance
	svc := server.NewServer(sl, config)

	// --------  initializing service  --------

	// AccountService
	pkg.InitAccountService(db, sl, config, svc)

	// @TodoService
	pkg.InitTodoService(db, sl, config, svc)

	// Listen and serve...
	svc.Listen()

}
