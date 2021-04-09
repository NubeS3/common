package arangodb

import "time"

type keyPair struct {
	Public       string       `json:"public"`
	Private      string       `json:"private"`
	BucketId     string       `json:"bucket_id"`
	GeneratorUid string       `json:"generator_uid"`
	ExpiredDate  time.Time    `json:"expired_date"`
	Permissions  []Permission `json:"permissions"`
}

type KeyPair struct {
	Public       string    `json:"public"`
	Private      string    `json:"private"`
	BucketId     string    `json:"bucket_id"`
	GeneratorUid string    `json:"generator_uid"`
	ExpiredDate  time.Time `json:"expired_date"`
	Permissions  []string  `json:"permissions"`
}

func (k *keyPair) toKeyPair() *KeyPair {
	var perms []string
	for _, perm := range k.Permissions {
		perms = append(perms, perm.String())
	}

	return &KeyPair{
		Public:       k.Public,
		Private:      k.Private,
		BucketId:     k.BucketId,
		GeneratorUid: k.GeneratorUid,
		ExpiredDate:  k.ExpiredDate,
		Permissions:  perms,
	}
}
