package handler

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuha-yuha/2023KosenFes_KosenGo_Backend.git/model"
)

var quizzes []model.Quize

func init() {
	quizzes = model.GetQuizeList()
}

func QuizeFetch(ctx *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(quizzes))

	//quize, err := model.QuizeSelect(index)
	quize, err := model.QuizePop(index, &quizzes)

	if len(quizzes) == 0 {
		quizzes = model.GetQuizeList()
	}
	if err != nil {
		ctx.Abort()
		panic(err)
	}

	ctx.JSON(200, quize)

}
