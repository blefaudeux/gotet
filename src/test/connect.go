package main

import (
	"fmt"
	"gotet"
)

func main() {
	client := gotet.Client{}
	client.Connect("localhost", "6555")

	// Basic API tests
	version, _ := client.Version()
	fmt.Println("Protocol version: ", version)

	iscalibrated, _ := client.IsCalibrated()
	fmt.Println("Is calibrated ? ", iscalibrated)

	iscalibrating, _ := client.IsCalibrated()
	fmt.Println("Is calibrating ? ", iscalibrating)
}
