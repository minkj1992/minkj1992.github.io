package main

import (
	"log"

	"github.com/minkj1992/go_nomad/accounts/accounts"
)

func main() {
	account := accounts.NewAccount("leoo")
	account.Deposit(10)
	err := account.Withdraw(10)
	if err != nil {
		log.Fatalln(err) // print & sig kill
	}
	account.ChangeOwner("mom")
}
