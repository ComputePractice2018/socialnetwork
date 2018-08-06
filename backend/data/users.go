package data
//User структура для хранения данных пользователя
type User struct {
	Name    string 'json:"name"'
    Surname string 'json:"surname"'
    Email   string 'json:"email"'
    Github  string 'json:"github"'
} 

//UserList хранимый список пользователей
var UserList = []User{User {
   Name: "Имя",
   Surname: "Фамилия",
   Email: "user@domain.ru",
   Github: "user"}}