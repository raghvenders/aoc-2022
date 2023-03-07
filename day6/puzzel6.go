package day6

import (
	"io/ioutil"
	"log"
)

func Puzzel6(){
	input, err := ioutil.ReadFile("./data/day6.txt")
	if err != nil {
		log.Fatal(err)
	}

	answer := -1

	for i := 0; i < len(input)-4; i++ {
		if input[i] != input[i+1] && input[i] != input[i+2] && input[i] != input[i+3] &&
			input[i+1] != input[i+2] && input[i+1] != input[i+3] &&
			input[i+2] != input[i+3] {

			answer = i + 4
			break

		}
	}

	log.Println(answer)
}
