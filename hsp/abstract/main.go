package main

import "fmt"

type Account struct {
	AccountNo string
	Pwd string
	Balance float64
}


// 方法
// 1 存款
func (account *Account) Deposite(money float64, pwd string)  {

	if pwd != account.Pwd {
		fmt.Println("Incorrect password ")
		return
	}

	if money <= 0 {
		fmt.Println("Incorrect money")
		return
	}
	account.Balance += money
	fmt.Println("Deposite Success")
}

// 2 取款
func (account *Account) WithDraw(money float64, pwd string)  {
	if pwd != account.Pwd {
		fmt.Println("Incorrect password ")
		return
	}

	if money <= 0 || money > account.Balance {
		fmt.Println("Incorrect money")
		return
	}
	account.Balance -= money
	fmt.Println("WithDraw Success")
}

// 3 查询余额
func (account *Account) Query(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("Incorrect password ")
		return
	}
	fmt.Printf("Your account is = %v balance = %v\n", account.AccountNo, account.Balance)
}

func main() {

	account := Account{
		AccountNo : "gs1111111",
		Pwd : "666666",
		Balance : 100.0,
	}

	//这里可以做的更加灵活，就是让用户通过控制台来输入命令...
	//菜单....
	account.Query("666666")
	account.Deposite(200.0, "666666")
	account.Query("666666")
	account.WithDraw(150.0, "666666")
	account.Query("666666")

}