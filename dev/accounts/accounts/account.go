package accounts

import (
	"errors"
	"fmt"
)

// Account is for user (private fields)
type Account struct {
	owner   string
	balance int
}

// errFoo is convention
var (
	errNoMoney = errors.New("Can't withdraw: you don't have that money")
)

// private method to check account info
func (a *Account) checkAccount() {
	fmt.Println(a)
}

// Balance is getter for balance of account
func (a Account) Balance() int {
	// receiver 복사해도 상관없다.
	return a.balance
}

// Owner is getter for owner of account
func (a Account) Owner() string {
	return a.owner
}

// NewAccount is factory to make Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	// 새로 생성한 account를 value로 copy 시키고 싶지 않기 때문에 &로 전달
	return &account
}

/*
Deposit amount money on your account
receiver의 conv는 struct의 앞글자 소문자
`Pointer receiver`: go 에서는 value로 전달하기 때문에 copy가 일어난다. 이를 방지하기 위해서 `*receiver`를 사용한다.
*/
func (a *Account) Deposit(ammount int) {
	defer a.checkAccount()
	a.balance += ammount
}

// Withdraw amount from a account
func (a *Account) Withdraw(amount int) error {
	defer a.checkAccount()
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil // error의 nil type
}

// ChangeOwner changes owner of account
func (a *Account) ChangeOwner(newOwner string) {
	defer a.checkAccount()
	a.owner = newOwner
}

// like __repr__
func (a Account) String() string {
	return fmt.Sprint("Owner: ", a.Owner(), "\nBalance: ", a.Balance(), "\n")
}
