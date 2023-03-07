package day9

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	positionMap = make(map[string]bool, 0)
)

const (
	start = "0_0"
)

func RopeBridge(file string) {
	log.Println("Day 9:: Rope Bridge")
	input, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	H := start
	L := start

	positionMap["0_0"] = true

	lVisit := 1

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()

		direcMove := strings.Split(line, " ")

		switch direcMove[0] {
		case "R", "L":
			i, _ = strconv.Atoi(direcMove[1])
			fp := strings.Split(H, "_")
			hx, _ := strconv.Atoi(fp[0])

			//_ = hx
			hy, _ := strconv.Atoi(fp[1])

			sp := strings.Split(L, "_")
			lx, _ := strconv.Atoi(sp[0])
			ly, _ := strconv.Atoi(sp[1])

			in := 0
			if direcMove[0] == "R" {
				in = 1
			} else {
				in = -1
			}
			for i > 0 {
				hx += in

				H = strconv.Itoa(hx) + "_" + strconv.Itoa(hy)

				if (hy-ly == 1 || hy-ly == -1) && (lx-hx == 2 || lx-hx == -2) {
					ly = hy
					lx = hx + (-in)
				} else if lx-hx == 2 || lx-hx == -2 {
					lx += in
				}

				L = strconv.Itoa(lx) + "_" + strconv.Itoa(ly)

				if positionMap[L] == false {
					lVisit++
					positionMap[L] = true
				}

				i--
			}

		case "U", "D":
			i, _ = strconv.Atoi(direcMove[1])
			fp := strings.Split(H, "_")
			hx, _ := strconv.Atoi(fp[0])

			//_ = hx
			hy, _ := strconv.Atoi(fp[1])

			//_ = hy
			sp := strings.Split(L, "_")
			lx, _ := strconv.Atoi(sp[0])
			ly, _ := strconv.Atoi(sp[1])

			in := 0
			if direcMove[0] == "U" {
				in = 1
			} else {
				in = -1
			}

			for i > 0 {
				hy += in
				H = strconv.Itoa(hx) + "_" + strconv.Itoa(hy)

				if (hx-lx == 1 || hx-lx == -1) && (ly-hy == 2 || ly-hy == -2) {
					lx = hx
					ly = hy + (-in)
				} else if ly-hy == 2 || ly-hy == -2 {
					ly += in
				}

				L = strconv.Itoa(lx) + "_" + strconv.Itoa(ly)

				if positionMap[L] == false {
					lVisit++
					positionMap[L] = true
				}

				i--
			}

		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(lVisit)
}
