package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 定义一个结构体，用于保存统计结果
type CharCount struct {
	ChCount int // 记录英文个数
	NumCount int // 记录数字的个数
	SpaceCount int // 记录空格的个数
	OtherCount int // 记录其它字符的个数
}

func main()  {
	// 打开一个文件，创建一个reader
	// 每读取一行，就统计一次 然后保存到结构体

	fileName := "./abc.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()
	// 定义CharCount 实例
	var count CharCount
	// 创建一个 Reader
	reader := bufio.NewReader(file)

	// 开始循环的读取 fileName 的内容
	for {
		str, err := reader.ReadString('\n')
		fmt.Println(str)
		if err == io.EOF { // 读取到文件末尾就退出
			break
		}
		// 遍历 str， 进行统计
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough // 穿透
			case v>= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++

			}
		}
	}
	//输出统计的结果看看是否正确
	fmt.Printf("字符的个数为=%v 数字的个数为=%v 空格的个数为=%v 其它字符个数=%v",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)

}
