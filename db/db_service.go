package db

import (
	"fmt"
	"sync"
	"web-app1/model"
)

var lock = sync.Mutex{}
var instance *DbService

type DbService struct {
	personTable []model.User
}

func (dbService *DbService) SavePerson(person model.User) {
	fmt.Println("DbService - save", person)
	dbService.personTable = append(dbService.personTable, person)
}
func (dbService *DbService) GetAllPerson() []model.User {
	fmt.Println("DbService - GetAllPerson. personTable size =", len(dbService.personTable))
	result := make([]model.User, len(dbService.personTable))
	copy(result, dbService.personTable)
	return result
}

func Instance() *DbService {
	if instance != nil {
		return instance
	}
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = &DbService{make([]model.User, 0)}
	}
	return instance
}
