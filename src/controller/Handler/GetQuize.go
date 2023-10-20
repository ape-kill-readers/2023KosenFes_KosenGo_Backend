package handler

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuha-yuha/2023KosenFes_KosenGo_Backend.git/model"
)

var Quizzes []model.Quize

func QuizeFetch(ctx *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(Quizzes))

	//quize, err := model.QuizeSelect(index)
	quize, err := model.QuizePop(index, &Quizzes)

	if err != nil {
		ctx.Abort()
		panic(err)
	}

	ctx.JSON(200, quize)

}
