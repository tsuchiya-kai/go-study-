package main //ファイルごとにpackageとして命名する、この単位でimportできるっぽい

// ビルトインの機能もimportして使う
import (
	"fmt"
	"time"
)

func main() {
  fmt.Println("hello world")
  fmt.Println(time.Now())
}