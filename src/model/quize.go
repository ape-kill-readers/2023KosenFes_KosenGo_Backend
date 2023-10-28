package model

import (
	"errors"
	"log"
	"os"

	"github.com/jszwec/csvutil"
)

type Quize struct {
	Id      int    `csv:"id" json:"-"`
	Content string `csv:"que" json:"que"`
	Answer  string `csv:"ans" json:"ans"`
}

func GetQuizeList(mode string) []Quize {
	var quizzes []Quize

	filename := "assets/" + mode + "Quize.csv"
	log.Println(mode)

	fb, _ := os.ReadFile(filename)

	_ = csvutil.Unmarshal(fb, &quizzes)

	log.Println(mode)

	return quizzes
}

/*func QuizeSelect(index int) (Quize, error) {

	var err error
	quizzes := GetQuizeList()

	if index >= len(quizzes) {
		err = errors.New("Error:Unexpected index")
	}
	return quizzes[index], err

}*/

func QuizePop(index int, quizzes *[]Quize) (Quize, error) {
	var err error
	if index >= len((*quizzes)) {
		err = errors.New("Error:Unexpected index")
	}
	PopedQuize := (*quizzes)[index]

	(*quizzes) = append((*quizzes)[:index], (*quizzes)[index+1:]...)
	return PopedQuize, err
}
