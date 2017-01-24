package main

import "fmt"

func CallWith(f func(string) int, who string) {
	fmt.Println(f(who))
}

type FunctionHolder struct {
	Function func(string) int
}

func main() {
	holder := FunctionHolder{func(who string) int { fmt.Println("Hello,", who); return 1 }}
	CallWith(holder.Function, "ernest")
}
