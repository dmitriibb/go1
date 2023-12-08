package user

import (
	"fmt"
	"web-app1/common/db"
)

var Repo UserRepository = &InMemoryRepository{db.New()}

type UserRepository interface {
	Save(user User)
	All() []User
	Delete(user User) bool
}

type InMemoryRepository struct {
	internalRepo *db.CommonInMemoryRepository
}

func (repo *InMemoryRepository) Save(user User) {
	repo.internalRepo.Save(user)
}

func (repo *InMemoryRepository) All() []User {
	allAny := repo.internalRepo.All()
	result := make([]User, len(allAny))
	for i, v := range allAny {
		result[i] = v.(User)
	}
	return result
}

func (repo *InMemoryRepository) Delete(user2 User) bool {
	res, err := repo.internalRepo.Delete(user2)
	if err != nil {
		fmt.Printf("ERROR - %v\n", err)
	}
	return res
}
