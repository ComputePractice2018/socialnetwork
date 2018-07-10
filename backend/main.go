package main

import "fmt"

//GetHelloWorldString возвращает строку "Hello World"
func GetHelloWorldString(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

func main() {
	fmt.Println(GetHelloWorldString("Xenia"))
}
