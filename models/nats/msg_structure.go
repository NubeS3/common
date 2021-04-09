//+build windows linux

package nats

const (
	GetById = iota
	Resolve
)

type ReqType int

type Msg struct {
	ReqType   ReqType `json:"req_type"`
	Data      string  `json:"data"`
	ExtraData *string `json:"extra_data,omitempty"`
}
