package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func contains(a1, a2, b1, b2 int) bool {
	return a1 <= b1 && a2 >= b2 || b1 <= a1 && b2 >= a2
}

func PuzzelDay4(){

  start := time.Now()
	input, err := os.Open("./data/day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	answer := 0

	for scanner.Scan() {
		pairs := strings.Split(scanner.Text(), ",")
		pair1 := strings.Split(pairs[0], "-")
		pair2 := strings.Split(pairs[1], "-")

		a1, _ := strconv.Atoi(pair1[0])
		a2, _ := strconv.Atoi(pair1[1])
		b1, _ := strconv.Atoi(pair2[0])
		b2, _ := strconv.Atoi(pair2[1])

		if contains(a1, a2, b1, b2) {
			answer += 1
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(answer)

  fmt.Println("---", time.Now().Sub(start))
}
