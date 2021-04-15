//+build windows linux

package nats

import "github.com/Nubes3/common/utils"

const (
	GetById = iota
	Resolve
	Check
	GetByParams
	Add
	Update
	Remove
)

type ReqType int

type Msg struct {
	ReqType   ReqType  `json:"req_type"`
	Data      string   `json:"data"`
	ExtraData []string `json:"extra_data"`
}

type MsgResponse struct {
	IsErr     bool     `json:"is_err"`
	Data      string   `json:"data"`
	ExtraData []string `json:"extra_data"`
}

type ErrMsg struct {
	ErrType utils.ErrorType `json:"err_type"`
	Message string          `json:"message"`
}
