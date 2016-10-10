package main

import (
	"fmt"
    "io"
    "os"
)

/*
    练习 7.2： 写一个带有如下函数签名的函数CountingWriter， 传入一个io.Writer接口类型， 返
    回一个新的Writer类型把原来的Writer封装在里面和一个表示写入新的Writer字节数的int64类
    型指针
 */
type NewIOWriter struct {
    iw io.Writer
    bytenums int
}

func (iw *NewIOWriter) Write (b []byte) (int, error) {
    iw.bytenums = len(b)
    return iw.bytenums, nil
}

func CountingWriter (iw io.Writer) (iwr NewIOWriter, bytenums *int64) {
    newbytes := []byte("This is new bytes after invoke")
    
    iwr.iw = iw
    iwr.Write(newbytes)

    return iwr, bytenums
}

func main() {
    file, err:= os.Open("./testfile.txt")
    if err == nil {
        iw, bytenums := CountingWriter(file)
        fmt.Println(iw.bytenums, bytenums)
    }
    defer file.Close()
}
