package main //ファイルごとにpackageとして命名する、この単位でimportできるっぽい。 命名がmainである必要がありそう

import "fmt"

// 明示的な定義は関数スコープ外で定義することが可能
var outsideI int = 1000

// 分離した関数
func other() {
  var s4 string = "other"
  fmt.Println(s4)
}

func main() {
  /*
  * int型の最大値/最小値 参考: https://zenn.dev/heromina/scraps/760807f6990659
  * int8
  * int16
  * int32
  * int64
  */
  var i int = 100 //このように宣言するとどのビットかは環境依存になる

  var i64 int64 = 10000
  // fmt.Println(i + i64) //型が違うと計算できない error: invalid operation: i + i64 (mismatched types int and int64)compilerMismatchedTypes

  fmt.Println(i)

  //現在の型を調べる
  fmt.Printf("%T\n",i64)

  // 型の変換
  fmt.Printf("%T\n",int32(i64))
}