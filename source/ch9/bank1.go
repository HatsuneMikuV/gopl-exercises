package main

var deposits = make(chan int)
var balances = make(chan int)

// Deposit function
func Deposit(amount int) {
	deposits <- amount
}

// Balance function
func Balance() int {
	return <-balances
}

// teller function
func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

// init function
func init() {
	go teller()
}
