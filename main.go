package main

import (
	"fmt"
)

func main() {

	ReadSettings()
	connectAndAuthorize()
	users := searchForUsers()

	fmt.Println(users)
}
