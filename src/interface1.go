//go语言中interface的两种存在形式
//1. 提供函数的调用,比如io.Reader
//2. 作为一种值类型，包装其他不同类型的值
//因为interface作为值类型，没有任何的方法可以调用，所以出现了断言。x.(assert type)
package main

import (
	"fmt"
)

//
func sqlQuote(x interface{}) string {
	if x == nil {
		return "NULL"
	} else if _, ok := x.(int); ok {
		return fmt.Sprintf("%d", x)
	} else if _, ok := x.(uint); ok {
		return fmt.Sprintf("%d", x)
	} else if _, ok := x.(string); ok {
		return fmt.Sprintf("%s", x)
	} else if b, ok := x.(bool); ok {
		if b {
			return "TRUE"
		}
		return "FALSE"
	} else {
		panic(fmt.Sprintf("unexcepted type %T:%v", x, x))
	}
}

//方法中x转换成了不同作用域中的同名变量先
func sqlQuoteSwitch(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x)
	case bool :
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return fmt.Sprintf("%s", x)
	default:
		panic(fmt.Sprintf("unexcepted type %T:%v", x, x))
	}
}

func main() {
	var a interface{} = string("http://baby.360.cn")
	fmt.Println(sqlQuote(a))

	var x interface{} = bool(false)
	fmt.Println(sqlQuoteSwitch(x))
}