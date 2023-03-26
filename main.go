package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	ent "lkuoch/ent-todo/ent/__generated__"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Create a user
	user, err := CreateUser(ctx, client)
	if err != nil {
		log.Fatal(err)
	}

	// Assign todos to user
	if err = CreateTodoForUser(ctx, client, user); err != nil {
		log.Fatal(err)
	}
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	user, err := client.
		User.
		Create().
		SetUsername("lkuoch").
		SetDisplayName("Law").
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	fmt.Println("✅ Added user:", user.DisplayName)

	return user, nil
}

func CreateTodoForUser(ctx context.Context, client *ent.Client, _user *ent.User) (_err error) {
	todo, err := client.
		Todo.
		Create().
		SetTitle("My first todo").
		Save(ctx)

	if err != nil {
		_err = errors.Join(err, fmt.Errorf("failed creating todo: %w", err))
	}

	user, err := client.
		User.
		UpdateOne(_user).
		AddTodos(todo).
		Save(ctx)

	if err != nil {
		_err = errors.Join(err, fmt.Errorf("failed assigning todo: %w", err))
	}

	fmt.Println("✅ Added todo:", todo.Title, "for user:", user.DisplayName)

	return
}
