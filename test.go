package main

import "fmt"

type S struct {
    n string
}

var ch chan S
var finish chan struct{}
var p []S

func foo() {
    count := 0
    for s := range ch {
        fmt.Printf("%p, %p, %d\n", &s, &p[count], count)
        count++
        ch <- S{"wmn"}
        if count == 1000 {
            finish <- struct{}{}
            close(ch)
            return
        }
    }
}

func main() {
    //ch = make(chan S, 100)
    //finish = make(chan struct{})
    //p = make([]S, 10000000)
    //for index := 0; index < 10; index++ {
    //    s := S{"lwh"}
    //    p = append(p, s)
    //    ch <- s
    //}
    //go foo()
    //select {
    //case <-finish:
    //    fmt.Println("完成哈哈哈!")
    //}
    //s := []int{1,2,3}
    //a := make([]int,0,10)
    //for _, d := range s {
    //    a = append(a, d)
    //}
    //a[0] = -1
    //fmt.Println(s)
    //fmt.Println(a)
    s := []int{1,2,3,4}
    fmt.Println(len(s))
    s = s[1:]
    s = s[1:]
    s = s[1:]
    fmt.Println(len(s))
}
