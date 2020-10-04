package test1

import (
	"fmt"

	"github.com/hyehyeong/learnGo/test1/accounts"
	"github.com/hyehyeong/learnGo/test1/mydict"
)

func Test1() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
	account.Deposit(10)
	fmt.Println(account.Balance())
	// 에러를 받아오지 않으면 아무 일도 일어나지 않는다.
	err := account.Withdraw(20)
	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())

	dictionary := mydict.Dictionary{"first": "First word"}
	fmt.Println(dictionary)

	definition, err := dictionary.Search("second")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	err1 := dictionary.Add("hello", "Greeting")
	if err1 != nil {
		fmt.Println(err)
	} else {
		hello, err2 := dictionary.Search("hello")
		if err2 != nil {
			fmt.Println(err)
		} else {
			fmt.Println("found", "hello", "definition:", hello)
		}
	}

	err3 := dictionary.Update("hello", "Updated Greeting")
	if err3 != nil {
		fmt.Println(err)
	} else {
		hello, err4 := dictionary.Search("hello")
		if err4 != nil {
			fmt.Println(err4)
		} else {
			fmt.Println("found", "hello", "definition:", hello)
		}

		// delete
		dictionary.Delete("hello")
		hello, err5 := dictionary.Search("hello")
		if err5 != nil {
			fmt.Println(err5)
		} else {
			fmt.Println("found", "hello", "definition:", hello)
		}
	}
}
