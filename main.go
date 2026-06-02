package main

import (
	"fmt"
	"log/slog"
)

const (
	Logging  = true            // Always on for logging. Set level with LogLevel property
	LogLevel = slog.LevelDebug // Change to Error for production
)

var Log *slog.Logger

func fatalOnErr(err error) {
	if err != nil {
		Log.Error(fmt.Sprintf("Fatal error encountered : %v. Program Aborting", err))
	}
}

func main() {
	// Configure logging
	Log = ConfigureLogging(LogLevel, Logging)

	w, err := NewWifi("wlan0", "rpi-config-ap")
	fatalOnErr(err)

	fatalOnErr(w.RescanInfo())

	if w.IsConnectedToNetwork() {
		Log.Info("Connected to wireless network %s with IP: %s\n", w.GetESSID(), w.GetIP())
	} else {
		Log.Error("Not connected to WIFI - ToDo: Enable AP here!\n")
		// Call function to get list if available networks and connect to one of them.
		availableSSID, err := w.GetAvailableNetworks()
		fatalOnErr(err)
		ss := make([]string, 0, len(availableSSID))
		for _, b := range availableSSID {
			ss = append(ss, string(b))
		}
		// log.Printf("Available networks:\n%s", strings.Join(ss, "\n"))
	}

	// The following blocks are for testing extract info from the nmcli output.
	/* availableSSID, err := w.GetAvailableNetworks()
	fatalOnErr(err)
	ss := make([]string, 0, len(availableSSID))
	for _, b := range availableSSID {
		ss = append(ss, string(b))
	}
	log.Printf("Available networks:\n%s", strings.Join(ss, "\n")) */

}

////////////////////////////////////////////////////////////////////////////////
