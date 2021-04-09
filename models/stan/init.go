package stan

import (
	"github.com/Nubes3/common-components/config"
	stan "github.com/nats-io/stan.go"
)

const (
	errSubj               = "nubes3_err"
	uploadFileSubj        = "nubes3_upload_file"
	downloadFileSubj      = "nubes3_download_file"
	stagingFileSubj       = "nubes3_staging_file"
	uploadFileSuccessSubj = "nubes3_upload_success_file"
	userSubj              = "nubes3_user"
	bucketSubj            = "nubes3_bucket"
	folderSubj            = "nubes3_folder"
	accessKeySubj         = "nubes3_accessKey"
	keyPairSubj           = "nubes3_keyPair"
)

var (
	sc stan.Conn
)

func InitStan() (func(), error) {
	url := config.Conf.StanUrl

	var err error
	sc, err = stan.Connect("nats-streaming", "nubes3", stan.NatsURL("nats://"+url))
	if err != nil {
		return nil, err
	}

	return cleanUp, nil
}

func cleanUp() {
	if sc != nil {
		_ = sc.Close()
	}
}
