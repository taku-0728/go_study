// 例3-20
package main

import (
	"fmt"

	"golang.org/x/net/idna"
)

func main5() {
	src := "握力王"
	ascii, err := idna.ToASCII(src)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s -> %s\n", src, ascii)
}
