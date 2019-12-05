package main

import "sync"

func Len(sm sync.Map) int {
	length := 0
	f := func(key, value interface{}) bool {
		length++;
		return true
	}
	one := length
	length = 0;
	sm.Range(f)
	if one != length {
		one = length
		length = 0
		sm.Range(f)
		if one < length {
			return length
		}
	}
	return one
}

func main() {
	m := make(map[int]int)
	go func() {
		for {
			_ = m[1]
		}
	}()

	go func() {
		for {
			m[2] = 2
		}
	}()
	select {}
}
