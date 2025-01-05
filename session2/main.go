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
  //明示的な変数定義
  //var 変数名 型 = 値
  var i int = 100
  var s string = "hello Go"

  //まとめて定義
  var t,f bool = true, false
  var(
    i2 int = 100
    s2 string = "string2"
  )

  fmt.Println(i,s,t,f,i2,s2)

  //値の更新
  i = 200
  fmt.Println(i)

  //暗黙的な定義
  //変数名 := 値
  i4 := 400
  fmt.Println(i4)

  // i4 := 500 //エラーになる: session2/main.go:29:6: no new variables on left side of :=
  // i4 = "hoge" //エラーになる: session2/main.go:30:8: cannot use "hoge" (untyped string constant) as int value in assignment

  fmt.Println(outsideI)

  other()
}