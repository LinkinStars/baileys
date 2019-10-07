package router

func InitRouter(port string) {
	r := gin.New()

	v1 := r.Group("/gtw/api/v1")
	
	{{range $i, $v := .}}
	// {{$v.Comment}}
    v1.POST("/{{$v.UnderlineName | UnderlineStr2Strikethrough}}", controller.Add{{$v.UpperCamelName}})
    v1.DELETE("/{{$v.UnderlineName | UnderlineStr2Strikethrough}}", controller.Remove{{$v.UpperCamelName}})
    v1.PUT("/{{$v.UnderlineName | UnderlineStr2Strikethrough}}", controller.Update{{$v.UpperCamelName}})
    v1.GET("/{{$v.UnderlineName | UnderlineStr2Strikethrough}}/:id", controller.Get{{$v.UpperCamelName}})
    v1.GET("/{{$v.UnderlineName | UnderlineStr2Strikethrough}}s", controller.Get{{$v.UpperCamelName}}s)
    v1.GET("/{{$v.UnderlineName | UnderlineStr2Strikethrough}}s/page", controller.Get{{$v.UpperCamelName}}sWithPage)
    
    {{end}}

	err := r.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
