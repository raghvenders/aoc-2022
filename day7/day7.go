package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	stack []string = make([]string, 0)
  dirs = make(map[string]*dir)
  dirSize = make(map[string]int)
)

type dir struct {
	child map[string]*dir
  name string
	size  int
}


func (bb *dir) String() string {
     dirs := make([]string,0)
     for keys := range bb.child{
     dirs = append(dirs, keys)
  } 
    return fmt.Sprintf("%s, %d, %v",bb.name, bb.size, dirs)
}


func SpaceOnDevice(file string) {
	fmt.Println("Day7 :: Space on Device")

	input, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
  lineNum := 0
	for scanner.Scan() {
		input := scanner.Text()
		processCommands(input, lineNum)
    lineNum += 1
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
   
/*
  a := &dir{
   child: map[string]*dir{},
   size: 0,
   name: "a",
  }
  dirs["/"] = a

  fmt.Printf("%v", dirs)
*/
  fmt.Printf("%v \n",dirs)
  fmt.Printf("%v \n",stack)
 fmt.Printf("%v \n",dirSize)

}

func processCommands(line string, num int) {

	re := regexp.MustCompile(`^\$ (cd (\w+|\.\.|\/)|ls)$`)
  fmt.Println("Line :", line, num) 
	if string(re.Find([]byte(line))) == line && !strings.Contains(line, "ls") {
		arr := strings.Split(line, " ")
		if arr[len(arr)-1] == ".." {
      
      last := stack[len(stack)-1]
      if dirs[last] != nil && dirs[last].size  <= 100000{
         dirSize[last] = dirs[last].size  
      }
			stack = stack[:len(stack)-1]
      if len(stack) >0 {
         last = stack[len(stack)-1]
         dirs[last].size +=  dirSize[last] 
      } 

		} else {
			stack = append(stack, arr[len(arr)-1])
		}
    if dirs[arr[len(arr)-1]] == nil{
       dirs[arr[len(arr)-1]] = &dir{name: arr[len(arr)-1],size: 0, child: make(map[string]*dir)}
    }
    fmt.Println( " --- stack", stack)
	} else {
     dirName := strings.Split(line, " ")
     parent := stack[len(stack)-1]
     if strings.Index(line, "dir") != -1{
       d := &dir{
        child: nil,
        size: 0,
        name: dirName[1],
      }
      dirs[parent].child[dirName[1]] =  d
      }else{
      x,_ := strconv.Atoi(dirName[0])
      fmt.Println("file", x)
      f := &dir{
        child: nil,
        size: x,
        name: dirName[1],
      }
      dirs[parent].child[dirName[1]] = f
      dirs[parent].size += x
      }
     fmt.Printf("--- process ::: %v \n", dirs)
    }
}
    
