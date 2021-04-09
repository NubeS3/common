//+build

package arangodb

import (
	"context"
	"github.com/Nubes3/common/config"
	arangoDriver "github.com/arangodb/go-driver"
	arangoHttp "github.com/arangodb/go-driver/http"
	"time"
)

const ContextExpiredTime = 30

var (
	ArangoConnection arangoDriver.Connection
	ArangoClient     arangoDriver.Client
	ArangoDb         arangoDriver.Database
)

func InitArango() {
	var err error
	ArangoConnection, err = arangoHttp.NewConnection(arangoHttp.ConnectionConfig{
		Endpoints: []string{config.Conf.ArangoHost},
	})

	if err != nil {
		panic(err)
	}

	ArangoClient, err = arangoDriver.NewClient(arangoDriver.ClientConfig{
		Connection:     ArangoConnection,
		Authentication: arangoDriver.BasicAuthentication(config.Conf.ArangoUser, config.Conf.ArangoPassword),
	})

	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), ContextExpiredTime*time.Second)
	defer cancel()

	dbExist, err := ArangoClient.DatabaseExists(ctx, "nubes3")
	if err != nil {
		panic(err)
	}

	if !dbExist {
		ArangoDb, _ = ArangoClient.CreateDatabase(ctx, "nubes3", &arangoDriver.CreateDatabaseOptions{
			Users: []arangoDriver.CreateDatabaseUserOptions{
				{
					UserName: config.Conf.ArangoUser,
					Password: config.Conf.ArangoPassword,
				},
			},
		})
	} else {
		ArangoDb, _ = ArangoClient.Database(ctx, "nubes3")
	}
}
