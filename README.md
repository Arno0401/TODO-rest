
# API документация для Todo приложения

## Описание
Этот проект представляет собой простое приложение для управления задачами (Todo), в котором реализованы возможности регистрации, авторизации и работы с задачами: создание, получение, обновление и удаление. Приложение использует JWT для авторизации пользователей и взаимодействует с базой данных через ORM GORM.

## Основные маршруты

### **POST** `/sign_up`
Регистрация нового пользователя.

**Запрос:**
- Тело запроса должно содержать JSON с полями:
```json
    {
       "user_name": "John",
       "login": "John1234",
       "password": "12312john!"
    }
```
####  Логин должен содержать только латинские символы от 5 символов
#### Пароль должен содержать только:
#### от 8 латинских символов
#### наличие цифр
#### наличие специальных символов "!,@,#,$,%"
      
**Ответ:**
- При успешной регистрации возвращается статус 201 с сообщением:
  ```json
  {
    "code": 200,
    "message": "Пользователь успешно зарегистрирован"
  }
  ```
- При ошибке возвращается ошибка с описанием проблемы.

### **POST** `/sign_in`
Авторизация пользователя и получение JWT токена.

**Запрос:**
- Тело запроса должно содержать JSON с полями:
   ```
   {
     "login": "John1234",
     "password": "12312john!"
   }
  ```

**Ответ:**
- При успешной авторизации возвращается токен:
  ```json
  {
    "code": 200,
    "message": "Вход выполнен успешно",
    "token": {
        "access_token": "JWT access token",
        "refresh_token": "JWT refresh token"
    }
  }
  ```
- При ошибке возвращается ошибка с описанием проблемы.

### **POST** `/todos`
Создание новой задачи.

**Запрос:**
- Требуется заголовок `Authorization` с токеном в формате `Bearer <JWT_TOKEN>`.
- Тело запроса должно содержать JSON с полями:
    - `title` (строка) — заголовок задачи.
    - `description` (строка) — описание задачи.

**Ответ:**
- При успешном создании задачи возвращается статус 201 с информацией о задаче:
  ```json
  {
    "id": 1,
    "title": "Sample Task",
    "description": "Task description",
    "user_id": 1
  }
  ```

### **GET** `/todos`
Получение списка задач пользователя.

**Запрос:**
- Требуется заголовок `Authorization` с токеном в формате `Bearer <JWT_TOKEN>`.

**Ответ:**
- При успешном запросе возвращается список задач:
  ```json
  [
    {
      "id": 1,
      "title": "Sample Task",
      "description": "Task description",
      "user_id": 1
    },
    ...
  ]
  ```

### **PUT** `/todos/:id`
Обновление задачи по её ID.

**Запрос:**
- Требуется заголовок `Authorization` с токеном в формате `Bearer <JWT_TOKEN>`.
- Параметр URL `id` — ID задачи, которую необходимо обновить.
- Тело запроса должно содержать JSON с полями:
    - `title` (строка) — новый заголовок задачи.
    - `description` (строка) — новое описание задачи.

**Ответ:**
- При успешном обновлении задачи возвращается статус 200 с сообщением:
  ```json
  {
    "message": "Task updated successfully"
  }
  ```

### **DELETE** `/todos/:id`
Удаление задачи по её ID.

**Запрос:**
- Требуется заголовок `Authorization` с токеном в формате `Bearer <JWT_TOKEN>`.
- Параметр URL `id` — ID задачи, которую необходимо удалить.

**Ответ:**
- При успешном удалении задачи возвращается статус 200 с сообщением:
  ```json
  {
    "message": "Task deleted successfully"
  }
  ```

---

## Авторизация и токены

Для взаимодействия с защищёнными маршрутам (например, создание, получение, обновление или удаление задач) необходимо использовать JWT токен, получаемый при авторизации пользователя.

Токен передаётся в заголовке `Authorization` в формате:

```
Authorization: Bearer <JWT_TOKEN>
```

### Пример получения токена

1. Авторизуйтесь с помощью **POST** запроса на `/sign_in`, передав логин и пароль.
2. В ответ вы получите JWT токен.
3. Используйте этот токен для доступа к защищённым маршрутам.

---

## Структура данных

### Пользователь (User)

При регистрации пользователя сохраняются следующие данные:
- `id` (int) — уникальный идентификатор пользователя.
- `user_name` (string) — имя пользователя.
- `login` (string) - логин пользователя.
- `password` (string) — хэш пароля.
- `role` (string) - роль пользователя.

### Задача (Todo)

Каждая задача имеет следующие поля:
- `id` (int) — уникальный идентификатор задачи.
- `title` (string) — заголовок задачи.
- `description` (string) — описание задачи.
- `user_id` (int) — ID пользователя, которому принадлежит задача.
- `status` (bool) - статус задачи.
- `created_at` (timestamp) - время создания задачи.

---

## Пример использования с Postman

1. **Регистрация**:
    - Метод: `POST`
    - URL: `http://localhost:2211/sign_up`
    - Тело запроса:
      ```json
      {
       "user_name": "John",
       "login": "John1234",
       "password": "12312john!"
      }
      ```

2. **Авторизация**:
    - Метод: `POST`
    - URL: `http://localhost:2211/sign_in`
    - Тело запроса:
      ```json
        {
          "login": "John1234",
          "password": "12312john!"
        }
      ```

3. **Создание задачи**:
    - Метод: `POST`
    - URL: `http://localhost:2211/todos`
    - Заголовок:
        - `Authorization: Bearer <JWT_TOKEN>`
    - Тело запроса:
      ```json
      {
         "title": "test",
         "description": "test",
         "done": false
      }
      ```

4. **Получение задач**:
    - Метод: `GET`
    - URL: `http://localhost:2211/todos`
    - Заголовок:
        - `Authorization: Bearer <JWT_TOKEN>`

5. **Обновление задачи**:
    - Метод: `PUT`
    - URL: `http://localhost:2211/todos/1`
    - Заголовок:
        - `Authorization: Bearer <JWT_TOKEN>`
    - Тело запроса:
      ```json
      {
         "title": "test3",
         "description": "test4",
         "done": true
      }
      ```

6. **Удаление задачи**:
    - Метод: `DELETE`
    - URL: `http://localhost:2211/todos/1`
    - Заголовок:
        - `Authorization: Bearer <JWT_TOKEN>`

---

## Ошибки

- **401 Unauthorized** — Нет авторизации или токен невалиден.
- **403 Forbidden** — У пользователя нет прав для выполнения действия.
- **404 Not Found** — Запрашиваемый ресурс не найден (например, задача с таким ID).
- **500 Internal Server Error** — Ошибка сервера.
