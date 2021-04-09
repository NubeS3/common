//+build windows linux

package nats

import (
	"github.com/Nubes3/common/config"
	nats "github.com/nats-io/nats.go"
)

const (
	MailSubj              = "nubes3_mail"
	UploadFileSubj        = "nubes3_upload_file"
	DownloadFileSubj      = "nubes3_download_file"
	StagingFileSubj       = "nubes3_staging_file"
	UploadFileSuccessSubj = "nubes3_upload_success_file"
	UserSubj              = "nubes3_user"
	BucketSubj            = "nubes3_bucket"
	FolderSubj            = "nubes3_folder"
	AccessKeySubj         = "nubes3_accessKey"
	KeyPairSubj           = "nubes3_keyPair"
)

var (
	Nc *nats.Conn
)

func InitNats() (func(), error) {
	url := config.Conf.NatsUrl

	var err error
	Nc, err = nats.Connect(url)
	if err != nil {
		return nil, err
	}

	return clearUp, nil
}

func clearUp() {
	if Nc != nil {
		Nc.Close()
	}
}
