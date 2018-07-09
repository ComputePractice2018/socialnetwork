# socialnetwork

## Usercases
1. Как пользователь, я хочу иметь возможность создать свою страницу(имя, фамилия, email, github), чтобы иметь свой профиль.
1. Как пользователь, я хочу иметь возможность редактировать свой профиль, чтобы исправить ошибки при создании.
1. Как пользователь, я хочу иметь возможность добавить друга, чтобы добавить своего друга.
1. Как пользователь, я хочу иметь возможность просмотреть страницу друзей, чтобы использовать эту информацию.
1. Как пользователь, я хочу иметь возможность удалить из друзей, чтобы не хранить неактуальную информацию.

## REST API

### POST /api/socialnetwork/users

Тело запроса:

```json
{
   "name": "Имя",
   "surname": "Фамилия",
   "email": "user@domain.ru",
   "github" "user"
}
```
Ответ: 201 Created
Location: /api/socialnetwork/users/1

### PUT /api/socialnetwork/users/1

Тело запроса:

```json
   {
        "name": "Имя",
        "surname": "Фамилия",
        "email": "user@domain.ru",
        "github" "user"
   }
```
Ответ: 202 Accepted
Location: /api/socialnetwork/users/1

### POST /api/socialnetwork/users/2/friends

Тело запроса:

```json
{
   "name": "Имя",
   "surname": "Фамилия",
   "email": "user@domain.ru",
   "github" "user"
}
```
Ответ: 201 Created
Location: /api/socialnetwork/users/friends/2

### GET /api/socialnetwork/users

Ответ: 200 OK

```json
[{
   "name": "Имя",
   "surname": "Фамилия",
   "email": "user@domain.ru",
   "github" "user"
}]
```

### DELETE /api/socialnetwork/users/friends/2

Ответ: 204 No Content

