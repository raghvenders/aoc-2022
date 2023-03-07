package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"

	//_ "regexp"
	"regexp"
	"strconv"
	"strings"
	_ "strings"
)

var (
	monkeys = make(map[string]*monkey, 0)
)

const (
	REGEX_MONKEY = `^Monkey (\d+):
  Start items: (\d+(, \d+)*)
  Operation: new = old (\*|\+|\/|\-) (\d+)
  Test: divisible by (\d+)
    If true: throw to monkey (\d+)
    If false: throw to monkey (\d+)`
)

type monkey struct {
	name      string
	items     []int
	operation [2]string
	test      [2]string
	truth     string
	falsy     string
}

func roundMoney(mky *monkey) {

	for _, item := range mky.items {
		switch mky.operation[0] {
		case "*":
			r, err := strconv.Atoi(mky.operation[1])
			if err != nil {
				r = item
			}
			item = item * r
		case "+":
			r, err := strconv.Atoi(mky.operation[1])
			if err != nil {
				r = item
			}
			item = item + r
		case "-":
			r, err := strconv.Atoi(mky.operation[1])
			if err != nil {
				r = item
			}
			item = item - r
		case "/":
			r, err := strconv.Atoi(mky.operation[1])
			if err != nil {
				r = item
			}
			item = item / r

		}

		wl := item / 3

		testInt, _ := strconv.Atoi(mky.test[1])
		//fmt.Println(item, wl, testInt)

		if wl%testInt == 0 {
			anotherMonkey := monkeys["Monkey"+mky.truth].items
			anotherMonkey = append(anotherMonkey, wl)
			monkeys["Monkey"+mky.truth].items = anotherMonkey

		} else {
			//log.Printf("%s", mky.falsy)
			anotherMonkey := monkeys["Monkey"+mky.falsy].items
			anotherMonkey = append(anotherMonkey, wl)
			monkeys["Monkey"+mky.falsy].items = anotherMonkey

		}

		if len(mky.items) > 0 {
			mky.items = mky.items[1:]
		}

		//fmt.Printf("Monkey 3 %#v -- %#v", *monkeys["Monkey3"], *monkeys["Monkey0"])
	}
}

func (m *monkey) string() {

	//fmt.Printf("%s -- %v -- %v -- %v -- %s %s", m.name, m.items, m.operation, m.test, m.truth, m.falsy)

}

func MonkeyInMiddle(file string) {
	onDouble := func(data []byte, atEOF bool) (advance int, token []byte, err error) {

		//log.Println(string(data))
		//log.Println(atEOF)
		for i := 0; i < len(data); i++ {
			if data[i] == '\n' && (i+1 < len(data) && data[i+1] == '\n') {
				return i + 2, data[:i], nil
			}
		}
		/*
			if !atEOF {
				return 0, nil, nil
			}
		*/
		// There is one final token to be delivered, which may be the empty string.
		// Returning bufio.ErrFinalToken here tells Scan there are no more tokens after this
		// but does not trigger an error to be returned from Scan itself.
		return 0, data, bufio.ErrFinalToken
	}

	strToIntArr := func(input string, sep string) []int {
		strArr := strings.Split(input, sep)
		intArr := make([]int, 0, len(strArr))
		for _, v := range strArr {
			i, _ := strconv.Atoi(v)
			intArr = append(intArr, i)
		}
		return intArr
	}

	//_ = strToIntArr
	input, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(onDouble)
	for scanner.Scan() {
		line := scanner.Text()

		//log.Printf("%s \n", line)
		//log.Printf("%s", REGEX_MONKEY)

		re := regexp.MustCompile(`^Monkey (\d+):\n  Starting items: (\d+(, \d+)*)\n  Operation: new = old (\*|\+|\/|\-) ((\d+)|old)\n  Test: divisible by (\d+)\n    If true: throw to monkey (\d+)\n    If false: throw to monkey (\d+)`)
		splits := re.FindStringSubmatch(line)

		//log.Printf("%d \n", len(splits))

		//log.Printf("%v", monkeys["Monkey"+splits[1]])

		if monkeys["Monkey"+splits[1]] == nil {
			m := &monkey{
				name:      "Monkey" + splits[1],
				items:     strToIntArr(splits[2], ", "),
				operation: [2]string{splits[4], splits[5]},
				test:      [2]string{"/", splits[7]},
				truth:     splits[8],
				falsy:     splits[9],
			}
			monkeys["Monkey"+splits[1]] = m
		}

		//if monkeys[line[1]]

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

		//fmt.Printf("%s %#v \n", k, *v)
		//roundMoney(v)

	for j := 0; j < 20; j++ {
		for i := 0; i < len(monkeys); i++ {
			//	fmt.Printf("%s %#v \n", k, *v)
			monkey := monkeys["Monkey"+strconv.Itoa(i)]
			roundMoney(monkey)
		}
	}

	for _, v := range monkeys {
		fmt.Printf("\n \\()/ %v", v.items)
	}

}
