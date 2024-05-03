# test
## Тестовое задание: Средний уровень по GO/ClickHouse

**Срок выполнения**: 3 часа

### Задачи:

### Задача 1: SQL-запросы для ClickHouse

1. Выборка всех уникальных `eventType`, в которых более 1000 событий.
2. Выборка событий, которые происходят в первый день каждого месяца.
3. Выборка пользователей, которые реализуют более 3 различных типов событий.

#### Тестовая таблица для всех задач:

```sql
CREATE TABLE events (
    eventID Int64,
    eventType String,
    userID Int64,
    eventTime DateTime,
    payload String
) ENGINE = MergeTree
ORDER BY (eventID, eventTime);
```

### Задача 2: Реализация на GO

1. Вставка тестовых данных в таблицу событий.
2. Вывод событий по заданному типу события и временному диапазону.

### Задача 3: Реализация GO JSON API

Реализовать JSON API на GO с использованием метода обработки событий в ClickHouse.

**Пример запроса:**

```http
POST /api/event
{
  "eventType": "login",
  "userID": 1,
  "eventTime": "2023-04-09 13:00:00",
  "payload": "{\"some_field\":\"some_value\"}"
}
```
