CREATE TABLE tasks
(
    task_id serial PRIMARY KEY,
    header text,
    description text,
    date timestamp,
    is_done boolean
)