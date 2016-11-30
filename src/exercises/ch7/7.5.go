package main

import (
	"fmt"
    "io"
    "io/ioutil"
//    "strings"
    "os"
//    "log"

)

/*
    练习 7.5： io包里面的LimitReader函数接收一个io.Reader接口类型的r和字节数n，
    并且返回另一个从r中读取字节,但是当读完n个字节后就表示读到文件结束的Reader。
    实现这个LimitReader函数.
 */
type MyLimitReader struct {
    R io.Reader
    N int64
}

func (myLimitReader *MyLimitReader) Read (p []byte) (n int, err error) {
    if myLimitReader.N <= 0 {
        return 0, io.EOF
    }
    if int64(len(p)) > myLimitReader.N {
        p = p[0:myLimitReader.N]
    }
    n, err = myLimitReader.R.Read(p)
    myLimitReader.N -= int64(n)
    return
}

func LimitReader(r io.Reader, n int64) io.Reader {
    return &MyLimitReader{r, n}
}

func main() {
    file, _ := os.Open("./testfile.txt")
    r := LimitReader(file, 10)
    filecontent, _ := ioutil.ReadFile("./testfile.txt")
    fmt.Println(r.Read(filecontent))
}
