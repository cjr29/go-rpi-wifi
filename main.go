package main

////////////////////////////////////////////////////////////////////////////////

import (
	"go-rpi-wifi/wifi"
	"log"
)

////////////////////////////////////////////////////////////////////////////////

func fatalOnErr(err error) {
	if err != nil {
		log.Fatalf(`Fatal error encountered : %s.
Program Aborting
`, err.Error())
	}
}

////////////////////////////////////////////////////////////////////////////////

func main() {
	// fatalOnErr(checkDependencies())

	w, err := wifi.New("wlan0", "rpi-config-ap")
	fatalOnErr(err)

	fatalOnErr(w.RescanInfo())

	if w.IsConnectedToNetwork() {
		log.Printf("Connected to wireless network with IP: %s\n", w.GetIP())
	} else {
		log.Printf("Not connected to WIFI - ToDo: Enable AP here!\n")
		// Call function to get list if available networks and connect to one of them.
		availableSSID, err := w.GetAvailableNetworks()
		fatalOnErr(err)
		log.Printf("Available networks:\n%s", string(availableSSID))
	}
	availableSSID, err := w.GetAvailableNetworks()
	fatalOnErr(err)
	log.Printf("Available networks:\n%s", string(availableSSID))
}

func init() {
	log.SetPrefix("")
	log.SetFlags(0)
}

////////////////////////////////////////////////////////////////////////////////
