package nats

import (
	"github.com/Nubes3/common-components/config"
	nats "github.com/nats-io/nats.go"
)

const (
	mailSubj              = "nubes3_mail"
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
	nc *nats.Conn
)

func InitNats() (func(), error) {
	url := config.Conf.NatsUrl

	var err error
	nc, err = nats.Connect(url)
	if err != nil {
		return nil, err
	}

	return clearUp, nil
}

func clearUp() {
	if nc != nil {
		nc.Close()
	}
}
