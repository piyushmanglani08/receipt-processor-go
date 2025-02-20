package main

import (
	"flag"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"receipt-processor-go/handlers"
)

func main() {
	var port int
	var useLogging bool
	flag.IntVar(&port, "port", 8080, "port used for the web server")
	flag.BoolVar(&useLogging, "logging", false, "enable detailed logging for all requests")
	flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/receipts/process", handlers.ProcessReceipt).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", handlers.GetPoints).Methods("GET")

	//middleware can be enabled using useLogging
	if useLogging {
		router.Use(loggingMiddleware)
	}

	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + flag.Lookup("port").Value.String(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server running on port %d...", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var headerDetails []string
		for key, values := range r.Header {
			headerDetails = append(headerDetails, key+": "+strings.Join(values, ","))
		}
		log.Printf("Request: %s %s, Headers: [%s]", r.Method, r.RequestURI, strings.Join(headerDetails, "; "))
		next.ServeHTTP(w, r)
	})
}
