package main

import (
	"fmt"
	"log"
)

var printFn = func(idx int, bean interface{}) error {
	fmt.Printf("%d: %#v\n", idx, bean.(*Account))
	return nil
}

func main()  {
	fmt.Println("Welcome bank of xorm!")

	count, err := getAccountCount()
	if err != nil {
		log.Fatalf("Fail to get account count: %v", err)
	}
	fmt.Println("Account count:", count)

	// 自动创建至 10 个账户
	for i := count; i< 10; i++ {
		if err = NewAccount(fmt.Sprintf("joe%d",i), float64(i)*100); err != nil {
			log.Fatalf("Fail to create account: %v", err)
		}
	}

	// 迭代查询
	fmt.Println("Query all columns:")
	x.Iterate(new(Account), printFn)

	// 更加灵活的迭代
	a := new(Account)
	rows, err := x.Rows(a)
	if err != nil {
		log.Fatalf("Fail to rows: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(a); err != nil {
			log.Fatalf("Fail get row: %v", err)
		}
		fmt.Printf("%#v\n",a)
	}

	// 查询特定字段
	fmt.Println("\nOnly query name:")
	x.Cols("name").Iterate(new(Account), printFn)

	// 排除特定字段
	fmt.Println("\nQuery all but name:")
	x.Omit("name").Iterate(new(Account), printFn)

	// 查询结果偏移
	fmt.Println("\nOffest 2 and limit 3:")
	x.Limit(3,2).Iterate(new(Account), printFn)

	// 测试 LRU 缓存
	getAccount(1)
	getAccount(1)
}