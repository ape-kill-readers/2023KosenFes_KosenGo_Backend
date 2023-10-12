package router

import (
	"github.com/gin-gonic/gin"
	handler "github.com/yuha-yuha/2023KosenFes_KosenGo_Backend.git/controller/Handler"
)

func Get() *gin.Engine {
	r := gin.Default()

	r.GET("/QuizeFetch", handler.QuizeFetch)

	return r
}
