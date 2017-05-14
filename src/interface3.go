package main

import (
	"io"
	"strings"
	"bytes"
	"fmt"
)

/*
在声明方法的时候，类型声明为指针和值需要明确的地方：
1. 虽然方法声明为了指针类型，但是通过值类型参数仍然可以调用.但这其实是编译器帮我们做了
处理，自动获取的变量的指针地址
type IntSet struct {
}
func (i *IntSet) String() string {
}

var is IntSet
is.String()

2. 还需要注意：变量和临时变量是有区别的，比如下面这样是会编译错误的。因为编译器获取不到临时变量的地址
IntSet{}.String()

3. 在接口实现的时候会明确区分值类型和指针类型，比如，下面是编译不通过的，is是值类型：
var _ fmt.Stringer = is

查看strings.NewReader实现Reader接口
type Reader struct {
	s	string
	i	int64
	prevRune	int
}

func (r *Reader) Read(b []byte) (n int, err error) {
	if r.i > int64(len(r.s)) {
		return 0, io.EOF
	}
	r.prevRune = -1
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}
 */


func main() {
	var sr = strings.NewReader("hello")
	var r io.Reader = sr

	var rep = []byte("he*** neojos!")
	r.Read(rep)

	var w = bytes.Buffer{}
	sr.WriteTo(&w)

	fmt.Printf("%s\n", rep)
	fmt.Fprintf(&w, "%s", rep)
	fmt.Println(w.String())
}
