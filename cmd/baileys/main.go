package main

import (
	"flag"
	"log"
	"net/http"

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

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", "")
	})
	router.GET("/converter/sql/code", handle.ConverterSql2Code)
	router.GET("/converter/go/pb", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "struct_2_pb.html", "")
	})

	router.POST("/gen/sql/code", handle.GenCode)
	router.POST("/gen/go/pb", handle.ConvertGoStruct2PbMessage)

	err := util.OpenBrowser("http://127.0.0.1:" + cache.WebPort + "/")
	if err != nil {
		log.Print("open browser error : ", err)
	}

	err = router.Run(":" + cache.WebPort)
	if err != nil {
		panic(err)
	}
}
