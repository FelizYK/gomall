package main

import (
	"net/http"
	"os"

	"github.com/FelizYK/gomall/frontend/conf"
	"github.com/FelizYK/gomall/frontend/handler"
	"github.com/FelizYK/gomall/frontend/router"
	"github.com/FelizYK/gomall/frontend/rpc"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	registerMiddleware(r)

	rpc.InitEtcd()
	defer rpc.CloseEtcd()
	rpc.InitClient()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.LoadHTMLGlob("templates/*")
	r.Static("assets", "./assets")

	r.GET("/", handler.Home)
	router.Auth(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func registerMiddleware(r *gin.Engine) {
	// redis session
	store, err := redis.NewStore(100, "tcp", conf.GetConf().Redis.Address, os.Getenv("REDIS_PASSWORD"), []byte(os.Getenv("SESSION_SECRET")))
	if err != nil {
		panic(err)
	}
	store.Options(sessions.Options{MaxAge: 86400})
	r.Use(sessions.Sessions("feliz-shop", store))
}
