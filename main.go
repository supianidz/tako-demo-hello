package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Response struct {
	Message   string    `json:"message"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Hostname  string    `json:"hostname"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	hostname, _ := os.Hostname()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		accept := r.Header.Get("Accept")
		if strings.Contains(accept, "text/html") {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>TAKO Demo Service</title>
  <style>
    * { box-sizing: border-box; margin: 0; padding: 0; font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif; }
    body { background: #09090b; color: #f4f4f5; display: flex; align-items: center; justify-content: center; min-height: 100vh; padding: 1.5rem; }
    .card { background: #18181b; border: 1px solid #27272a; border-radius: 12px; padding: 2.5rem; max-width: 480px; width: 100%; text-align: center; }
    .badge { display: inline-flex; align-items: center; gap: 6px; background: rgba(16, 185, 129, 0.1); border: 1px solid rgba(16, 185, 129, 0.25); color: #34d399; font-size: 0.75rem; font-weight: 600; padding: 4px 10px; border-radius: 9999px; margin-bottom: 1.5rem; }
    .dot { width: 6px; height: 6px; background: #10b981; border-radius: 50%; }
    h1 { font-size: 1.75rem; font-weight: 700; margin-bottom: 0.5rem; color: #ffffff; }
    p { font-size: 0.875rem; color: #a1a1aa; line-height: 1.6; margin-bottom: 1.5rem; }
    .info { background: #09090b; border: 1px solid #27272a; border-radius: 8px; padding: 1rem; font-family: monospace; font-size: 0.8125rem; color: #e4e4e7; text-align: left; }
    .info div { display: flex; justify-content: space-between; margin-bottom: 0.25rem; }
    .info div:last-child { margin-bottom: 0; }
    .info span:first-child { color: #71717a; }
  </style>
</head>
<body>
  <div class="card">
    <div class="badge"><span class="dot"></span> Online & Serving</div>
    <h1>Hello World from TAKO v3.3! 🐙</h1>
    <p>This service was deployed automatically using TAKO Control Plane.</p>
    <div class="info">
      <div><span>Status:</span> <span>ok</span></div>
    </div>
  </div>
</body>
</html>`, hostname, port)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(Response{
			Message:   "Hello World from TAKO! 🐙",
			Status:    "ok",
			Timestamp: time.Now().UTC(),
			Hostname:  hostname,
		})
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK\n")
	})

	log.Printf("🐙 Hello World server listening on :%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server exited with error: %v", err)
	}
}
