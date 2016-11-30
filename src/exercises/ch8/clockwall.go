package main

import (
    "flag"
    "log"
    "strings"
    "net"
    "io"
    "os"
)

type WorldClock struct {
    dst  string
    addr string
    conn net.Conn
}

func mustCopy(dst io.Writer, src io.Reader) {
    if _, err := io.Copy(dst, src); err != nil {
        log.Fatal(err)
    }
}

func Dial(worldClock WorldClock) {
    conn, err := net.Dial("tcp", worldClock.addr)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()
    go mustCopy(os.Stdout, conn)
    mustCopy(conn, os.Stdin)
}

func main() {
    flag.Parse()

    // The command like "NewYork=localhost:8010"
    commands := flag.Args()
    if len(commands) == 0 {
        log.Fatal("Please assign a destination with port like 'NewYork=[IP]:[PORT]' ")
    }

    var worldClocks []WorldClock
    for _, command := range commands {
        com := strings.Split(command, "=")
        if len(com) != 2 {
            log.Fatal("Please check you command")
        }
        dst, addr := com[0], com[1]
        worldClock := WorldClock{dst:dst, addr:addr}
        worldClocks = append(worldClocks, worldClock)
    }
    for _, worldClock := range worldClocks {
        go Dial(worldClock)
    }

    select {

    }
}
