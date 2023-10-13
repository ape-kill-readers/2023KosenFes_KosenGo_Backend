package handler

import (
	"log"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuha-yuha/2023KosenFes_KosenGo_Backend.git/model"
)

func QuizeFetch(ctx *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(model.GetQuizeList()))

	quize, err := model.QuizeSelect(index)

	log.Println(ctx.Request.Header)

	if err != nil {
		ctx.Abort()
		panic(err)
	}

	ctx.JSON(200, quize)

}
