Цитатник — QuotationBook.
REST API на Go для хранения и управления цитатами. Данные хранятся в памяти.  
Реализован с использованием только стандартной библиотеки Go + `gorilla/mux`.

Установка и запуск:

    Шаг 1: Клонировать репозиторий
          git clone https://github.com/himka0-0/QuotationBook.git
          cd QuotationBook
    Шаг 2 (вариант 1): Запуск без Docker
          go run main.go
    Шаг 2 (вариант 2): Запуск через Docker
          # Сборка
            docker build -t quotationbook .
          # Запуск
            docker run -p 8080:8080 quotationbook
Тесты:

    go test ./models
    go test ./handler

Примеры запросов:

```bash
# Добавить цитату
curl -X POST http://localhost:8080/quotes \
  -H "Content-Type: application/json" \
  -d '{"author":"Конфуций", "quote":"Жизнь проста, но мы её усложняем."}'

# Получить все цитаты
curl http://localhost:8080/quotes

# Фильтр по автору
curl http://localhost:8080/quotes?author=Конфуций

# Случайная цитата
curl http://localhost:8080/quotes/random

# Удалить цитату
curl -X DELETE http://localhost:8080/quotes/1
```

Особенности:

Все данные хранятся в памяти (по ТЗ).
Используется только net/http и gorilla/mux.
