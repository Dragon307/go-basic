package decorator

import (
	"net/http"
	"os"
	"runtime"
	"testing"
)

func TestAuth(t *testing.T) {
	if IsCI() {
		return
	}

	http.HandleFunc("/", Auth(f))

	if err := http.ListenAndServe(":8088", nil); err != nil {
		t.Error(err)
	}
}

func TestReflectDecorator(t *testing.T) {
	type MyFoo func(int, int, int) int
	var myfoo MyFoo
	Decorator(&myfoo, foo)
	myfoo(1,2,3)

	mybar := bar
	Decorator(&mybar, bar)
	mybar("hello", "world!")
}


func IsCI() bool {
	name, _ := os.Hostname()
	if name != "pibigstar" && runtime.GOOS == "linux" {
		return true
	}
	return false
}