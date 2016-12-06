package main

import (
    "fmt"
    //"time"
    "time"
)

// 临界资源的大小
const MaxSize int = 10

// 全局变量,用来标记产品的id
var id int

// 结构体Item待放入资源池的产品
type Item struct {
    Id int
}

// 生产者函数
func producer(resources chan <- Item) {
    for {
        item := Item{Id:id}
        id += 1
        resources <- item
        fmt.Printf("+ 添加商品[%d]进入资源池\n", item.Id)
    }
}

// 消费者函数
func consumer(resources <-chan Item) {
    for {
        item := <-resources
        fmt.Printf("- 从资源池消费商品[%d]\n", item.Id)
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    // 用一个channel类型的数据结构表示临界资源,大小为MaxSize
    resources := make(chan Item, MaxSize)

    // 用goroutine的方式并发运行生产者函数和消费者函数
    go producer(resources)
    go consumer(resources)

    // 阻塞在此,防止main函数退出时生产者消费者函数同时退出.
    select {}
}
