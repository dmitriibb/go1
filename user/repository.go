package user

import (
	"database/sql"
	"fmt"
	"web-app1/common/db"
)

// var Repo UserRepository = &InMemoryRepository{db.New()}
var Repo UserRepository = &DbRepository{}

type UserRepository interface {
	DbSave(user User)
	DbAll() []User
	DbDelete(user User) bool
}
type InMemoryRepository struct {
	internalRepo *db.CommonInMemoryRepository
}

type DbRepository struct{}

func (repo *InMemoryRepository) DbSave(user User) {
	repo.internalRepo.Save(user)
}
func (repo *InMemoryRepository) DbAll() []User {
	allAny := repo.internalRepo.All()
	result := make([]User, len(allAny))
	for i, v := range allAny {
		result[i] = v.(User)
	}
	return result
}
func (repo *InMemoryRepository) DbDelete(user2 User) bool {
	res, err := repo.internalRepo.Delete(user2)
	if err != nil {
		fmt.Printf("ERROR - %v\n", err)
	}
	return res
}

func (repo *DbRepository) DbSave(user User) {
	query := "INSERT INTO usr (name, age) VALUES "
	queryParams := fmt.Sprintf("('%v', %d)", user.Name, user.Age)
	query = query + queryParams
	f := func(db *sql.DB) any {
		res, err := db.Exec(query)
		if err != nil {
			fmt.Println("ERROR - ", err)
		}
		return res
	}
	res := db.UseConnection(f)
	fmt.Println("Save user result", res)
}
func (repo *DbRepository) DbAll() []User {
	userList := make([]User, 0)
	query := "SELECT id, name, age FROM usr"
	f := func(db *sql.DB) any {
		rows, err := db.Query(query)
		if err != nil {
			fmt.Println("ERROR - ", err)
		}
		for rows.Next() {
			user := User{}
			if err := rows.Scan(&user.Id, &user.Name, &user.Age); err != nil {
				fmt.Println("ERROR - ", err)
			} else {
				userList = append(userList, user)
			}
		}
		return userList
	}
	dbRes := db.UseConnection(f)
	fmt.Println("select all users db result", dbRes)
	return userList
}
func (repo *DbRepository) DbDelete(user User) bool {
	if user.Id < 1 {
		fmt.Println("ERROR - ", "can't delete user with id 0")
		return false
	}
	query := fmt.Sprintf("DELETE FROM usr where id = %d", user.Id)
	f := func(db *sql.DB) any {
		res, err := db.Exec(query)
		if err != nil {
			fmt.Println("ERROR - ", err)
		}
		return res
	}
	dbRes := db.UseConnection(f)
	fmt.Println("delete user db result", dbRes)
	return true
}
