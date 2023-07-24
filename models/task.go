package models

import (
 "context"
 "github.com/jackc/pgx/v4/pgxpool"
)

type Task struct {
 ID       int    `json:"id"`
 Title    string `json:"title"`
 Completed bool  `json:"completed"`
}

type TaskModel struct {
 db *pgxpool.Pool
}

func NewTaskModel(db *pgxpool.Pool) *TaskModel {
 return &TaskModel{db: db}
}

func (m *TaskModel) GetAllTasks(ctx context.Context) ([]Task, error) {
 rows, err := m.db.Query(ctx, "SELECT * FROM tasks")
 if err != nil {
  return nil, err
 }
 defer rows.Close()

 tasks := make([]Task, 0)
 for rows.Next() {
  var task Task
  err := rows.Scan(&task.ID, &task.Title, &task.Completed)
  if err != nil {
   return nil, err
  }
  tasks = append(tasks, task)
 }

 return tasks, nil
}

func (m *TaskModel) GetTaskByID(ctx context.Context, id int) (*Task, error) {
 var task Task
 err := m.db.QueryRow(ctx, "SELECT * FROM tasks WHERE id = $1", id).Scan(&task.ID, &task.Title, &task.Completed)
 if err != nil {
  return nil, err
 }

 return &task, nil
}

func (m *TaskModel) CreateTask(ctx context.Context, task *Task) (int, error) {
 var id int
 err := m.db.QueryRow(ctx, "INSERT INTO tasks (title, completed) VALUES ($1, $2) RETURNING id", task.Title, task.Completed).Scan(&id)
 if err != nil {
  return 0, err
 }

 return id, nil
}

func (m *TaskModel) UpdateTask(ctx context.Context, task *Task) error {
 _, err := m.db.Exec(ctx, "UPDATE tasks SET title = $1, completed = $2 WHERE id = $3", task.Title, task.Completed, task.ID)
 if err != nil {
  return err
 }

 return nil
}

func (m *TaskModel) DeleteTask(ctx context.Context, id int) error {
 _, err := m.db.Exec(ctx, "DELETE FROM tasks WHERE id = $1", id)
 if err != nil {
  return err
 }

 return nil
}
