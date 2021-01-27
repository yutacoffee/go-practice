//標準エラー出力に値を出力するサンプルコード

/*
これらの関数の第１パラメーターは出力先になるので
第一パラメーターには「os.Stderr」（標準エラー出力）を指定する
*/

package main 

import (
	"fmt"
	"os"
)

func main() {

	fmt.Fprint(os.Stderr, "北海道", "の降水確率は", 100, "%です\n")

	fmt.Fprintln(os.Stderr, "北海道", "の降水確率は", 100, "%です。")

	fmt.Fprintf(os.Stderr, "%sの降水確率は%d%%です。\n", "北海道", 100)
}