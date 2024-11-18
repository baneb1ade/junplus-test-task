# Javacode Test Task(Golang version >= 1.23.0)

Это приложение реализует сервис управления кошельками с использованием REST API.

## Описание

### API

- **POST api/v1/wallet**

  Создает или обновляет баланс кошелька на основе типа операции (DEPOSIT или WITHDRAW) и указанной суммы.

  Пример запроса:
  ```json
  {
    "walletId": "UUID",
    "operationType": "DEPOSIT | WITHDRAW",
    "amount": 1000
  }
  ```
  Пример ответа:
  ```json
  {
	"uuid": "cc7a9d85-f728-4c44-b55b-34e354f5937a",
	"balance": 15500.5
  }
  ```

- **GET api/v1/wallets/{WALLET_UUID}**

    Возвращает текущий баланс кошелька.

    Пример ответа:
  ```json
  {
	"uuid": "a3c8a350-5b69-4d75-a16e-8d5bfa2b7a29",
	"balance": 150.75
  }
  ```
- **GET api/v1/wallets**
    
    Возвращает все кошельки.
    
    Пример ответа:
  ```json
    [
        {
        "uuid": "a3c8a350-5b69-4d75-a16e-8d5bfa2b7a29",
        "balance": 150.75
        },
        {
        "uuid": "bbd9c3f1-8a5f-4f3e-87e6-9c8b4a9d69c0",
        "balance": 2000
        },
        {
        "uuid": "cc7a9d85-f728-4c44-b55b-34e354f5937a",
        "balance": 500.5
        },
        {
        "uuid": "dde3f8e2-91a7-47fc-b09e-4f52934912a8",
        "balance": 750.25
        }
    ]
  ```

Стек технологий

    Golang – основной ЯП.
    Goose - инструмент миграций.
    Postgresql – база данных для хранения информации о кошельках и операциях.
    Docker – для контейнеризации приложения и базы данных.
    Docker Compose – для автоматического поднятия всей системы (приложения и базы данных).

### Запуск

Клонируйте репозиторий:

```bash
git clone https://github.com/baneb1ade/javacode-test-task.git
```
Перейдите в директорию:
```bash
cd javacode-test-task
```
При необходимости измените config.env и параметры подключения к БД в docker/docker-compose.yml.
Запустите deploy.sh
```bash
./deploy.sh
```
Скрипт поднимет 2 контейнера: app_container и postgres_container с приложением и БД соответственно
и накатит миграции на БД

Приложение будет доступно по адресу

    http://localhost:{PORT}

Запустить тесты командой
```bash
go test ./app/internal/wallet/tests -v
```