package main

import (
	"Insurance/internal/api_db/redis_db"
	"Insurance/internal/api_db/reindexer_db"
	"Insurance/internal/handler"
	"Insurance/internal/model"
	"Insurance/internal/service"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/restream/reindexer/v3"
	_ "github.com/restream/reindexer/v3/bindings/cproto"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"log"
	"resenje.org/logging"
)

func main() {
	viper.SetConfigFile("config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		logging.Info(fmt.Errorf("config: %v", err))
	}

	db := reindexer.NewReindex(fmt.Sprintf("%s://%s:%s/%s",
		viper.GetString("db.scheme"),
		viper.GetString("db.hostname"),
		viper.GetString("db.port"),
		viper.GetString("db.path"),
	))

	if err := db.Status().Err; err != nil {
		logging.Info("reindexer connection: " + err.Error())
	}

	if err := db.OpenNamespace("employees", reindexer.DefaultNamespaceOptions(), model.EmployeeItem{}); err != nil {
		logging.Info("open namespace employees: " + err.Error())
	}

	logging.Info("Connection to reindexer DB successful!")

	rdb7 := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", viper.GetString("redis.ip"), viper.GetString("redis.port")),
		Password: "",
		DB:       11,
	})

	if err := rdb7.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("redis db11 connection: %v", err)
	}

	logging.Info("Connection to redis DB11 successful!")

	rdb8 := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", viper.GetString("redis.ip"), viper.GetString("redis.port")),
		Password: "",
		DB:       12,
	})

	if err := rdb8.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("redis db12 connection: %v", err)
	}

	logging.Info("Connection to redis DB12 successful!")

	employeeApiDB := reindexer_db.NewEmployeeApiDB(db)
	authRedisApiDB := redis_db.NewRedisApiDB(rdb7, rdb8)

	s := service.NewService(employeeApiDB, authRedisApiDB)

	h := handler.NewHandler(s)

	server := fasthttp.Server{
		Handler: h.InitRoutes,
	}

	if err := server.ListenAndServe(viper.GetString("http.port")); err != nil {
		log.Fatalf("start server: %v", err)
	}
}
