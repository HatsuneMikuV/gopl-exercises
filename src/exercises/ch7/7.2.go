package main

import (
    "io"
    "log"
    "fmt"
)

type countWriter1 int64
type countWriter2 struct {
    writer io.Writer
    nbytes *int64
}

func (cw *countWriter1) Write(b []byte) (int, error) {
    log.Print("Run countWriter1 writer\n")
    *cw += countWriter1(len(b))
    return len(b), nil
}

func (cw *countWriter2) Write(b []byte) (int, error) {
    log.Print("Run countWriter2 writer\n")
    nbytes := int64(len(b))
    cw.nbytes = &nbytes
    return len(b), nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
    cw2 := &countWriter2{writer:w}
    cw2.Write([]byte("Hello, magical world"))
    return cw2, cw2.nbytes
}

func main() {
    var cw1 countWriter1
    cw1.Write([]byte("Hello world."))
    cw2, nbytes := CountingWriter(&cw1)
    fmt.Println(cw2, nbytes)
}
