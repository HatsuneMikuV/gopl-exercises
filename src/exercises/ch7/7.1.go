package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func wordCounter(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	var count = 0
	for {
		token := scanner.Scan()
		if token {
			count += 1
		} else {
			err := scanner.Err()
			if err != nil {
				log.Fatal(err)
			}
			return count
		}
	}
}

func lineCounter(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	var count = 0
	for {
		token := scanner.Scan()
		if token {
			count += 1
		} else {
			err := scanner.Err()
			if err != nil {
				log.Fatal(err)
			}
			return count
		}
	}
}

func main() {
	file, err := os.Open("test")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Printf("The reader has %d word(s)\n", wordCounter(file))
	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The reader has %d line(s)\n", lineCounter(file))

}
