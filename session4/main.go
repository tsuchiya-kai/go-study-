package main

import "fmt"

// 定数

const Pi = 3.14 // 頭文字を大文字にすることでグローバル化する様

//変数の場合はintの最大値を超えて定義できない
// var Big int = 9223372036854775807 + 1
//定数の場合は可能
const Big  = 9223372036854775807 + 1

//連続する整数の連番を生成する iota
const(
	c0 = iota
	c1
	c2
)

func main() {
	fmt.Println(Pi,
		Big-1,//ただし、使用時にintの制限以下にする必要がありそう
	)

  fmt.Println(
		c0,
		c1,
		c2,
	)
}