package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/sahilsh-dev/go-api/internal/app"
	"github.com/sahilsh-dev/go-api/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "go backend port")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
	defer app.DB.Close()
	app.Logger.Println("Created Application")

	r := routes.SetupRoutes(app)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	app.Logger.Println("Starting server started on PORT", port)
	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
