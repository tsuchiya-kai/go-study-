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

  ///// -------------int型------------- /////
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

  //// -------------float型------------- /////
  var fl64 float64 = 2.4
  fmt.Println(fl64)

  fl := 6.4 //暗黙的な定義の場合はintと違ってfloat64が適用される
  fmt.Println(fl64 + fl) //計算が可能

  var fl32 float32 = 2.4 //float32は基本的にはあまり使わない
  fmt.Println(fl32)
  
  //特殊な数値
  zero := 0.0
  pinf := 1.0 / zero //ポジティブインフと読む、正の無限大
  ninf := -1.0 / zero //ネガティブインフと読む、負の無限大
  nan := zero / zero
  fmt.Println(pinf, ninf, nan)

  // > 【uint型(+の整数）、complex(複素数型)】
  // > 他にもuint(+の整数) complex(複素数型）などの型があるが、あまり使用頻度は無いので今回は紹介だけにします。

  // var u8 uint = 255   //uint型
  // var c64 complex64 = -5 + 12i //complex型

  //// -------------string型------------- /////
  //複数行を扱うとき
  var st string = `テスト
  テスト
  テスト`
  fmt.Println(st)
  fmt.Println(string(st[0]))// stringはbyte型の配列として扱われるためこのようにして1文字目を取得できる, stringに変換する必要あり

  //// -------------string型------------- /////
}