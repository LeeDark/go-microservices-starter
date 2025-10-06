package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/LeeDark/go-microservices-starter/handlers"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	log.Println("Hello, World!")

	// 	d, err := io.ReadAll(r.Body)
	// 	if err != nil {
	// 		http.Error(w, "Failed to read request body", http.StatusBadRequest)
	// 		return
	// 	}

	// 	log.Printf("Data %s\n", d)

	// 	//w.Write([]byte("Hello, World!"))
	// 	fmt.Fprintf(w, "Hello %s!", d)
	// })

	// http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
	// 	log.Println("Goodbye, World!")
	// 	w.Write([]byte("Goodbye, World!"))
	// })

	l := log.New(os.Stdout, "miscroservice ", log.LstdFlags)

	// hh := handlers.NewHello(l)
	// gh := handlers.NewGoodbye(l)
	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()
	// sm.Handle("/", hh)
	// sm.Handle("/goodbye", gh)
	sm.Handle("/", ph)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// l.Println("Starting server on port 9090")
	// http.ListenAndServe(":9090", sm)

	go func() {
		l.Println("Starting server on port 9090")
		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()
	s.Shutdown(ctx)
}
