package main

import (
	"fmt"
)

// 请求生产
func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g iGun) {
	fmt.Printf("Gun: %s", getName())
	fmt.Println()
	fmt.Printf("Power: %s", getPower())
	fmt.Println()
}
