package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello, World!")

	d, err := io.ReadAll(r.Body)
	if err != nil {
		h.l.Println("Error reading request body:", err)

		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	h.l.Printf("Data %s\n", d)

	fmt.Fprintf(w, "Hello, %s!", d)
}
