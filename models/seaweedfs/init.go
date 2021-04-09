//+build windows linux

package seaweedfs

import (
	"github.com/Nubes3/common/config"
	"github.com/linxGnu/goseaweedfs"
	"net/http"
	"time"
)

var (
	Sw *goseaweedfs.Seaweed
)

const (
	CHUNK_SIZE = 8096
)

// Return a cleanup function if success, please defer call this function to clean up fs.
func InitFs() (func(), error) {
	masterUrl := config.Conf.SeaweedMasterUrl

	var err error
	Sw, err = goseaweedfs.NewSeaweed(masterUrl, []string{}, CHUNK_SIZE, &http.Client{Timeout: 5 * time.Minute})

	if err != nil {
		return nil, err
	}

	return cleanUp, nil
}

func cleanUp() {
	if Sw != nil {
		_ = Sw.Close()
	}
}
