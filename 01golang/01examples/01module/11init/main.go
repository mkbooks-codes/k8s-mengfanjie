package main

import (
	"fmt"

	_ "github/mkbooks-codes/k8s-mengfanjie/01golang/01examples/01module/11init/a"
	_ "github/mkbooks-codes/k8s-mengfanjie/01golang/01examples/01module/11init/b"
)

func init() {
	fmt.Println("main init")
}
func main() {
	fmt.Println("main")
}
