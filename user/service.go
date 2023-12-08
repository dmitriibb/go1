package user

import (
	"fmt"
	"sync"
	"web-app1/db"
	"web-app1/model"
)

var lock = &sync.Mutex{}
var instance *PersonService

type PersonService struct {
	dbService *db.DbService
}

func GetAll() []model.User {
	return service.dbService.GetAllPerson()
}
func Save(p *model.User) {
	fmt.Println("PersonService - save", p)
	service.dbService.SavePerson(*p)
}

//func PersonServiceInstance() *PersonService {
//	if instance != nil {
//		return instance
//	}
//	lock.Lock()
//	defer lock.Unlock()
//	if instance == nil {
//		dbService := db.Instance()
//		instance = &PersonService{dbService}
//	}
//	return instance
//}

//type beanWrapper struct {
//	metadata factory.BeanMetadata
//	instance *PersonService
//}
//
//func (b *beanWrapper) BeanMetadata() *factory.BeanMetadata {
//	return &b.metadata
//}
//
//func BeanWrapper() *beanWrapper {
//	if beanWrapperInstance != nil {
//		return beanWrapperInstance
//	}
//	lock.Lock()
//	defer lock.Unlock()
//	if instance == nil {
//		metadata := factory.BeanMetadata{fmt.Sprintf("%T", instan)}
//		instance = &PersonService{}
//	}
//	return instance
//}
