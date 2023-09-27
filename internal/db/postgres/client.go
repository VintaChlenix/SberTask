package postgres

import (
	"SberTask/internal/db"
	"SberTask/internal/model"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"time"
)

type Client struct {
	db *pgx.Conn
}

func NewClient(connectionString string) (*Client, error) {
	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}
	return &Client{db: conn}, nil
}

var _ db.DB = Client{}

func (c Client) CreateTask(ctx context.Context, task model.Task) (int, error) {
	q := `
		INSERT INTO
		  tasks(header, description, date, is_done)
		VALUES
		  ($1, $2, $3, $4)
		RETURNING
		  task_id
	`
	var taskID int
	if err := c.db.QueryRow(ctx, q, task.Header, task.Description, task.Date, false).Scan(&taskID); err != nil {
		return 0, fmt.Errorf("failed to create task: %w", err)
	}

	return taskID, nil
}

func (c Client) SelectTask(ctx context.Context, taskID int) (*model.Task, error) {
	q := `
		SELECT
		  header, description, date, is_done
		FROM
		  tasks
		WHERE
		  task_id = $1
	`
	var task model.Task
	if err := c.db.QueryRow(ctx, q, taskID).Scan(&task.Header, &task.Description, &task.Date, &task.IsDone); err != nil {
		return nil, fmt.Errorf("failed to parse task: %w", err)
	}
	return &task, nil
}

func (c Client) UpdateTask(ctx context.Context, taskID int, task model.Task) error {
	tx, err := c.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()
	q := `
		SELECT
		  task_id
		FROM
		  tasks
		WHERE
		  task_id = $1
	`
	if row := c.db.QueryRow(ctx, q, taskID); row.Scan() == pgx.ErrNoRows {
		return fmt.Errorf("failed to select task ID: %w", pgx.ErrNoRows)
	}
	q = `
		UPDATE
		  tasks
		SET
		  header = COALESCE($2, header),
		  description = COALESCE($3, description),
		  date = COALESCE($4, date),
		  is_done = COALESCE($5, is_done)
		WHERE
		  task_id = $1
	`
	if _, err := c.db.Exec(ctx, q, taskID, task.Header, task.Description, task.Date, task.IsDone); err != nil {
		return fmt.Errorf("failed to update task by id: %w", err)
	}
	return nil
}

func (c Client) DeleteTask(ctx context.Context, taskID int) error {
	tx, err := c.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()
	q := `
		SELECT
		  task_id
		FROM
		  tasks
		WHERE
		  task_id = $1
	`
	if row := c.db.QueryRow(ctx, q, taskID); row.Scan() == pgx.ErrNoRows {
		return fmt.Errorf("failed to select task ID: %w", pgx.ErrNoRows)
	}
	q = `
		DELETE FROM
		  tasks
		WHERE
		  task_id = $1
	`
	if _, err := c.db.Exec(ctx, q, taskID); err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}
	return nil
}

func (c Client) SelectTasks(ctx context.Context, pagination model.Pagination) ([]model.Task, error) {
	q := `
		SELECT
		  header, description, date, is_done
		FROM
		  tasks
		WHERE
		  date BETWEEN $1 AND $2 AND is_done = ANY($3)
		ORDER BY date LIMIT $4 OFFSET $5
	`
	var rows pgx.Rows
	var err error
	if pagination.StartTimestamp == nil {
		pagination.StartTimestamp = new(time.Time)
		*pagination.StartTimestamp = time.Unix(0, 0)
	}
	if pagination.EndTimestamp == nil {
		pagination.EndTimestamp = new(time.Time)
		*pagination.EndTimestamp = time.Unix(1<<56, 0)
	}
	if pagination.IsDone != nil {
		rows, err = c.db.Query(ctx, q, pagination.StartTimestamp, pagination.EndTimestamp, []bool{*pagination.IsDone}, pagination.Rows, (*pagination.Page-1)*(*pagination.Rows))
	} else {
		rows, err = c.db.Query(ctx, q, pagination.StartTimestamp, pagination.EndTimestamp, []bool{true, false}, pagination.Rows, (*pagination.Page-1)*(*pagination.Rows))
	}
	if err != nil {
		return nil, fmt.Errorf("failed to select tasks: %w", err)
	}
	defer rows.Close()

	tasks := make([]model.Task, 0)
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.Header, &task.Description, &task.Date, &task.IsDone); err != nil {
			return nil, fmt.Errorf("failed to parse task: %w", err)
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to parse tasks: %w", err)
	}

	return tasks, nil
}
