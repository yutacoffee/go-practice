//標準入力から値を読み込むサンプルコード

package main

import "fmt"

func main(){

	var fname, name string
	fmt.Println("あなたの名前は(姓 名)?")
	fmt.Scanln(&fname, &name)

	fmt.Println("あなたの年齢は？")
	var age int
	fmt.Scanf("%d", &age) //scanfは書式を指定して読み込む

	fmt.Printf("名前:%s %s 年齢: %d\n", fname, name, age)
}