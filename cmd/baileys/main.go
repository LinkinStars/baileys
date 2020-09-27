package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/LinkinStars/baileys/internal/cache"
	"github.com/LinkinStars/baileys/internal/handle"
	"github.com/LinkinStars/baileys/internal/util"
)

func init() {
	flag.StringVar(&cache.ConfPath, "c", "./conf/conf.yml", "default config path")
	flag.StringVar(&cache.WebPort, "p", "5272", "default web port")
}

func main() {
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "./static")

	router.GET("/", handle.LoadingIndex)
	router.POST("/gen", handle.GenCode)

	err := util.OpenBrowser("http://127.0.0.1:" + cache.WebPort + "/")
	if err != nil {
		log.Print("open browser error : ", err)
	}

	err = router.Run(":" + cache.WebPort)
	if err != nil {
		panic(err)
	}
}
