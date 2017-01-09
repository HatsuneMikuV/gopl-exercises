package main

var deposits = make(chan int)
var balances = make(chan int)
var withdraws = make(chan int)

// Deposit function
func Deposit(amount int) {
	deposits <- amount
}

// Balance function
func Balance() int {
	return <-balances
}

// Withdraw function
func Withdraw(amount int) bool {
	balance := Balance()
	if amount > balance {
		return false
	}
	balance -= amount
	withdraws <- balance
	return true
}

// teller function
func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case withdraw := <-withdraws:
			balance -= withdraw
		}
	}
}

// init function
func init() {
	go teller()
}
