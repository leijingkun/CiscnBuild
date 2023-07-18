package main

import (
	"CiscnMap/utils"
	"fmt"
)

func main() {
	if utils.Isglastopf("36.152.44.95", 80) {
		fmt.Println("is a honeypot")
	}
}
