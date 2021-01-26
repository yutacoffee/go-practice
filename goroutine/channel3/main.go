//チャネルを利用した同期
package main 
import(
	"fmt"
	"math/rand"
	"time"
)

const goroutines = 3

func main() {
	c := make(chan int)
 
	for i := 0; i < goroutines; i++ {

	go func(s chan<- int) {

		//ダミーの処理として乱数により1~10秒待機
		time.Sleep(
			time.Duration(rand.Int63n(10)) * time.Second)

			fmt.Println("処理完了")

			s <- 0
	}(c)
}

	for i := 0; i < goroutines; i++ {
			<-c
	}

	fmt.Println("全て完了")
}