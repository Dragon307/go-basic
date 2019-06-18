package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func ReadFile(filePath string) []byte {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Panicln("Read error")
	}
	return content
}

// 使用buffer读取大文件

func ReadFileUseBuffer(filePath string, handle func([]byte)) error {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)

	for {
		buf, err := buf.ReadBytes('\n')

		if err != nil {
			if err == io.EOF{
				return nil
			}
			return err
		}
		handle(buf)

		return nil
	}
}

// 分片处理

func ReadFileUseSlice(filePath string, handle func([]byte)) error {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("can't opened this file")
		return err
	}
	defer f.Close()
	s := make([]byte, 4096)

	for {
		switch nr, err := f.Read(s[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading: %s\n", err.Error())
			os.Exit(1)
		case nr == 0:
			return nil
		case nr > 0:
			handle(s[0:nr])
		}
	}
	return nil
}

func printFile(b []byte)  {
	fmt.Println(b)
}

func main()  {
	// ReadFileUseSlice("./videos/ddd.mp4", printFile)
	ReadFileUseBuffer("./videos/ddd.mp4", printFile)
}