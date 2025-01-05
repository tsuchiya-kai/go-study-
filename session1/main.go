package main //ファイルごとにpackageとして命名する、この単位でimportできるっぽい。 命名がmainである必要がありそう

// ビルトインの機能もimportして使う
import (
	"fmt"
	"time"
)

//命名がmainである必要がありそう
func main() {

  fmt.Println("hello world")
  fmt.Println(time.Now())


}