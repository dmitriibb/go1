package db

import (
	"errors"
	"fmt"
	"github.com/reugn/async"
	"reflect"
)

type CommonInMemoryRepository struct {
	lock *async.ReentrantLock
	data []any
}

func New() *CommonInMemoryRepository {
	lock := &async.ReentrantLock{}
	data := make([]any, 0)
	return &CommonInMemoryRepository{lock, data}
}
func (repo *CommonInMemoryRepository) Save(entity any) error {
	if err := isNotStruck(entity); err != nil {
		return err
	}
	repo.lock.Lock()
	defer repo.lock.Unlock()
	repo.data = append(repo.data, entity)
	return nil
}
func (repo *CommonInMemoryRepository) All() (result []any) {
	repo.lock.Lock()
	defer repo.lock.Unlock()
	result = make([]any, len(repo.data))
	for i, v := range repo.data {
		result[i] = v
	}
	return
}
func (repo *CommonInMemoryRepository) Delete(entity any) (bool, error) {
	if err := isNotStruck(entity); err != nil {
		return false, err
	}
	repo.lock.Lock()
	defer repo.lock.Unlock()
	if index := repo.find(entity); index > -1 {
		repo.data = append(repo.data[:index], repo.data[index+1:]...)
		return true, nil
	} else {
		return false, nil
	}
}

func (repo *CommonInMemoryRepository) find(entity any) int {
	repo.lock.Lock()
	defer repo.lock.Unlock()
	for i, v := range repo.data {
		fmt.Printf("checking index %v, for delete\n", i)
		if reflect.DeepEqual(v, entity) {
			return i
		}
	}
	return -1
}

func isNotStruck(entity any) (err error) {
	kind := reflect.TypeOf(entity).Kind()
	if kind != reflect.Struct {
		err = errors.New(fmt.Sprintf("Entity has type kind %v, expected only struct", kind))
	}
	return
}
