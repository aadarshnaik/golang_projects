package main

import (
	"fmt"
	"os"
	"text/scanner"
)

func main() {

	file, err := os.Open("step1/valid.json")
	checkErr(err)
	defer file.Close()

	var s scanner.Scanner
	s.Init(file)
	for {
		tok := s.Scan()
		if tok == scanner.EOF {
			break
		}
		fmt.Printf("Token: %s, Value: %s\n", s.TokenText(), string(tok))
	}
}

func checkErr(err error) error {
	if err != nil {
		return err
	}
	return nil
}
