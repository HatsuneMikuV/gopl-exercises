package main

import (
    "flag"
    "log"
    "net"
    "time"
    "fmt"
    "io"
)

// Presume port 8010 return NewYork time
// 8020 return London time
// 8030 return Tokyo time

var locTime map[string]*time.Location

const FORMAT = "Jan 2, 2006 at 3:04:05pm"

var port = flag.String("port", "8000", "Assign a port")

func handleConnn(c net.Conn) {
    defer c.Close()
    for {
        t, _ := time.ParseInLocation(FORMAT, time.Now().Format(FORMAT), locTime[*port])
        _, err := io.WriteString(c, "\r" + locTime[*port].String() + " :" + t.String())
        if err != nil {
            fmt.Println("Disconnected.")
            return // e.g., client disconnected
        }
        time.Sleep(2 * time.Second)
    }
}

func init() {
    locTime = make(map[string]*time.Location)
    loc, err := time.LoadLocation("America/New_York")
    if err != nil {
        log.Fatal(err)
    }
    locTime["8010"] = loc

    loc, err = time.LoadLocation("Europe/London")
    if err != nil {
        log.Fatal(err)
    }
    locTime["8020"] = loc

    loc, err = time.LoadLocation("Asia/Tokyo")
    if err != nil {
        log.Fatal(err)
    }
    locTime["8030"] = loc

    for k, v := range locTime {
        println(k, v.String())
    }
}

func main() {
    flag.Parse()

    listern, err := net.Listen("tcp", "localhost:" + *port)
    fmt.Println("begin listining on port " + *port)
    if err != nil {
        log.Fatal(err)
    }

    for {
        conn, err := listern.Accept()
        fmt.Println("Accept a connection.")
        if err != nil {
            log.Print(err)
            continue
        }
        go handleConnn(conn)
    }
}
