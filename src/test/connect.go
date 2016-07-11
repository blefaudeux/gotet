package main

import (
	"fmt"
	"gotet"
)

func main() {
	client := gotet.Client{}
	client.Connect("localhost", "6555")

	// Basic API tests. Check all status calls
	version, _ := client.Version()
	fmt.Println("Protocol version: ", version)

	trackerstate, _ := client.Trackerstate()
	fmt.Println("Tracker state: ", trackerstate)

	framerate, _ := client.Framerate()
	fmt.Println("Framerate: ", framerate)

	iscalibrated, _ := client.IsCalibrated()
	fmt.Println("Is calibrated ? ", iscalibrated)

	iscalibrating, _ := client.IsCalibrated()
	fmt.Println("Is calibrating ? ", iscalibrating)

	screenindex, _ := client.ScreenIndex()
	fmt.Println("Screen index: ", screenindex)

	srh, _ := client.ScreenResH()
	fmt.Println("Screen horizontal resolution: ", srh)

	srv, _ := client.ScreenResW()
	fmt.Println("Screen vertical resolution: ", srv)

	spw, _ := client.ScreenPsyW()
	fmt.Println("Screen physical width: ", spw)

	sph, _ := client.ScreenPsyH()
	fmt.Println("Screen physical height: ", sph)

	calibres, _ := client.CalibResult()
	fmt.Println("Calib results: \n", calibres)

	fd, _ := client.FrameData()
	fmt.Println("Frame data: \n", fd)
}
