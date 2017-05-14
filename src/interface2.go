package main

import (
	"fmt"
	"bufio"
)

/*
io下write接口的定义
type Writer interface {
	Write(p[]byte) (int, error)
}

os.Stdout变量的定义为一个*os.File类型
Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")

*byte.Buffer类型中实现了writer\reader接口
*/
type ByteCount int

//实现接口io.Writer
func (c *ByteCount) Write(p []byte) (int, error) {
	*c += ByteCount(len(p))
	return len(p), nil
}

func main() {
	var c ByteCount = 0
	var name = "neojos"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)

	var line []byte = []byte("good night neojos!")
	for start, width := 0, 0; start < len(line); start += width {
		advance, token, error := bufio.ScanWords(line[width:], true)
		fmt.Printf("%d, %s, %v\n", advance, string(token), error)
		width += advance
	}
}
