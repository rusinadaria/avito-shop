## Инструкция по запуску

1. Склонируйте репозиторий
2. Перейдите в корневую папку проекта
3. Соберите образы с помощью команды `docker-compose build`
4. Запустите контейнеры командой `docker-compose up`
5. Все готово! 
6. Для запуска юнит-тестов перейдите в директорию ./internal/handlers и воспользуйтесь  командой `go test -v`, для запуска интеграционных тестов используйте эту команду в директории ./tests

env:\
PORT=:8080\
DB_PATH=user=postgres password=root dbname=shop sslmode=disable\
DATABASE_PORT=5432\
DATABASE_USER=postgres\
DATABASE_PASSWORD=root\
DATABASE_NAME=shop\
DATABASE_HOST=db\
SERVER_PORT=8080