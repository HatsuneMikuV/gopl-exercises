package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
    "time"
)

type Process struct {
	name      string
	resourceA Resource
	resourceB Resource
	resourceC Resource
}

type Resource struct {
	name       string
	max        float64
	allocation float64
	need       float64
}

var available map[string]float64
var orders []Process
var finish chan struct{}

type simpleQueue struct {
	data []Process
	max  int
}

func (q *simpleQueue) length() int {
	return len(q.data)
}

func (q *simpleQueue) full() bool {
	return q.length() >= q.max
}

func (q *simpleQueue) empty() bool {
	return len(q.data) == 0
}

func (q *simpleQueue) pop() Process {
	if q.length() <= 0 {
		panic("队列为空")
	}
	ret := q.data[0]
	q.data = q.data[1:]
	return ret

}

func (q *simpleQueue) top() Process {
    return q.data[0]
}

func (q *simpleQueue) push(process Process) {
	if q.full() {
		fmt.Println("队列已满,后续添加无效")
		return
	}
	q.data = append(q.data, process)
}

func (q *simpleQueue) traverse() {
	fmt.Printf("size = %d, max = %d\n", q.length(), q.max)
	for _, data := range q.data {
		fmt.Println(data.name)
	}
}

var queue simpleQueue

func getDataset(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return data
}

func initProcesses(datasetJson interface{}) {
	queue.data = make([]Process, 0, 1024)
	m := datasetJson.(map[string]interface{})
	ps := m["Process"]
	processesd := ps.([]interface{})
	finish = make(chan struct{})
	queue.max = len(processesd)
	for _, process := range processesd {
		p := process.(map[string]interface{})
		name := p["name"].(string)
		resourceAJson := p["resourceA"].(map[string]interface{})
		resourceBJson := p["resourceB"].(map[string]interface{})
		resourceCJson := p["resourceC"].(map[string]interface{})
		resourceA := Resource{name: "A", max: resourceAJson["Max"].(float64),
			allocation: resourceAJson["Allocation"].(float64), need: resourceAJson["Need"].(float64)}
		resourceB := Resource{name: "B", max: resourceBJson["Max"].(float64),
			allocation: resourceBJson["Allocation"].(float64), need: resourceBJson["Need"].(float64)}
		resourceC := Resource{name: "C", max: resourceCJson["Max"].(float64),
			allocation: resourceCJson["Allocation"].(float64), need: resourceCJson["Need"].(float64)}
		ps := Process{name: name, resourceA: resourceA, resourceB: resourceB, resourceC: resourceC}
		queue.push(ps)
	}

	rs := m["Available"].(map[string]interface{})
	available = make(map[string]float64)
	available["A"] = rs["A"].(float64)
	available["B"] = rs["B"].(float64)
	available["C"] = rs["C"].(float64)

}

// 分配资源
func allocation(process *Process) {
	process.resourceA.allocation -= process.resourceA.need
	available["A"] += process.resourceA.need
	process.resourceA.need = 0

	process.resourceB.allocation -= process.resourceB.need
	available["B"] += process.resourceB.need
	process.resourceB.need = 0

	process.resourceC.allocation -= process.resourceC.need
	available["C"] += process.resourceC.need
	process.resourceC.need = 0
}

// 安全性算法
func checkSecurity() bool {
	dupAvailable := make(map[string]float64)
	finish := make(map[string]bool)
	for key, value := range available {
		dupAvailable[key] = value
	}
	count := 0
	for {
		for _, process := range queue.data {
            printProcessStatus(queue.data, dupAvailable)
			if _, ok := finish[process.name]; !ok && process.resourceA.need <= dupAvailable["A"] &&
				process.resourceB.need <= dupAvailable["B"] &&
				process.resourceC.need <= dupAvailable["C"] {
                fmt.Printf("安全性检查:%s\n", process.name)
				dupAvailable["A"] += process.resourceA.allocation + process.resourceA.need
				dupAvailable["B"] += process.resourceB.allocation + process.resourceB.need
				dupAvailable["C"] += process.resourceC.allocation + process.resourceC.need
				finish[process.name] = true
			}
			if len(finish) == queue.length() {
				return true
			}
		}
		count++
		if count > queue.length() {
			return false
		}

	}
}

// 选择一条满足条件的进程
func selectProcess() {
	for !queue.empty() {
		//queue.traverse()
		process := queue.top()
        queue.traverse()
        //printProcessStatus(queue.data, available)
		if process.resourceA.need <= available["A"] && process.resourceB.need <= available["B"] &&
			process.resourceC.need <= available["C"] {
            fmt.Printf("进程%s满足条件\n", process.name)
			if checkSecurity() {
				allocation(&process)
                queue.pop()
				fmt.Println("资源已分配给进程" + process.name)
				orders = append(orders, process)
			} else {
				fmt.Printf("进程%s无法通过安全性检查!\n", process.name)
				queue.push(process)
				fmt.Printf("size:%d, max:%d\n", queue.length(), queue.max)
				if queue.full() {
					fmt.Println("存在死锁状态,无法分配资源!")
				}
			}
		} else {
            queue.push(queue.pop())
        }
        time.Sleep(2*time.Second)
	}
	//for index, ps := range queue.data {
	//    if ps.resourceA.need <= available["A"] && ps.resourceB.need <= available["B"] &&
	//        ps.resourceC.need <= available["C"] {
	//        // 满足条件,进行安全性检查
	//        if checkSecurity(ps.name) {
	//            allocation(&processes[index])
	//            fmt.Println("资源已分配给进程" + ps.name)
	//            orders = append(orders, ps)
	//        } else {
	//            fmt.Printf("进程%s无法通过安全性检查!\n", ps.name)
	//        }
	//    } else {
	//        fmt.Println("进程" + ps.name + "不满足分配要求")
	//        processes = append(processes, ps)
	//        processes = processes[index:]
	//    }
	//}

	//for ps := range processesch {
	//    if ps.resourceA.need <= available["A"] && ps.resourceB.need <= available["B"] &&
	//        ps.resourceC.need <= available["C"] {
	//        // 满足条件,进行安全性检查
	//        if checkSecurity(ps.name) {
	//            allocation(&ps)
	//        } else {
	//            fmt.Printf("进程%s无法通过安全性检查!\n", ps.name)
	//        }
	//    } else {
	//        continue
	//    }
	//}
}

func printProcessStatus(processes []Process, avail map[string]float64) {
	fmt.Println("name\tMax   | Allocation | Need  | Available|")
	fmt.Println("    \tA B C | A   B    C | A B C | A   B   C|")
	for _, ps := range processes {
		fmt.Println(ps.name, "    ", ps.resourceA.max, ps.resourceB.max, ps.resourceC.max, "|",
			ps.resourceA.allocation, ps.resourceB.allocation, ps.resourceC.allocation, "     |",
			ps.resourceA.need, ps.resourceB.need, ps.resourceC.need, "|",
			avail["A"], " ", avail["B"], "", avail["C"], "|")
	}
}

func main() {
	datasetFile := getDataset("./data1.json")
	var datasetJson interface{}
	json.Unmarshal(datasetFile, &datasetJson)
	initProcesses(datasetJson)
	orders = make([]Process, 0)
	//printProcessStatus(queue.data, available)
	selectProcess()
	for _, order := range orders {
		fmt.Println("#", order)
	}
}
