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
		log.Printf("Connected to wireless network %s with IP: %s\n", w.GetESSID(), w.GetIP())
	} else {
		log.Printf("Not connected to WIFI - ToDo: Enable AP here!\n")
		// Call function to get list if available networks and connect to one of them.
		availableSSID, err := w.GetAvailableNetworks()
		fatalOnErr(err)
		ss := make([]string, 0, len(availableSSID))
		for _, b := range availableSSID {
			ss = append(ss, string(b))
		}
		// log.Printf("Available networks:\n%s", strings.Join(ss, "\n"))
	}

	// The following blocks are for testing the regexes to extract info from the nmcli output.
	/* availableSSID, err := w.GetAvailableNetworks()
	fatalOnErr(err)
	ss := make([]string, 0, len(availableSSID))
	for _, b := range availableSSID {
		ss = append(ss, string(b))
	}
	log.Printf("Available networks:\n%s", strings.Join(ss, "\n")) */

}

func init() {
	log.SetPrefix("")
	log.SetFlags(0)
}

////////////////////////////////////////////////////////////////////////////////
