package main

import "fmt"
import "flag"
import "./utils"

func main() {
	var name = flag.String("name", "Xenia", "имя для приветствия")
	flag.Parse()
	fmt.Println(utils.GetHelloWorld(*name))
}
