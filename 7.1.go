package main

import (
	"fmt"
    "io/ioutil"
    //"bufio"
)

/*
    练习 7.1： 使用来自ByteCounter的思路， 实现一个针对对单词和行数的计数器。 你会发现
    bufio.ScanWords非常的有用。(本解答不使用bufio.ScanWords,从执行效率来说真的不好用)
 */

const (
   INIT_STATE, SPACE_COMMA_STATE, CHAR_STATE = "INIT_STATE", "SPACE_COMMA", "CHAR_STATE"
)
func CountWordsLines(b []byte) (int, int) {
    // 统计单词数,每个单词默认以空格,或逗号分隔
    s := string(b[:])
    r := []rune(s)
    words := 0
    lines := 9
    state := INIT_STATE
    for i:=0; i<len(r)-1; {
        switch state {
        case INIT_STATE:
            i++
            if r[i] == ' ' || r[i] == ','  || r[i] == '.' {
                state = SPACE_COMMA_STATE
            } else {
                state = CHAR_STATE
                words++
            }
            if r[i] == '\n' {
                lines++
            }
        case SPACE_COMMA_STATE:
            i++
            if r[i] == ' ' || r[i] == ','  || r[i] == '.' {
                state = SPACE_COMMA_STATE
            } else {
                words++
                state = CHAR_STATE
            }
            if r[i] == '\n' {
                lines++
            }
        case CHAR_STATE:
            i++
            if r[i] == ' ' || r[i] == ','  || r[i] == '.' {
                state = SPACE_COMMA_STATE
            } else {
                state = CHAR_STATE
            }
            if r[i] == '\n' {
                lines++
            }
        }
    }
    return words, lines
}
func main() {
    fcontent, _:= ioutil.ReadFile("./testfile.txt")
    words, lines:= CountWordsLines(fcontent)
    fmt.Println(words, lines)

}
