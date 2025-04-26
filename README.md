# Тестовое задание для Effective Mobile 𝐄𝐌

REST API для управления пользователями, написанный на Go, PostgreSQL и автоматической генерацией Swagger-документации.

## 📦 Возможности

- Создание нового пользователя
- Получение пользователей по фильтрам
- Обновление данных пользователя
- Удаление пользователя
- Swagger-документация

## 🚀 Запуск с Docker Compose

### 1. Создайте `.env` файл:

```env
PORT=your_port
DB_USER=your_user
DB_PASSWORD=your_password
DB_HOST=your_host
DB_PORT=your_db_port
DB_NAME=your_db_name
DB_SSLMODE=disable
```

2. Собери и запусти контейнеры:

```bash
docker-compose up --build
```

Сервер будет доступен по адресу:
📍 http://localhost:8080

Swagger UI:
📚 http://localhost:8080/swagger/index.html

### 🛠 Примеры запросов

`POST /person`

```json
{
	"name": "Ivan",
	"surname": "Ivanov",
	"patronymic": "Ivanovich",
	"age": 30,
	"gender": "male",
	"nationality": "RU"
}
```

`GET /person?age=30&gender=male` - фильтрация пользователей по параметрам.

`PUT /person` - обновление существующего пользователя.

`DELETE /person?id=1` - удаление пользователя по ID.

## 🧑‍💻 Автор

Дамир Усетов [ damirdgj@gmail.com ]
