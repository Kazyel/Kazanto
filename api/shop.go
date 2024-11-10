package api

import "fmt"

type Shop struct {
	Items []Item
}

func GetShop() {
	fmt.Println("Shop")

}
