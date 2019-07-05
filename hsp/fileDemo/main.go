package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 自己编写一个函数，接收两个文件路径 srcFileName dstFileName

func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {

	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)

	// 打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY| os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	// 通过dstFile，获取到 Writer
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	return io.Copy(writer, reader)
}

func main()  {
	// 打开文件
	// 概念说明： file 的叫法
	// 1 file 叫 file对象
	// 2 file 叫 file指针
	// 3 file 叫 file 文件的句柄
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	// 当函数退出时，要及时关闭file
	defer file.Close()

	// 通过bufio读取文件内容
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf(str)
	}

	// 直接读取文件
	content, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		fmt.Printf("read file err=%v", err)
	}
	fmt.Printf("%v", string(content))

	fmt.Println("文件读取结束。。。")


	file3, err := os.OpenFile("./abc.txt", os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}

	str := "hello,Gardon\r\n"
	writer := bufio.NewWriter(file3)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)

	}
	//因为writer是带缓存，因此在调用WriterString方法时，其实
	//内容是先写入到缓存的,所以需要调用Flush方法，将缓冲的数据
	//真正写入到文件中， 否则文件中会没有数据!!!
	writer.Flush()
	file3.Close()
	fmt.Println("文件读写。。。")

	// 文件copy
	srcFile := "./ASCII.png"
	dstFile := "./ascii1.png"

	_, err = CopyFile(dstFile, srcFile)
	if err == nil {
		fmt.Printf("copy done\n")
	} else {
		fmt.Printf("copy err=%v\n", err)
	}

}
