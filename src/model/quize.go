package model

import (
	"errors"
	"os"

	"github.com/jszwec/csvutil"
)

type Quize struct {
	Id      int    `csv:"id" json:"-"`
	Content string `csv:"que" json:"que"`
	Answer  string `csv:"ans" json:"ans"`
}

func GetQuizeList() []Quize {
	var quizzes []Quize

	fb, _ := os.ReadFile("quize.csv")

	_ = csvutil.Unmarshal(fb, &quizzes)

	return quizzes
}

func QuizeSelect(index int) (Quize, error) {

	var err error
	quizzes := GetQuizeList()

	if index >= len(quizzes) {
		err = errors.New("Error:Unexpected index")
	}
	return quizzes[index], err

}

func QuizePop(index int, quizzes *[]Quize) (Quize, error) {
	var err error
	if index >= len((*quizzes)) {
		err = errors.New("Error:Unexpected index")
	}
	PopedQuize := (*quizzes)[index]

	(*quizzes) = append((*quizzes)[:index], (*quizzes)[index+1:]...)
	return PopedQuize, err
}
