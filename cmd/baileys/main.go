package main

import (
	"embed"
	"flag"
	"html/template"
	"log"
	"net/http"

	"github.com/LinkinStars/baileys/internal/util"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/LinkinStars/baileys/internal/cache"
	"github.com/LinkinStars/baileys/internal/handle"
)

func init() {
	flag.StringVar(&cache.ConfPath, "c", "./conf/conf.yml", "default config path")
	flag.StringVar(&cache.WebPort, "p", "5272", "default web port")
}

//go:embed templates
var tmpl embed.FS

//go:embed static
var staticFS embed.FS

func main() {
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	t, err := template.ParseFS(tmpl, "templates/*.html")
	if err != nil {
		panic(err.Error())
	}
	router.SetHTMLTemplate(t)
	//router.LoadHTMLGlob("templates/*.html")
	router.StaticFS("/static/", http.FS(staticFS))

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", "")
	})
	router.GET("/converter/sql/code", handle.ConverterSql2Code)
	router.GET("/converter/go/pb", func(ctx *gin.Context) { ctx.HTML(http.StatusOK, "struct_2_pb.html", "") })
	router.GET("/converter/json/go", func(ctx *gin.Context) { ctx.HTML(http.StatusOK, "json_2_struct.html", "") })

	router.POST("/gen/sql/code", handle.GenCode)
	router.POST("/gen/go/pb", handle.ConvertGoStruct2PbMessage)

	if err := util.OpenBrowser("http://127.0.0.1:" + cache.WebPort + "/"); err != nil {
		log.Print("open browser error : ", err)
	}

	err = router.Run(":" + cache.WebPort)
	if err != nil {
		panic(err)
	}
}
