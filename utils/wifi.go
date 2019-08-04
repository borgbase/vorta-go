package utils

import (
	"fmt"
	"howett.net/plist"
	"os"
)

const wifiPlist = "/Library/Preferences/SystemConfiguration/com.apple.airport.preferences.plist"

type sparseBundleHeader struct {
	InfoDictionaryVersion string `plist:"CFBundleInfoDictionaryVersion"`
	BandSize              uint64 `plist:"band-size"`
	BackingStoreVersion   int    `plist:"bundle-backingstore-version"`
	DiskImageBundleType   string `plist:"diskimage-bundle-type"`
	Size                  uint64 `plist:"size"`
}

func UpdateWifiList() {
	var data sparseBundleHeader

	f, err := os.Open(wifiPlist)
	if err != nil {
		Log.Error("Issue reading Wifis from plist", err)
	}
	decoder := plist.NewDecoder(f)
	err = decoder.Decode(&data)
	if err != nil {
		Log.Error(err)
	}
	fmt.Println(data)

}
