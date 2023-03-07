package day10

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)



var (
  cycle = []int{20, 60, 100, 140, 180, 220}
  signalStrength = make(map[int]int)
  result int
)


func CathodeRayTube(file string) int{
  fmt.Println("Cathode Ray tube")

 input, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
  X := 1
  current := 0
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
    switch input[0] {
    case "noop":{
      current += 1
      if len(cycle) >0 && current >= cycle[0]{
         signalStrength[cycle[0]] = cycle[0] * X
        if cycle[0] == 1{
           cycle = []int{}
        }else {
          cycle = cycle[1:]
      }
      }
    }
    case "addx":{
      current += 2
     if len(cycle) >0 && current >= cycle[0]{
         signalStrength[cycle[0]] = cycle[0] * X
        if cycle[0] == 1{
           cycle = []int{}
        }else {
          cycle = cycle[1:]
      }
      }
      
      ins, _ := strconv.Atoi(input[1])
      X += ins
    }

    default:{
      log.Fatal("Invalid instruction")
    }
      
    }
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


  fmt.Println(signalStrength , X, current)
  for _,v := range signalStrength{
      result += v
  }
  return result
}
