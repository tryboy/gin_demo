package main

import "fmt"

func main() {
	adidasFactory, _ := getSportsFactory("adidas")
	nikeFactory, _ := getSportsFactory("nike")

	nikeShoe := makeShoe()
	nikeShirt := makeShirt()

	adidasShoe := makeShoe()
	adidasShirt := makeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

// 规定类型约束传入值 - 无论什么牌子都好，有什么华丽呼哨的功能都好，首要是具备鞋子的基本属性
func printShoeDetails(s iShoe) {
	fmt.Printf("Logo: %s", getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", getSize())
	fmt.Println()
}

// 规定类型约束传入值 - 无论什么牌子都好，有什么华丽呼哨的功能都好，首要是具备衣服的基本属性
func printShirtDetails(s iShirt) {
	fmt.Printf("Logo: %s", getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", getSize())
	fmt.Println()
}
