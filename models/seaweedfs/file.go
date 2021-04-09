package seaweedfs

import (
	"github.com/Nubes3/common-components/utils"
	"github.com/linxGnu/goseaweedfs"
	"io"
)

func UploadFile(filename string, size int64, reader io.Reader) (*goseaweedfs.FilePart, error) {
	meta, err := sw.Upload(reader, filename, size, "", "")
	if err != nil {
		return nil, &utils.ModelError{
			Msg:     err.Error(),
			ErrType: utils.FsError,
		}
	}

	return meta, nil
}

func DownloadFile(id string, callback func(reader io.Reader) error) error {
	_, err := sw.Download(id, nil, callback)

	if err != nil {
		return &utils.ModelError{
			Msg:     err.Error(),
			ErrType: utils.FsError,
		}
	}

	return nil
}
