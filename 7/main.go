package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	for {
		var sum float64
		var err error
		line, _, err := r.ReadLine()
		if err != nil {
			return
		}
		marks := strings.Fields(string(line))
		var flag bool
		for _, mark := range marks {
			score, err := getScore(mark)
			if err != nil {
				fmt.Println(err.Error())
				flag = true
				break
			}
			sum += score
		}
		if !flag {
			fmt.Printf("%.2f\n", sum/float64(len(marks)))
		}
	}

}

func getScore(mark string) (float64, error) {
	switch mark {
	case "A":
		return 4.0, nil
	case "B":
		return 3.0, nil
	case "C":
		return 2.0, nil
	case "D":
		return 1.0, nil
	case "F":
		return 0, nil
	default:
		return 0, errors.New("Unknown")
	}
}
