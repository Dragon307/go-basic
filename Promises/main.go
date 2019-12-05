package main

import (
	"fmt"
	"github.com/fanliao/go-promise"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	p := promise.NewPromise()
	p.OnSuccess(func(v interface{}) {

		fmt.Printf("%s success", v)
	}).OnFailure(func(v interface{}) {
		fmt.Println(v, "Fail")
	}).OnComplete(func(v interface{}) {
		fmt.Println("Complete")
	})

	go func() {
		url := "http://example.com"
		resp, err := http.Get(url)
		defer resp.Body.Close()
		if err != nil {
			p.Reject(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			p.Reject(err)
		}
		p.Resolve(body)
	}()

	time.Sleep(time.Second * 10)
	//r, err := p.Get()
	//fmt.Printf("%s, %s", r, err)
}
