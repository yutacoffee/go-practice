//CSVファイルを読み込む
//「encoding/csv」パッケージの「Reader」型を使用する


package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	file, err := os.OpenFile("test.csv", os.O_RDONLY, 0)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer func() {
		file.Close()
	}()

	//csv.Readerを作成
	r := csv.NewReader(file)

	for {
		//1行ずつ読み込み
		record, err := r.Read()
		if err != nil {
			break
		}

		fmt.Println(record)
	}

}
