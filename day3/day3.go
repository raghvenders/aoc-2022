package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
)


var charItems [52]uint8
func sumItemPriority(a string) uint8{

  charItems := [52]uint8{}

  //a := "vJrwpWtwJgWrhcsFMMfFFhFp"
  //a := "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"
  //a := "PmmdzqPrVvPwwTWBwg"
 //A-Z : 65-90 - 38
 //a-z : 97-122 -96
  half := len(a)/2
  for i:=0; i< half; i++{
    //fmt.Println(reflect.TypeOf(a[i]))
    var itemIndex uint8
    switch {
    case a[i] >= 65 && a[i] <=90:
      itemIndex = a[i] - 38
      break
    case a[i] >= 97 && a[i] <= 122:
     itemIndex = a[i] - 96
      break
    default:
     log.Println(fmt.Errorf("This is not a valid Ascill English Alphabet %d", a[i]))
     return 0
     } 
        
     charItems[itemIndex-1] = 1
  }

   for i:=len(a)-1 ; i >=half; i--{
    var result uint8
    if a[i] >= 65 && a[i] <= 90{
         result = a[i] - 38
    }else if a[i] >= 97 && a[i] <=122{
         result = a[i] - 96     
    } 
    if charItems[result-1] == 1{
       return result
    }
  } 




return 0
}

func RucksackReorg(file string){
input, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

  priorities := 0
	for scanner.Scan() {
       priorities += int(sumItemPriority(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(priorities)
}
