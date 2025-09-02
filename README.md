# ToDo Application

Сервис управления задачами.

## Клонирование и запуск

```bash
git clone https://github.com/scwern/ToDo.git
cd todo
docker-compose up -d
```

Доступ после запуска:

* **API:** [http://localhost:8080](http://localhost:8080)
* **Swagger Docs:** [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
* **База данных:** 5432 (PostgreSQL)

## Основные возможности API

* CRUD для задач и пользователей
* История задач пользователя
* Буфферизированные каналы для удаления и обработки задач
* **Gzip компрессия:**

  * Принимает сжатые запросы (Content-Encoding)
  * Отдаёт сжатые ответы клиенту (Accept-Encoding)
  * Работает для application/json и text/html
* Graceful Shutdown сервера
* Поддержка in-memory и PostgreSQL хранилищ
* Миграции базы данных (migrations)

## Конфигурация

* Настройки через JSON файл (`internal/config`) и переменные окружения

## Инструменты и качество кода

* Линтер: golangci-lint
* Бенчмарки для основных операций
