# Common Lib 

Библиотека общих компонентов для микросервисной архитектуры проекта **Mountain Tour**. Пакет содержит унифицированные инструменты для работы с конфигурацией, базами данных, логированием и безопасностью, обеспечивая единство кода во всех сервисах.

## Установка

Для использования библиотеки в своем микросервисе добавьте её через `go get`:

```bash
go get [github.com/yousefggg/common-lib](https://github.com/yousefggg/common-lib)

Основные модули
1. Config (pkg/config)
Обеспечивает строго типизированную загрузку настроек из переменных окружения (ENV). Поддерживает валидацию обязательных полей и автоматическую конвертацию типов данных.

2. Logger (pkg/logger)
Высокопроизводительный логгер на базе uber-go/zap. Настроен для вывода структурированных JSON-логов, что идеально подходит для систем сбора и анализа логов (ELK Stack, Grafana Loki).

Доступные уровни: debug, info, warn, error

3. Errors (pkg/errors)
Кастомная реализация интерфейса error, позволяющая прокидывать строковые коды (Code) и оборачивать исходные ошибки для сохранения контекста и упрощения отладки.

4. JWT (pkg/jwt)
Модуль для работы с JSON Web Tokens (алгоритм шифрования HS256). Интегрирован с пакетом uuid для надежной идентификации пользователей.

5. Database (pkg/database)
Инициализация пула соединений PostgreSQL (sql.DB). Включает в себя:

Поддержку PingContext для проверки соединения при старте.

Гибкую настройку лимитов подключений (MaxOpenConns, MaxIdleConns).

Переменные окружения (Environment Variables)
Для корректной работы модулей config и database необходимо задать следующие ENV-переменные:
Переменная,Тип,Описание
APP_PORT,string,Порт приложения
APP_ENVIRONMENT,string,Окружение (dev/prod)
APP_LOG_LEVEL,string,Уровень логов (debug/info/warn/error)
DATABASE_URL,string,DSN: postgres://user:pass@host:port/db?sslmode=disable
DATABASE_MAX_OPEN_CONNS,int,Максимальное кол-во открытых соединений
DATABASE_MAX_IDLE_CONNS,int,Максимальное кол-во соединений в простое
DATABASE_CONN_TIMEOUT,duration,"Таймаут подключения (например, 5s)"
AUTH_JWT_SECRET,string,Секретный ключ для подписи токенов
AUTH_TOKEN_TIME,duration,"Время жизни токена (например, 24h)"

Зависимости (Dependencies)
Согласно go.mod, библиотека использует:

Core: go 1.25.7

UUID: github.com/google/uuid v1.6.0

Postgres Driver: github.com/lib/pq v1.12.3

JWT: github.com/golang-jwt/jwt/v5 v5.3.1

Logging: go.uber.org/zap v1.28.0

Установка через go get github.com/yousefggg/common-lib