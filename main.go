package main

import (
    "context"
    "fmt"
    "log"

    "github.com/prisma/prisma-client-go/runtime/builder"
    "github.com/steebchen/prisma-client-go"
)

func main() {
    client := prisma.NewClient()
    if err := client.Connect(); err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer func() {
        if err := client.Disconnect(); err != nil {
            log.Fatalf("Failed to disconnect from database: %v", err)
        }
    }()

    ctx := context.Background()

    // Create a new user
    newUser, err := client.User.CreateOne(
        prisma.User.Name.Set("Ahmad"),
        prisma.User.Email.Set("ahmad@example.com"),
    ).Exec(ctx)
    if err != nil {
        log.Fatalf("Failed to create user: %v", err)
    }

    fmt.Printf("Created user: %v\n", newUser)

    // Fetch all users
    users, err := client.User.FindMany().Exec(ctx)
    if err != nil {
        log.Fatalf("Failed to fetch users: %v", err)
    }

    fmt.Printf("Users: %v\n", users)
}