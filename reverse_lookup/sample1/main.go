//値を文字列にフォーマットするサンプルコード

package main

import "fmt"

func main(){

	s := fmt.Sprintf("%sの降水確率は%d%%です", "沖縄", 30)
	
	fmt.Println(s)


}