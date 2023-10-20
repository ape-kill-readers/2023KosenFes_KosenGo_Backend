package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	handler "github.com/yuha-yuha/2023KosenFes_KosenGo_Backend.git/controller/Handler"
	"github.com/yuha-yuha/2023KosenFes_KosenGo_Backend.git/middleware"
)

func Get() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},

		AllowMethods: []string{
			"GET",
		},
	}))

	r.GET("/QuizeFetch", middleware.IsQuizeListEmpty(), handler.QuizeFetch)
	r.GET("/ClearQuizeProgress", handler.ClearQuizeProgress)

	return r
}
