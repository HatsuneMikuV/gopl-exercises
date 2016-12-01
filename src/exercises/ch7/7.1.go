package main

import (
    "os"
    "bufio"
    "log"
    "io"
    "fmt"
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
            if err == nil {
                log.Println("Finish count words.")
            } else {
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
    fmt.Printf("The reader has %d words\n", wordCounter(file))
}