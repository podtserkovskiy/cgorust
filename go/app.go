package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"siutsin.com/app/go/core"
)

func main() {
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		result := core.Add(5, 7)
		if _, err := fmt.Fprintf(w, "5 + 7 = %d (computed in Rust)\n", result); err != nil {
			slog.Error("Failed to write response", "error", err)
		}
	})

	slog.Info("Server listening", "port", 8080)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("Server failed", "error", err)
		os.Exit(1)
	}
}
