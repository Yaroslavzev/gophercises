package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scanf("%s", &input)

	wordsCounter := 1
	for _, v := range input {
		//fmt.Println(reflect.TypeOf(v))

		if strings.ToUpper(string(v)) == string(v) {
			wordsCounter++
		}
	}

	fmt.Println(wordsCounter)
}
