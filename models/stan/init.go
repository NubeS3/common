//+build windows linux

package stan

import (
	"github.com/Nubes3/common/config"
	stan "github.com/nats-io/stan.go"
)

const (
	ErrSubj               = "nubes3_err"
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
	Sc stan.Conn
)

func InitStan() (func(), error) {
	url := config.Conf.StanUrl

	var err error
	Sc, err = stan.Connect("nats-streaming", "nubes3", stan.NatsURL("nats://"+url))
	if err != nil {
		return nil, err
	}

	return cleanUp, nil
}

func cleanUp() {
	if Sc != nil {
		_ = Sc.Close()
	}
}
