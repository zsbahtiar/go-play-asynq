package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/zsbahtiar/go-play-asynq/client/config"
	"github.com/zsbahtiar/go-play-asynq/client/handler"
	"github.com/zsbahtiar/go-play-asynq/client/module/user"
	"log"
)

func main() {
	cfg := config.Get()
	asynqCli := initAsynq(cfg.RedisAddress)
	userSvc := user.NewService(asynqCli)
	userHandler := handler.NewUserHandler(userSvc)

	router := gin.Default()
	router.POST("/create-user-csv", userHandler.CreateUsersCsv)

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}

func initAsynq(redisAddr string) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
}
