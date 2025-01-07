package main

import "fmt"

//演算子

func main() {
	//四則演算はJSと同じ
  fmt.Println(1+2)

	//+に関しては文字列でも可能
	fmt.Println("Hello" + " " + "World")

	/////---------比較演算子---------/////
	fmt.Println(1 == 1)
	fmt.Println(1 != 1)
	//他はJSと同じ

	/////---------論理演算子---------/////
	//JSと同じ
	fmt.Println(true && false == true)
	fmt.Println(true || false == true)
	fmt.Println(!true)
}