package utils

import (
	"howett.net/plist"
	"os"
	"time"
	"vorta/models"
)

const wifiPlist = "/Library/Preferences/SystemConfiguration/com.apple.airport.preferences.plist"

type knownNetwork struct {
	AddedAt time.Time `plist:"AddedAt"`
	SSIDString   string `plist:"SSIDString"`
}

type appleAirportPreferences struct {
	KnownNetworks  map[string]knownNetwork `plist:"KnownNetworks"`
}

// Read plist of known Wifis and save them to settings DB.
// from https://godoc.org/howett.net/plist#example-Decoder-Decode
func UpdateWifiList() {
	var data appleAirportPreferences

	f, err := os.Open(wifiPlist)
	if err != nil {
		Log.Error("Issue reading Wifis from plist", err)
	}
	decoder := plist.NewDecoder(f)
	err = decoder.Decode(&data)
	if err != nil {
		Log.Error(err)
	}
	for _, network := range data.KnownNetworks {
		pp := []models.Profile{}
		models.DB.Find(&pp)
		for _, profile := range pp {
			newNetwork := models.KnownWifi{
				SSID:          network.SSIDString,
				LastConnected: network.AddedAt,
				Allowed:       true,
				ProfileID:     profile.ID,
			}
			models.DB.FirstOrCreate(&newNetwork, newNetwork)
		}
	}
}
