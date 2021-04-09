//+build windows linux

package nats

const (
	GetById = iota
	Resolve
)

type Msg struct {
	ReqType   string  `json:"req_type"`
	Data      string  `json:"data"`
	ExtraData *string `json:"extra_data,omitempty"`
}
