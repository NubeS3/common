package common

import (
	_ "github.com/Nubes3/common/config"
	"github.com/Nubes3/common/models/arangodb"
	"github.com/Nubes3/common/models/nats"
	"github.com/Nubes3/common/models/seaweedfs"
	"github.com/Nubes3/common/models/stan"
)

func InitCoreComponents(initArango, initSeaweedfs, initStan, initNats bool) []func() {
	cleanUps := []func(){}

	if initArango {
		arangodb.InitArango()
	}
	if initSeaweedfs {
		clean, err := seaweedfs.InitFs()
		if err != nil {
			panic(err)
		}

		cleanUps = append(cleanUps, clean)
	}
	if initNats {
		clean, err := nats.InitNats()
		if err != nil {
			panic(err)
		}

		cleanUps = append(cleanUps, clean)
	}
	if initStan {
		clean, err := stan.InitStan()
		if err != nil {
			panic(err)
		}

		cleanUps = append(cleanUps, clean)
	}

	return cleanUps
}
