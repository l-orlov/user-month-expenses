# user-month-expenses
API для получения предсказаний трат пользователей по различным категориям

## Оглавление
- [Конфигурация](#configuration)
- [Развертывание](#deployment)

<a name="configuration"></a>
## Конфигурация
Конфигурация происходит следующим образом:
1. Читается конфиг по пути, указанному в переменной окружения `CONFIG_PATH`.
   Файл должен иметь формат YAML и иметь определенную структуру.
2. Читаются оставшиеся настройки из переменных окружения.
   При дублировании настроек переменные окружения затирают параметры конфига.

Список переменных окружения:
```
CONFIG_PATH=configs/config.yaml
LOGGER_LEVEL=debug
LOGGER_FORMAT=default
PG_ADDRESS=0.0.0.0:5432
PG_USER=user-month-expenses
PG_PASSWORD=123
PG_DATABASE=user-month-expenses
```

<a name="deployment"></a>
## Развертывание
1. Для того, чтобы развернуть сервис в docker:  
   ```docker-compose up -d --build```

   Опустить контейнеры:  
   ```docker-compose down```
2. Чтобы выполнить начальную миграцию для базы данных, нужно установить <a href="https://github.com/golang-migrate/migrate">эту утилиту</a> и выполнить команду:  
   ```migrate -path ./schema -database 'postgres://user-month-expenses:123@localhost:54320/user-month-expenses?sslmode=disable' up```

   Откатить миграцию:  
   ```migrate -path ./schema -database 'postgres://user-month-expenses:123@localhost:54320/user-month-expenses?sslmode=disable' down```  
