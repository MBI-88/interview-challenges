package main

import (
	"fmt"
	"strconv"
	"strings"
)

func minFrame(matiz [][]string) string {
	var pos int
	for _, frame := range matiz {
		maximun := len(frame)
		for j, next := range matiz {
			if len(next) > maximun {
				maximun = len(next)
				pos = j
			}
		}
	}
	return strings.Join(matiz[pos], "")

}

func longestAltSubtring(str string) string {
	var (
		frame      []string
		totalTrack [][]string
	)
	for i := 0; i < len(str); i++ {
		if len(frame) == 0 {
			frame = append(frame, string(str[i]))
		} else {
			last, _ := strconv.Atoi(frame[len(frame)-1])
			step, _ := strconv.Atoi(string(str[i]))
			if last%2 != step%2 {
				frame = append(frame, string(str[i]))
			} else {
				totalTrack = append(totalTrack, frame)
				frame = []string{string(str[i])}
			}
		}

	}
	return minFrame(totalTrack)
}

func main() {
	fmt.Println(longestAltSubtring("225424272163254474441338664823")) // -> "272163254"
	fmt.Println(longestAltSubtring("594127169973391692147228678476")) // -> "16921472" 
	fmt.Println(longestAltSubtring("721449827599186159274227324466")) // ->  "7214"
}
