# Subscription Management API

Этот REST сервис предназначен для агрегации и управления данными об онлайн-подписках пользователей.

### Установка и Запуск

1.  **Клонируйте репозиторий:**
    ```bash
    git clone https://github.com/fresh132/REST-API-agregating.git 
    cd REST-API-agregating 
    ```

2.  **Настройте переменные окружения:**
    Создайте файл `.env` в корне проекта (рядом с `docker-compose.yaml`) со следующим содержимым:
    ```env
    POSTGRES_USER=user
    POSTGRES_PASSWORD=password
    POSTGRES_DB=postgres
    POSTGRES_PORT=5432

    DATABASE_URL=postgres://user:password@postgres-db:5432/postgres?sslmode=disable

    PORT=9091 # Или 8080, убедитесь, что порт свободен

    GIN_MODE=debug
    ```

3.  **Установите Go зависимости:**
    ```bash
    go mod tidy
    ```

4.  **Сгенерируйте Swagger документацию:**
    ```bash
    go install github.com/swaggo/swag/cmd/swag@latest
    ```

5.  **Запустите сервисы с Docker Compose:**
    ```bash
    docker-compose up --build -d
    ```
    *Если получите ошибку "container name already in use", выполните `docker-compose down --remove-orphans` и повторите запуск.*

### 🔍 Проверка

-   **Health Check:** `curl http://localhost:8080/health`
    (Замените `8080` на порт из вашего `.env`, если он отличается)
-   **Swagger UI:** Откройте в браузере `http://localhost:8080/swagger/index.html`
    (Замените `8080` на порт из вашего `.env`, если он отличается)

## API Эндпоинты

Swagger UI предоставляет интерактивную документацию и возможность тестирования всех доступных эндпоинтов.