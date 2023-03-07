package day2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


var finalScore int
var SCORES = map[string]int{
  "AX" : 4,
  "AY" : 8,
  "AZ" : 3,
  "BX" : 1,
  "BY" : 5,
  "BZ" : 9,
  "CX" : 7,
  "CY" : 2,
  "CZ" : 6,
}
func RockPaperSisWin(file string) int{
  //fmt.Println("RPC Winner")
  dataFile, err := os.Open(file)

  if err != nil{
    err := fmt.Errorf("Unable to open the file %w", err)
    fmt.Println(err.Error())
  }
 scan := bufio.NewScanner(dataFile)

  for scan.Scan(){
    game := strings.Split(scan.Text(), " ")
    //fmt.Println(game)
   // if len(game) != 2 || len(game) == 0{
    //   fmt.Println("Ignore this game and Proceed")
     //  continue
    //}
    finalScore += SCORES[game[0]+game[1]]
  }

   fmt.Println(finalScore)

   return finalScore
} 
