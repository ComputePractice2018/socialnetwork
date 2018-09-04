package utils

import "fmt"

//GetHelloWorldString возвращает строку "Hello World"
func GetHelloWorld(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
