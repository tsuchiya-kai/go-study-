package session1 //ファイルごとにpackageとして命名する、この単位でimportできるっぽい

// ビルトインの機能もimportして使う
import (
	"fmt"
	"time"
)

//セッション1
func session1() {

  fmt.Println("hello world")
  fmt.Println(time.Now())


}