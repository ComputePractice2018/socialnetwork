package data

import (
	"fmt"
)
//User структура для хранения данных пользователя
type User struct {
	Name    string 'json:"name"'
    Surname string 'json:"surname"'
    Email   string 'json:"email"'
    Github  string 'json:"github"'
} 

//users хранимый список пользователей
var users []User

//GetUsers возвращает список пользователей
func GetUsers() []User{
    return users
}

//AddUsers добавляет пользователя в конец списка и возвращает id
func AddUsers(user User) int {
    id := len(users)
    users  = append(users,user)
    return id
}

//EditUsers изменяет пользователя с id на users
func EditUsers(user User, id int) error {
    if id < 0 || id >= len(users){
        return fmt.Errorf("incorrect ID")
    }
    users[id] = user
    return nil
}

//RemoveUsers удаляет контакт по id
func RemoveUsers(id int) error {
    if id < 0 || id >= len(users){
        return fmt.Errorf("incorrect ID")
    }
    users = append(users[:id],users[id+1:]...)
    return nil

}