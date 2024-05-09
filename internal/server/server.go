package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
	"xorm.io/xorm"
)

type Server struct {
	Router *mux.Router
	Engine *xorm.Engine
}

// StartServer creates a new server instance
func StartServer(engine *xorm.Engine) *Server {
	fmt.Println("Starting HTTP server...")
	router := mux.NewRouter()
	server := &Server{
		Router: router,
		Engine: engine,
	}
	ConfigureRoutes(router, engine)

	fmt.Println("Server started successfully!")
	return server
}

func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rateLimitRequests := os.Getenv("RATE_LIMIT_REQUESTS")
		rateLimitBurst := os.Getenv("RATE_LIMIT_BURST")

		requestsPerSecond, err := strconv.ParseFloat(rateLimitRequests, 64)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		burst, err := strconv.Atoi(rateLimitBurst)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		limiter := rate.NewLimiter(rate.Limit(requestsPerSecond), burst)

		if !limiter.Allow() {
			http.Error(w, "Limit Exceeded", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
