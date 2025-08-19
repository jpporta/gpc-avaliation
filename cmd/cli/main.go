package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jpporta/gpc-avaliation/internal/services"
	"github.com/jpporta/gpc-avaliation/internal/utils"
)

type Application struct {
	Services *services.Service
}

func main() {
	service := services.New()
	app := &Application{
		Services: service,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if len(os.Args) < 2 {
		panic("command is required")
	}
	switch os.Args[1] {
	case "create":
		app.createCmd(ctx, os.Args[2:])
		break
	default:
		panic("unknown command")
	}
}

func (a Application) createCmd(ctx context.Context, args []string) {
	if len(args) < 1 {
		panic("create command is required")
	}
	switch args[0] {
	case "juror":
		name := args[1]
		if name == "" {
			panic("name is required")
		}
		id := utils.Must(a.Services.CreateJuror(ctx, name))
		fmt.Printf("User created with ID: %d\n", id)
		break
	case "category":
		name := args[1]
		if name == "" {
			panic("name is required")
		}
		description := args[2]
		if description == "" {
			panic("description is required")
		}
		id := utils.Must(a.Services.CreateCategory(ctx, name, description))
		fmt.Printf("Category created with ID: %d\n", id)
	case "model":
		name := args[1]
		if name == "" {
			panic("name is required")
		}
		author := args[2]
		if author == "" {
			panic("author is required")
		}
		description := args[3]
		if description == "" {
			panic("description is required")
		}
		categoryId := utils.Must(strconv.Atoi(args[4]))
		id := utils.Must(a.Services.CreateModel(ctx, name, author,
			description, int64(categoryId)))
		fmt.Printf("Model created with ID: %d\n", id)
	default:
		panic("unknown create command")
	}
}
