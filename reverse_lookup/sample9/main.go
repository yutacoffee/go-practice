package main

import "os"

func main() {

	os.MkdirAll("a/b/c/f", 0777)
	os.MkdirAll("x/y/z", 0777)

	os.RemoveAll("x")

	file, _ := os.Create("test")
	file.Close()

	os.Remove("test")
}