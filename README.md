# SberTask

## Инструкция по запуску
1) Скачать репозиторий
2) docker-compose build
3) docker-compose up
4) В папке /internal/db/migrations лежит .sql файл с SQL запросами для инициализации базы данных. Нужно зайти в админер на localhost:8081 и выполнить их.

### 1) Метод создания задачи. Принимает заголовок, описание и дату.

### Запрос:
```
  curl --location 'localhost:8080/create_task' \
    --header 'Content-Type: application/json' \
    --data '{
        "header": "header1",
        "description": "desc1",
        "date": "2023-09-26T18:00:00Z"
    }'
```
### Ответ:
```
  {
    "task_id": 16
  }
```

### 2) Метод получения задачи. Принимает в URL ID задачи.

### Запрос:
 ```
   curl --location 'localhost:8080/get_task/16'
 ```
### Ответ:
```
  {
    "task": {
        "header": "header1",
        "description": "desc1",
        "date": "2023-09-26T18:00:00Z",
        "is_done": false
    }
  }
```
### 3) Метод обновления задачи. Принимает поля задачи(не обязательно все): заголовок, описание, дату или статус, а также в URL ID задачи.

### Запрос:
 ```
   curl --location 'localhost:8080/update_task/16' \
     --header 'Content-Type: application/json' \
     --data '{
       "is_done": true
     }'
 ```
### Ответ:
```
    200 OK
```

### 4) Метод удаления задачи по ID. Принимает в URL ID задачи.

### Запрос:
```
    curl --location --request DELETE 'localhost:8080/delete_task/15'
```
### Ответ:
```
    200 OK
```

### 5) Метод получения задач по критериям. Принимает статус, даты(от, до), номер страницы и количество строк на странице.

### Запрос:
```
    curl --location --request GET 'localhost:8080/get_tasks' \
    --header 'Content-Type: application/json' \
    --data '{
        "is_done": false,
        "start_timestamp": "2023-09-25T00:00:00Z",
        "end_timestamp": "2023-09-28T00:00:00Z",
        "page": 1,
        "rows": 10
    }'
```
### Ответ:
```
    {
        "tasks": [
            {
                "header": "test-test",
                "description": "try description",
                "date": "2023-09-26T18:00:00Z",
                "is_done": false
            },
            {
                "header": "test-test",
                "description": "try description",
                "date": "2023-09-26T18:00:00Z",
                "is_done": false
            },
            {
                "header": "test-test",
                "description": "try description",
                "date": "2023-09-26T18:00:00Z",
                "is_done": false
            }
        ]
    }
```

