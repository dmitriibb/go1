package user

import (
	"fmt"
)

func GetAll() []User {
	return Repo.All()
}
func Save(user *User) {
	fmt.Println("USerService - save", user)
	Repo.Save(*user)
}

func Delete(user *User) bool {
	return Repo.Delete(*user)
}

//func PersonServiceInstance() *PersonService {
//	if instance != nil {
//		return instance
//	}
//	lock.Lock()
//	defer lock.Unlock()
//	if instance == nil {
//		dbService := repository.Instance()
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
//		metadata := factory.BeanMetadata{fmt.Sprintf("%T", instant)}
//		instance = &PersonService{}
//	}
//	return instance
//}
