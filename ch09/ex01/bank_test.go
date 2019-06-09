// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank_test

import (
	"testing"

	bank "github.com/ryo-utsunomiya/go-programming-language-exercises/ch09/ex01"
)

//func TestBank(t *testing.T) {
//	done := make(chan struct{})
//
//	// Alice
//	go func() {
//		bank.Deposit(200)
//		fmt.Println("=", bank.Balance())
//		done <- struct{}{}
//	}()
//
//	// Bob
//	go func() {
//		bank.Deposit(100)
//		done <- struct{}{}
//	}()
//
//	// Wait for both transactions.
//	<-done
//	<-done
//
//	if got, want := bank.Balance(), 300; got != want {
//		t.Errorf("Balance = %d, want %d", got, want)
//	}
//}

func TestWithDraw(t *testing.T) {
	bank.Deposit(200)

	if bank.WithDraw(100) != true {
		t.Error("want: true, got: false")
	}
	if bank.Balance() != 100 {
		t.Error("want: 100, got: 100")
	}

	if bank.WithDraw(300) != false {
		t.Error("want: true, got: false")
	}
	if bank.Balance() != 100 {
		t.Error("want: 100, got: 100")
	}
}
