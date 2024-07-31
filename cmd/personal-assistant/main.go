package main

import (
    "flag"
    "fmt"
    "sync"
    "time"

    "github.com/sebastianhevia/personal-assistant/internal/db"
    sqlc "github.com/sebastianhevia/personal-assistant/internal/db/sqlc"
    "github.com/sebastianhevia/personal-assistant/internal/tasks"
)

func main() {
    db.InitDB("tasks.db")

    // Define flags
    action := flag.String("action", "list", "Action to perform (add/list)")
    name := flag.String("name", "", "Name of the task")
    status := flag.String("status", "to-do", "Status of the task (to-do/not for now/completed)")
    dueDate := flag.String("due", "", "Due date of the task (YYYY-MM-DD)")
    priority := flag.Int("priority", 0, "Priority of the task")
    flag.Parse()

    var wg sync.WaitGroup
    ch := make(chan sqlc.Task)

    switch *action {
    case "add":
        // Parse due date
        var due time.Time
        var err error
        if *dueDate != "" {
            due, err = time.Parse("2006-01-02", *dueDate)
            if err != nil {
                fmt.Println("Invalid due date format. Use YYYY-MM-DD.")
                return
            }
        } else {
            due = time.Now().Add(24 * time.Hour) // Default due date is tomorrow
        }

        // Create a new task
        wg.Add(1)
        go tasks.CreateTask(*name, *status, due, *priority, &wg)

    case "list":
        // Retrieve tasks
        wg.Add(1)
        go tasks.GetTasks("", ch, &wg)

        go func() {
            wg.Wait()
            close(ch)
        }()

        // Print tasks
        for task := range ch {
            fmt.Printf("%d: %s (Status: %s, Due: %s, Priority: %d)\n", task.ID, task.Name, task.Status, task.DueDate.Format("2006-01-02"), task.Priority)
        }

    default:
        fmt.Println("Unknown action. Use 'add' or 'list'.")
    }

    wg.Wait()
}

