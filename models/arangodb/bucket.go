//+build windows linux

package arangodb

import "time"

type Bucket struct {
	Id     *string `json:"id,omitempty"`
	Uid    string  `json:"uid"`
	Name   string  `json:"name" binding:"required"`
	Region string  `json:"region" binding:"required"`
	// DB Info
	CreatedAt time.Time `json:"created_at"`
}
