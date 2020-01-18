package main

import (
	"fmt"

	"./ProtoType"
)

func main() {
	fmt.Println("prototype")
	result, error := ProtoType.GetTshirClone(ProtoType.WHITE)

	if error == nil {
		fmt.Println(result.GetInfo())
	}
}
