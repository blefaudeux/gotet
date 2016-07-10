package main

import (
	"fmt"
	"gotet"
)

func main() {
	client := gotet.Client{}
	client.Connect("localhost", "6555")
	version, _ := client.Version()
	fmt.Println("Protocol version: ", version)
}