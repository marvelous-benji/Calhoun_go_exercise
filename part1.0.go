

package main


import (
	"fmt"
	"io/ioutil"
	"strings"
)


type exam struct{
	question string
	answer string
	user_answer string
}


func readfile(path string) []string{

	f, err := ioutil.ReadFile(path)
	if err != nil{
		panic(err)
	}
	text := string(f)
	return strings.Split(text, "\n")
}

func prepareExam() []exam {
	marking_array := []exam{}
	exam_struct := exam{}
	marking_scheme := readfile("problem.csv")
	for _, QandA_pairs := range marking_scheme{
		QandA_array := strings.Split(QandA_pairs, ",")
		if len(QandA_array) == 2{
			exam_struct.question = QandA_array[0]
			exam_struct.answer = QandA_array[1]
			marking_array = append(marking_array, exam_struct)
		}
	}
	return marking_array
}


func askQuestion(questions []exam) []exam{
	for i, _ := range questions{
		fmt.Printf("Problem %d# %s = ", i+1, questions[i].question)
		fmt.Scanf("%s", &questions[i].user_answer)
	}
	return questions
}


func main(){
	response := askQuestion(prepareExam())
	score := 0
	question_num := 0
	for _, result := range response{
		if result.answer == result.user_answer{
			score++
		}
		question_num++
	}
	fmt.Println("You scored",score,"out of",question_num)

}