package tasks

import (
    "context"
    "fmt"
    "sync"
    "time"

    "github.com/sebastianhevia/personal-assistant/internal/db"
    "github.com/sebastianhevia/personal-assistant/internal/db/sqlc"
)

func CreateTask(name, status string, dueDate time.Time, priority int, wg *sync.WaitGroup) {
    defer wg.Done()
	queries := sqlc.New(db.Database:w http.ResponseWriter, r *http.Request)
    err := queries.CreateTask(context.Background(), sqlc.CreateTaskParams{
        Name:     name,
        Status:   status,
        DueDate:  dueDate,
        Priority: priority,
    })
    if err != nil {
        fmt.Println("Error creating task:", err)
    }
}

func GetTasks(statusFilter string, ch chan<- sqlc.Task, wg *sync.WaitGroup) {
    defer wg.Done()
    queries := sqlc.New(db.Database)
    tasks, err := queries.GetTasks(context.Background(), statusFilter)
    if err != nil {
        fmt.Println("Error retrieving tasks:", err)
    }
    for _, task := range tasks {
        ch <- task
    }
    close(ch)
}

