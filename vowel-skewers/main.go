package main

import (
	"fmt"
	"strings"

)

var (
	vowels = []string{"A","E","I","O","U"}
	consonant = []string{"Q","W","R","T","Y","P","L","K","J","H","G","F","D","S","M","N","B","V","C","X","Z"}
)

func validateAlt(str string, skewer []string) bool {
	if strings.Contains(strings.Join(vowels[:], ""), str ) && 
	   strings.Contains(strings.Join(consonant[:], ""), skewer[len(skewer) - 1]) {
		return true
	}else if strings.Contains(strings.Join(consonant[:], ""), str ) && 
			 strings.Contains(strings.Join(vowels[:], ""), skewer[len(skewer) - 1]) {
		return true
	}
	return false
}

func validateSpace(sk []string, spaces int) bool {
	count := 0 
	for _, sp := range sk {
		if sp == "-" {
			count++
		}else if sp != "-" && sp != "" {
			if count != spaces {
				return false
			}
			count = 0
		}
	}
	return true
}

func isAuthenticSkewer(skewer string) bool {
	arrSkewer := strings.Split(skewer, "-")
	countSkewer := []string{}
	spaces := 0

	if !strings.Contains(strings.Join(consonant[:], ""), arrSkewer[0]) || 
		!strings.Contains(strings.Join(consonant[:], ""), arrSkewer[len(arrSkewer) - 1]) {
		return false
	}	
	
	for _, str := range strings.Split(skewer, "")[1:] {
		if str == "-" {
			spaces++
		}else if str != "-" {
			break
		}
	}
	
	if !validateSpace(strings.Split(skewer, "")[1:], spaces) {
		return false
	}

	countSkewer = append(countSkewer, arrSkewer[0])
	for _, str := range arrSkewer[1: len(arrSkewer) - 2] {
		if !validateAlt(str,countSkewer) {
			return false
		}
		countSkewer = append(countSkewer, str)
	}
	return true
}

func main() {
	fmt.Printf("Skewer %s result:> %t\n","B--A--N--A--N--A--S",isAuthenticSkewer("B--A--N--A--N--A--S"))
}