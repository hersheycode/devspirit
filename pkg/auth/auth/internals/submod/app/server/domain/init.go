package domain

import (
	dbClient "apppathway.com/pkg/db_api/internals/project/pkg/db_client"
	dbTypes "apppathway.com/pkg/db_api/internals/project/pkg/types"
	dt "apppathway.com/pkg/user/auth/internals/project/pkg/types"
)

type Services dt.DomainServices
type Logic struct {
	dbTech    string
	dbService dbTypes.DBClientService
	store     dbTypes.Client
}

func InitServices(dbTech string, dbService dbTypes.DBClientService, streams dbTypes.DBStreams) Logic {
	queryChan := streams.RecvQueryChan()
	mutationChan := streams.RecvMutationChan()
	return Logic{dbTech: dbTech, dbService: dbService, store: dbClient.NewClient(dbTech, queryChan, mutationChan, dbService)}
}
