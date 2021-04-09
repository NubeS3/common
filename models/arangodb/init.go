package arangodb

import (
	"context"
	"github.com/Nubes3/common-components/config"
	arangoDriver "github.com/arangodb/go-driver"
	arangoHttp "github.com/arangodb/go-driver/http"
	"time"
)

const ContextExpiredTime = 30

var (
	arangoConnection arangoDriver.Connection
	arangoClient     arangoDriver.Client
	arangoDb         arangoDriver.Database
	userCol          arangoDriver.Collection
)

func InitArango() {
	var err error
	arangoConnection, err = arangoHttp.NewConnection(arangoHttp.ConnectionConfig{
		Endpoints: []string{config.Conf.ArangoHost},
	})

	if err != nil {
		panic(err)
	}

	arangoClient, err = arangoDriver.NewClient(arangoDriver.ClientConfig{
		Connection:     arangoConnection,
		Authentication: arangoDriver.BasicAuthentication(config.Conf.ArangoUser, config.Conf.ArangoPassword),
	})

	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), ContextExpiredTime*time.Second)
	defer cancel()

	dbExist, err := arangoClient.DatabaseExists(ctx, "nubes3")
	if err != nil {
		panic(err)
	}

	if !dbExist {
		arangoDb, _ = arangoClient.CreateDatabase(ctx, "nubes3", &arangoDriver.CreateDatabaseOptions{
			Users: []arangoDriver.CreateDatabaseUserOptions{
				{
					UserName: config.Conf.ArangoUser,
					Password: config.Conf.ArangoPassword,
				},
			},
		})
	} else {
		arangoDb, _ = arangoClient.Database(ctx, "nubes3")
	}
}
