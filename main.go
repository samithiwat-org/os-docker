package main

import (
	"fmt"
	"os"
)

func main() {
	secretUser := os.Getenv("SECRET_USER")

	fmt.Println(fmt.Sprintf("Hello, %s", secretUser))
}
