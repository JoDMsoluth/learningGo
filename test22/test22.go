package test22

import (
	"fmt"
	"math/rand"
	"time"
)

// Thread가 없을 때는 프로그램을 하나만 돌릴 수 있다.
// Thread가 나오면서 다른 여러 프로그램을 돌릴 수 있다.
// 하나에 OS로 여러 Thread에 Context Switching할 때 Switching 비용이 너무 많이 들면 문제가 생긴다.
// Go에서는 CPU개수 만큼 OS Thread를 만들고 필요할 경우 OS Thread를 잘게 잘라서 Go thread를 만든다.

// 프로그램 >= 프로세스 > 쓰레드
// 프로그램 : 하나 or 여러 실행파일 + 데이터
// 더블클릭 -> 프로그램 실행 -> exe 파일을 메모리에 적제(load) -> os는 메모리의 ip포인트부터 한줄씩 실행
// 프로세스 : 적제된 실행파일을 프로세스라 한다.

// 쓰레드
// 장점
// 1. 멀티테스킹
// 단점
// 1. 동기화 : 각 쓰레드가 하나의 메모리를 돌려쓰기 때문에 메모리 내용을 동기화 해야함
// 해결 ) Mutex
// 2. deadlock
// 해결 ) Channel : 1. Fixed Sized, 2. Safe Thread, 3. Queue

type Account struct {
	balance int
}

func (a *Account) Widthdraw(val int) {
	a.balance -= val
}

func (a *Account) Deposit(val int) {

	a.balance += val

}

func (a *Account) Balance() int {

	balance := a.balance

	return balance
}

var accounts []*Account

func Transfer(sender, receiver int, money int) {

	accounts[sender].Widthdraw(money)
	accounts[receiver].Deposit(money)

}

func GetTotalBalance() int {

	total := 0
	for i := 0; i < len(accounts); i++ {
		total += accounts[i].Balance()
	}

	return total
}

func RandomTranfer() {
	var sender, balance int
	for {
		sender = rand.Intn(len(accounts))
		balance = accounts[sender].Balance()
		if balance > 0 {
			break
		}
	}

	var receiver int
	for {
		receiver = rand.Intn(len(accounts))
		if sender != receiver {
			break
		}
	}

	money := rand.Intn(balance)
	Transfer(sender, receiver, money)
}

func GoTransfer() {
	for {
		RandomTranfer()
	}
}

func PrintTotalBalance() {
	fmt.Printf("Total: %d\n", GetTotalBalance())
}

func Test22() {
	for i := 0; i < 20; i++ {
		accounts = append(accounts, &Account{balance: 1000})
	}

	PrintTotalBalance()

	for i := 0; i < 10; i++ {
		go GoTransfer()
	}

	for {
		PrintTotalBalance()
		time.Sleep(100 * time.Millisecond)
	}
}
