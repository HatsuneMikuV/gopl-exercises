// 执行流程:
// 运行后进行第一次锐捷连接,将结果写进log
// 第一次运行后休眠

package main

import (
	"bufio"
	"flag"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

var sh = flag.String("path", ".", "shell file location")
var account = flag.String("u", "", "RuiJie account")
var password = flag.String("p", "", "RuiJie password")
var regularArg = []string{"-d", "1", "-u", "-p"}
var logFile *os.File
var stopTime = 30 * time.Second
var repeatTime = 2 * time.Minute
var infoLogger *log.Logger
var requrl = "https://www.baidu.com" // 嗯,发挥百度的光荣作用

// /home/lwh/rjsupplicant/./rjsupplicant.sh -d 1 -u 201424133254 -p wh5622

func init() {
	logfile := "log.log"
	var err error
	logFile, err = os.Create(logfile)
	if err != nil {
		log.Fatal(err)
	}
	// year, month, day := time.Now().Date()
	//date := string(year) + "-" + string(month) + "-" + string(day)
	infoLogger = log.New(logFile, "[info] ", log.Ltime)
	infoLogger.Println("launch program...")
}

func checkConnect() int {
	client := &http.Client{}
	request, err := http.NewRequest("GET", requrl, nil)
	if err != nil {
		infoLogger.Fatal(err)
	}
	response, err := client.Do(request)
	if err != nil {
		infoLogger.Fatal(err)
	}
	return response.StatusCode
}

func runDial() {
	flag.Parse()

	cmd := exec.Command(*sh, regularArg[0], regularArg[1], regularArg[2], *account, regularArg[3], *password)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		infoLogger.Fatal(err)
	}
	defer stdin.Close()

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	go func() {
		if err := cmd.Start(); err != nil {
			infoLogger.Fatal(err)
		}
		infoLogger.Printf("start dial, and pid is %d\n", cmd.Process.Pid)
		r := bufio.NewReader(os.Stdout)
		d, _, _ := r.ReadLine()
		infoLogger.Println(d)
	}()

	tick := time.Tick(1 * time.Second)
	for countdown := 3; countdown > 0; countdown-- {
		<-tick
	}

	infoLogger.Println("stop dial")
}

func main() {
	runDial()
	//statusCode := checkConnect()
	//if statusCode != 200 {
	//    //
	//}
}
