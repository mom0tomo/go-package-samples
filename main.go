package main

import (
	"fmt"

	"github.com/mom0tomo/go-package-samples/rodentia"
)

func main() {
	fmt.Println(AppName())

	fmt.Println(rodentia.GopherLike())
	fmt.Println(rodentia.BeaverLike())
	fmt.Println(rodentia.HamsterLike())
}
