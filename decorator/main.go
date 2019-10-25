package main
// https://coolshell.cn/articles/17929.html
import (
	"fmt"
	"reflect"
)

func decorator(f func(s string)) func(s string) {
	return func(s string) {
		fmt.Println("Started")
		f(s)
		fmt.Println("Done")
	}
}

func Hello(s string) {
	fmt.Println(s)
}

// 泛型的装饰器
func Decorator(decoPtr, fn interface{}) (err error) {
	var decotatedFunc, targetFunc reflect.Value

	decotatedFunc = reflect.ValueOf(decoPtr).Elem()
	targetFunc = reflect.ValueOf(fn)

	v := reflect.MakeFunc(targetFunc.Type(), func(in []reflect.Value) (out []reflect.Value) {
		fmt.Println("before")
		out = targetFunc.Call(in)
		fmt.Println("after")
		return
	})
	decotatedFunc.Set(v)
	return
}

func foo(a, b, c int) int {
	fmt.Printf("%d, %d, %d \n", a, b,c)
	return a+b+c
}

func bar(a, b string) string {
	fmt.Printf("%s, %s \n", a, b)
	return  a+b
}

type MyFoo func(int, int, int) int

func main() {
	var myfoo MyFoo
	Decorator(&myfoo, foo)
	myfoo(1,2,3)
	// decorator(Hello)("Hello, World!")
}