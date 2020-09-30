package accounts

import (
	"errors"
	"fmt"
)

// Account Struct
type Account struct {
	// owner, balance도 대문자로 해야 export 되어 public처럼 쓸 수 있다.
	owner   string
	balance int
	deposit int
}

//컨벤션 : 에러 변수의 이름은 err000 이어야 한다.
var errNoMoney = errors.New("Cant't withdraw")

// Go에는 constructor가 없기에 따로 만드는 패턴이 있다.
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}

	// 새로운 struct를 만들기 싫으니 레퍼런스를 리턴한다.
	return &account
}

// 컨벤션 : struct의 첫글자를 소문자로 사용하여 a로 쓴다.
// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a *Account) Balance() int {
	return a.balance
}

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

// 오버라이팅 할 수 있다. String()은 클래스명 자체를 프린트했을 때 실행됨
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
