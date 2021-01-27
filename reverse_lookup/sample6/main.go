//CSVファイルを出力する
package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	file, err := os.OpenFile("test.csv", os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println(err)

	os.Exit(1)
	}

	defer func() {
		file.Close()
	}()

	//cev.Writerを作成

w := csv.NewWriter(file)

w.Write([]string{"No.", "商品", "価格", "備考"})
w.Write([]string{"１", "キャベツ", "150", "とれたて新鮮"})
w.Write([]string{"2", "人参", "100", ""})
w.Write([]string{"3", "秋刀魚", "99", "今日の特価品です"})

w.Flush();

}