package main

import (
	"fmt"
	"context"
	"os"
	"os/signal"

	"github.com/okekefrancis112/orders-api/application"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app: ", err)
	}

}