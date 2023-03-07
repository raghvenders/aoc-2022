package day8

import (
	"bufio"
	"log"
	"os"
)

func TreeTopHouse(file string) int{

	//a := [][]int{{3, 0, 3, 7, 3}, {2, 5, 5, 1, 2}, {6, 5, 3, 3, 2}, {3, 3, 5, 4, 9},{3, 5, 3, 9,0}}

  a := make([][]int,0)


 input, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
  lineNum := 0
	for scanner.Scan() {
		input := scanner.Text()
    row := make([]int,0)
    for _, v := range input{
      t := int(v - '0')
      row = append(row, t)
    }
    a = append(a,row)
    lineNum += 1
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}



	// 30373
	// 25512
	// 65332
  // 33549
  // 35390
  

  // 2(r) + 2(c-2)
  /*
 if len(a) <= 2{
    return len(a[0]) * len(a)
  }

if len(a) > 1 && len(a[0][:]) ==1{
   return len(a)
  }

*/

 treeView := 2 * (len(a)) + 2 * (len(a[0])-2)
 //treeView := 0
 
   

	//fmt.Printf("%d %d", len(a), len(a[0]))
	//column, row := len(a), len(a[0])
	//size := column + row
	//i, j := 1, 1
	row, column := len(a)-1, len(a[0])-1
	//k, l, m, n := i-1, i+1, j-1, j+1
  for i:=1 ; i <= len(a)-2; i++ {
    for j:= 1 ; j <= len(a[0])-2; j++ {
			k := 1
			r, l, u, d := true, true, true, true
			for {
				if r && (j+k <= column && a[i][j] <= a[i][j+k]) {
					r = false
				}
				if l && (j-k >= 0 && a[i][j] <= a[i][j-k]) {
					l = false
				}
				if d && (i+k <= row && a[i][j] <= a[i+k][j]) {
					d = false
				}
        if u && (i-k >= 0 && a[i][j] <= a[i-k][j]) {
          u = false
        }

        if !r && !l && !u && !d{
          break
        }

        if (r && j+k == column) || (l && j-k == 0) || (u && i-k == 0) || (d && i+k == row){
          treeView++
          break
        }

				k++
			}
		}
	}

 // fmt.Printf("Tree View :: %d", treeView)
 return treeView
}
