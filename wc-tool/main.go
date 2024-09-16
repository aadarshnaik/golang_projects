package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) > 3 {
		fmt.Println(`
Invalid Format
Usgae [command] -[option] [filename]
		`)
	}
	var input *os.File
	var err error
	if len(args) == 3 {
		input, err = os.Open(args[2])
		if err != nil {
			fmt.Println("error opening file:", err)
		}
		defer input.Close()
	} else {
		input = os.Stdin
	}

	option := args[1]
	switch option {
	case "-c":
		byt, err := os.ReadFile(args[2])
		if err != nil {
			fmt.Println("Unable to read file")
		}
		fmt.Println(len(byt))
	case "-l":
		in := bufio.NewScanner(input)

		count := 0
		for in.Scan() {
			count++
		}
		fmt.Println(count)
	case "-w":
		byt, err := os.ReadFile(args[2])
		if err != nil {
			fmt.Println("Unable to read file")
		}
		str := string(byt)
		words := strings.Fields(str)
		fmt.Println(len(words))
	case "-m":
		byt, err := os.ReadFile(args[2])
		if err != nil {
			fmt.Println("Unable to read file")
		}
		ch := []rune(string(byt))
		fmt.Println(len(ch))
	default:
		byt, err := os.ReadFile(args[1])
		if err != nil {
			fmt.Println("Unable to read file")
		}
		str := string(byt)
		words := strings.Fields(str)

		file, err := os.Open(args[1])
		if err != nil {
			fmt.Println("error opening file:", err)
		}
		defer file.Close()

		in := bufio.NewScanner(file)

		count := 0
		for in.Scan() {
			count++
		}

		fmt.Println(count, " ", len(words), " ", len(byt), " ", os.Args[1])
	}
}
