//配列とスライスを同時に作成する

package main 

import(
	"fmt"
)

func main() {
	s1 := make([]int, 10, 20)
	
	fmt.Println(s1)
	fmt.Println("len:", len(s1))
	fmt.Println("cap:", cap(s1))

	//空行
	fmt.Println()

	//キャパシティは省略化
	s2 := make([]float32, 5)

	fmt.Println(s2)
	fmt.Println("len:", len(s2))
	fmt.Println("cap:", cap(s2))
}