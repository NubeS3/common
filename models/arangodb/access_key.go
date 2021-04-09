//+build

package arangodb

import "time"

type accessKey struct {
	Key         string       `json:"key"`
	BucketId    string       `json:"bucket_id"`
	ExpiredDate time.Time    `json:"expired_date"`
	Permissions []Permission `json:"permissions"`
	Uid         string       `json:"uid"`
}

type AccessKey struct {
	Key         string    `json:"key"`
	BucketId    string    `json:"bucket_id"`
	ExpiredDate time.Time `json:"expired_date"`
	Permissions []string  `json:"permissions"`
	Uid         string    `json:"uid"`
}

func (a *accessKey) toAccessKey() *AccessKey {
	var perms []string
	for _, perm := range a.Permissions {
		perms = append(perms, perm.String())
	}

	return &AccessKey{
		Key:         a.Key,
		BucketId:    a.BucketId,
		ExpiredDate: a.ExpiredDate,
		Permissions: perms,
		Uid:         a.Uid,
	}
}
