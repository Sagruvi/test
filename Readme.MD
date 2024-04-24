
Задание 2
Выборки всех уникальных eventType, в которых более 1000 событий.
SELECT eventType
FROM events
GROUP BY eventType
HAVING count(*) > 1000;

Выборки событий, которые происходят в первый день каждого месяца.
SELECT *
FROM events
WHERE toStartOfMonth(eventTime) = eventTime;

Выборки пользователей, которые реализуют более 3 различных типов событий.
SELECT userID
FROM (
SELECT userID, count(DISTINCT eventType) AS numEventTypes
FROM events
GROUP BY userID
HAVING numEventTypes > 3
);
