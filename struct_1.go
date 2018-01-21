/**
 *	数组的正确声明方式：var array [3]int = [3]int{1, 2, 3}
 *	数组是一个值类型，所有的值类型变量在赋值和作为参数传递时都将产生一次复制动作
 *	数组在作为参数传递时，需要指明长度. 函数声明中使用[]int类型的参数都是切片类型
 *	
 *	这个例子还有挺有帮助的：对一个struct的结构体声明方法，并且声明为指针对象，值对象的struct完全有不同的效果
 *	有一种类的感觉。跟PHP传对象还不一样，PHP中对象是指向对象引用地址的拷贝，而这里传值就是拷贝对象。
 */
package main


import (
	"fmt"
	"strconv"
)

type stack struct {
	point int
	data [10]int
}

func (s *stack) push(v int) {
	if s.point >= 9 {
		return
	}
	s.data[s.point] = v
	s.point++
}

func (s *stack) pop() int {
	s.point--
	return s.data[s.point]
}

func (s *stack) string() string{
	var str string
	for i := 0; i < s.point; i++ {
		str = str + "[" + 
				strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	} 
	return str
}

func main() {
	var s stack
	s.push(29)
	s.pop()
	s.push(12)
	fmt.Println(s.string())
	fmt.Printf("%v\n", s)
} 


