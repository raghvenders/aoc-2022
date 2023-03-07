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

func CampCleanup(file string) {

  start := time.Now()
	input, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
  sum := 0
	for scanner.Scan() {
    camp := strings.Split(scanner.Text(), ",")
    if len(camp) == 2{
     first := strings.Split(camp[0],"-")
     second := strings.Split(camp[1], "-")
     if len(first) !=2 || len(second) != 2{
       log.Fatal("Not Enough cap assgignments range") 
      }
      b1,_ := strconv.Atoi(second[0])
      b2,_ := strconv.Atoi(second[1])
      a1,_ := strconv.Atoi(first[0])
      a2,_ := strconv.Atoi(first[1])

      switch {
       case a1 == b1 || a2 == b2:
        sum += 1
        break
       case b2 > a2 && b1 < a1:
        sum += 1
        break
       case a2 > b2 && a1 < b1:
        sum += 1
        break
       }
    }else{
      log.Fatal("Not enough Camp Assignments")
    }

	}

  log.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

  fmt.Println(time.Now().Sub(start))

}
