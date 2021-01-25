package main

import(
	"fmt"
)

func main() {

	s := []string{"a", "b", "c"}

	test(s...)

	//こうした時と渡される値は同じ
	test("a", "b", "c")


}
func test(s ...string) {

	fmt.Println(len(s), s)
}