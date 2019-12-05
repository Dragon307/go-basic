package main

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/scrypt"
)

func main()  {
	dk, err := scrypt.Key([]byte("123456"), []byte("qwer"), 1 << 15, 8, 1,32)
	if err != nil {
		return
	}
	fmt.Println(base64.StdEncoding.EncodeToString(dk))
}
