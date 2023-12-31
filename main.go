package main

import (
	"fmt"
	"os"
	"rag-sun/aoc2022/day1"
	"rag-sun/aoc2022/day10"
	"rag-sun/aoc2022/day11"
	"rag-sun/aoc2022/day12"
	"rag-sun/aoc2022/day2"
	"rag-sun/aoc2022/day3"
	"rag-sun/aoc2022/day4"
	"rag-sun/aoc2022/day5"
	"rag-sun/aoc2022/day6"
	"rag-sun/aoc2022/day7"
	"rag-sun/aoc2022/day8"
	"rag-sun/aoc2022/day9"
	"strconv"
	"strings"
	"time"
)

const FOLDER_PREFIX = "./data/"
const FILE_EXT = ".txt"

func main() {
	//first := time.Now()
	var inputArgs int
	var err error
	if len(os.Args) > 1 {
		inputArgs, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Invalid Input Args %d, valid are 1-24", inputArgs)
			os.Exit(1)
		}
	}

	inputFile := FOLDER_PREFIX + "Day" + strings.ToLower(strconv.Itoa(inputArgs)) + FILE_EXT

	switch "Day" + strconv.Itoa(inputArgs) {
	case "Day1":
		start := time.Now()
		//fmt.Printf("Top 3 calories %v",day1.CalorieCounting(inputFile))
		day1.Puzzel1()
		fmt.Print(time.Now().Sub(start))
	case "Day2":
		//fmt.Printf("RPC Winner Stat %d", day2.RockPaperSisWin(inputFile))
		end := time.Now()
		day2.RockPaperSisWin(inputFile)
		fmt.Println(time.Now().Sub(end))

		second := time.Now()
		day2.Puzzel2()
		fmt.Println(time.Now().Sub(second))
	case "Day3":
		end := time.Now()
		day3.RucksackReorg(inputFile)
		fmt.Println(time.Now().Sub(end))

		second := time.Now()
		day3.RuckSackPuzzel3()
		fmt.Println(time.Now().Sub(second))
	case "Day4":
		//end := time.Now()
		// day4.CampCleanup(inputFile)
		// fmt.Println(time.Now().Sub(end))

		//      second := time.Now()
		go day4.PuzzelDay4()
		go day4.CampCleanup(inputFile)
		//    fmt.Println(time.Now().Sub(second))

		//time.Sleep(2000 * time.Millisecond)
	case "Day5":
		go day5.SupplyStacks(inputFile)
		go day5.PuzzelDay5()

	case "Day6":
		day6.Devices(inputFile)
		day6.Puzzel6()
	case "Day7":
		day7.SpaceOnDevice(inputFile)
	case "Day8":
    end := time.Now()

	  result :=	day8.TreeTopHouse(inputFile)

    fmt.Println("Tree view ::", result)
    fmt.Println(time.Now().Sub(end))
  case "Day9":
    end := time.Now()
    day9.RopeBridge(inputFile)
    fmt.Println(time.Now().Sub(end))
	case "Day10":
		result := day10.CathodeRayTube(inputFile)
		fmt.Printf("Cathode Ray tube :: Sum of signal Strength %d", result)
	case "Day11":
    end := time.Now()
    day11.MonkeyInMiddle(inputFile)
    fmt.Println(time.Now().Sub(end))
  case "Day12":
    day12.HillClimbing(inputFile)

	default:
		fmt.Println("Roma is not built yet !!!")

	}

	//time.Sleep(2 * time.Second)
	//end := time.Now()
	//fmt.Println(end.Sub(first))

	//second := time.Now()
	//day1.Puzzel1()
	//sendd := time.Now()

	//fmt.Println(sendd.Sub(second))
}
