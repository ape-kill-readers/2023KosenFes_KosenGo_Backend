package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yuha-yuha/2023KosenFes_KosenGo_Backend.git/model"
)

func ClearQuizeProgress(ctx *gin.Context) {
	Quizzes = model.GetQuizeList()
	ctx.Status(200)
}
