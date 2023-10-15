package model

import (
	"testing"

	"github.com/yuha-yuha/2023KosenFes_KosenGo_Backend.git/model"
)

func TestPop(t *testing.T) {
	quizzes := model.GetQuizeList()
	quizzesConfirm := quizzes

	for _, quizeConfirm := range quizzesConfirm {
		if quize, _ := model.QuizePop(0, &quizzes); quize != quizeConfirm {
			t.Errorf("quizePop: want %s actual %s", quizeConfirm.Content, quize.Content)
		}
	}

}
