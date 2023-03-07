package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)


var (
    stack [][]string
    moveOps = make([][3]int,0)
  )


func moveCrates(stack [][] string, moves [][3]int){
  
  for _, v := range moves{
    for j:=0; j < v[0]; j++{
      fromCol := stack[v[1]-1]

      crate := fromCol[len(fromCol)-1]

      stack[v[2]-1] = append(stack[v[2]-1], crate)

      stack[v[1]-1] = fromCol[:len(stack[v[1]-1])-1] 

    }
  }

  
  result := ""
  for _, v := range stack{
     result += v[len(v)-1]
  }

 log.Println(result)


}


func SupplyStacks(file string){
   start := time.Now()
   ops := func(s string, a ...string) [3]int {
		for _, v := range a {
			//fmt.Println(i, v)
			s = strings.Replace(s, v, "", 1)
		}
		reArr := [3]int{}

		for i, v := range strings.Split(s, " ") {
			reArr[i], _ = strconv.Atoi(v)
		}

		return reArr
	}

	input, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
  
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
    line := scanner.Text()

    if strings.Index(line, "move") == 0{
        moveOps = append(moveOps, ops(line, "move ", "from ", "to "))
        continue
    }

    if stack == nil{

      stack = make([][]string, len(line)/4+1)


    }

    rowCrates := strings.Split(line, "")


    for i:=0; i < len(rowCrates);i+=4{
      if rowCrates[i] == "["{
        column := i/4
        if stack[column] == nil{
            stack[column] = make([]string, 0)
        }
        //crate :=  rowCrates[i] + rowCrates[i+1]+ rowCrates[i+2]
        stack[column] = append(stack[column], rowCrates[i+1])
         }
  }
  }



  for _, v := range stack{

    for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
        v[i], v[j] = v[j], v[i]
    }
  }

if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}



  moveCrates(stack, moveOps)
   fmt.Println(time.Now().Sub(start))
  }
