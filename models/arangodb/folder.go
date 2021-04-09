//+build windows linux

package arangodb

type Folder struct {
	Id       string        `json:"-"`
	OwnerId  string        `json:"owner_id"`
	Name     string        `json:"name"`
	Fullpath string        `json:"fullpath"`
	Children []FolderChild `json:"children"`
}

type FolderChild struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	IsHidden bool   `json:"is_hidden"`
}
