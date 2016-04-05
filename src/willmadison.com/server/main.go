package main

import (
	"fmt"
	"willmadison.com/flavor"
)

func main() {
	var f flavor.Flavor

	f = flavor.Sweet(16)

	fmt.Println(f.Taste())
}
