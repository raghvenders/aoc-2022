package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

var topCal = [3]int{0,0,0}
func SummedValues(c, quit chan int){
    defer wg.Done()
  for{
   select {
    case v := <-c:
       // fmt.Println(topCal)
        if v > topCal[0]{
          topCal[2],topCal[1], topCal[0] = topCal[1], topCal[0], v
        }else if v > topCal[1]{
          topCal[2], topCal[1] = topCal[1], v
        }else if v > topCal[2]{
          topCal[2] = v
        }
     case <-quit:
			 // fmt.Println("quit")
			  return
	  }
  }

}

func CalorieCounting(file string) [3]int{
 
  wg.Add(1)

  dataFile, err := os.Open(file)

  if err != nil{
    err := fmt.Errorf("Unable to open the file %w", err)
    fmt.Println(err.Error())
  }

  s := bufio.NewScanner(dataFile)

  c := make(chan int)
  quit := make(chan int)
  calories := 0
  go SummedValues(c,quit)
  for s.Scan() {
   if s.Text() == "" {
      c <- calories
      calories = 0
    } else {
        cal, err := strconv.Atoi(s.Text())
        if err != nil{
        fmt.Println("%w",err)
      }
        calories += cal
      
    }

      
  }

if calories >= 0{
   c <- calories 
  }
  
//time.Sleep(1 * time.Millisecond)
quit <- 0

  wg.Wait()


  return topCal

}



