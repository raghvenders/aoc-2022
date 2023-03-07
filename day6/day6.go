package day6

import (
	"bufio"
	"log"
	"os"
)

func Devices(file string) {

	//input := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"

	// input := "bvwbjplbgvbhsrlpgdmjqwftvncz"

	//input := "nppdvjthqldpwncqszvftbrmjlhg"

	//input := "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"

	//input := "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"

	//m - 109
	//j - 106
	//q - 113
	input, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	dev := make(map[byte]interface{})
	count := 0
	var i int = -1

	// reading crates state (empty line is a delimeter)
	//inputState := make([]string, 0)
	for scanner.Scan() {
    input := scanner.Text()
		for i = 0; i < len(input); i++ {
			//log.Println(string(input[i]), count, i)
			if dev[input[i]] == nil {
				dev[input[i]] = i
				count += 1
				log.Println(string(input[i]), count, i)
			} else if count <= 4 {
				count = 0
				v, _ := dev[input[i]].(int)
				i = v + 1

				for k := range dev {
					delete(dev, k)
				}
				log.Println("Latest Dev", dev)
				continue
			}
			if count == 4 {
				break
			}

		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
    
    if i==0 {
     i = -1
  }
		log.Println("Marker Character :", i)

	//return i +1

}
