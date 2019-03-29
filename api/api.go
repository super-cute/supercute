package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Api struct {
	router *gin.Engine
	front  *gin.Engine
}

func New() *Api {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(allowCors())

	front := gin.New()
	front.Use(gin.Recovery())

	api := &Api{
		router: router,
		front:  front,
	}
	api.init()

	return api
}

func (api *Api) init() {

	ag := api.router.Group("/api/v0")
	ag.GET("/time", api.Time)

	jg := ag.Group("/json")
	jg.POST("/escape", api.UnMarshal)
	jg.POST("/unescape", api.JsonUnescape)

	api.front.StaticFS("/", http.Dir("../../web"))
}

func (api *Api) Start() error {
	hs := &http.Server{
		Addr:           fmt.Sprintf(":%d", 9527),
		Handler:        api.router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	hs2 := &http.Server{
		Addr:           fmt.Sprintf(":%d", 80),
		Handler:        api.front,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go hs.ListenAndServe()

	return hs2.ListenAndServe()
}

func allowCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, UPDATE, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers",
			"Authorization, origin, content-type, accept, x-requested-with, token")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	}
}
