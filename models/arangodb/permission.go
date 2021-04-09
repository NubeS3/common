//+build windows linux

package arangodb

import "github.com/Nubes3/common/utils"

type Permission int

const (
	GetFileList Permission = iota
	GetFileListHidden
	Download
	DownloadHidden
	Upload
	MarkHidden
	DeleteFile
	RecoverFile
)

func (perm Permission) String() string {
	return [...]string{
		"GetFileList",
		"GetFileListHidden",
		"Download",
		"DownloadHidden",
		"Upload",
		"MarkHidden",
		"DeleteFile",
		"RecoverFile",
	}[perm]
}

func parsePerm(p string) (Permission, error) {
	switch p {
	case "GetFileList":
		return GetFileList, nil
	case "GetFileListHidden":
		return GetFileListHidden, nil
	case "Download":
		return Download, nil
	case "DownloadHidden":
		return DownloadHidden, nil
	case "Upload":
		return Upload, nil
	case "MarkHidden":
		return MarkHidden, nil
	case "DeleteFile":
		return DeleteFile, nil
	case "RecoverFile":
		return RecoverFile, nil
	default:
		return -1, &utils.ModelError{
			Msg:     "invalid permission: " + p,
			ErrType: utils.Invalid,
		}
	}
}
