package factory

import (
	"sync"
	"web-app1/db"
	"web-app1/infra/context"
)

const (
	DbServiceType     = "dbServiceType"
	PersonServiceType = "personServiceType"
)

var lock = sync.Mutex{}
var existingDependencies = make(map[string]any)
var typeToCreationFuncMap = map[string]interface{}{
	DbServiceType: getDbServiceInstance,
}

func getDbServiceInstance() *db.DbService {
	res, ok := existingDependencies[DbServiceType]
	if ok {
		return res.(*db.DbService)
	} else {
		dbService := db.Instance()
		existingDependencies[DbServiceType] = dbService
		return dbService
	}
}

type DependenciesRequest struct {
	ResponseChan    chan context.BeanMetadata
	NeededBeanTypes []string
}
