package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	handler "github.com/yuha-yuha/2023KosenFes_KosenGo_Backend.git/controller/Handler"
	"github.com/yuha-yuha/2023KosenFes_KosenGo_Backend.git/model"
)

func IsQuizeListEmpty() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(handler.Quizzes) == 0 {
			mode := ctx.Query("mode")
			handler.Quizzes = model.GetQuizeList(mode)
			log.Println("aaa")
		}

		log.Println("before")

		ctx.Next()

		if len(handler.Quizzes) == 0 {
			mode := ctx.Query("mode")
			handler.Quizzes = model.GetQuizeList(mode)
			log.Println("bbb")
		}

		log.Println("after")
	}
}
