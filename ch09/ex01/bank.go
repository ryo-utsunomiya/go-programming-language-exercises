// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

type WithdrawResult struct {
	amount    int
	isSuccess chan bool
}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraws = make(chan WithdrawResult)

func Deposit(amount int) { deposits <- amount }
func Withdraw(amount int) bool {
	isSuccess := make(chan bool)
	withdraws <- WithdrawResult{amount, isSuccess}
	return <-isSuccess
}
func Balance() int { return <-balances }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case w := <-withdraws:
			if balance >= w.amount {
				balance -= w.amount
				w.isSuccess <- true
			} else {
				w.isSuccess <- false
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
