package server

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"balancing_algorithms_analysis/internal/server/balancer"
)

type Server struct {
	balancer       balancer.Balancer
	averageReqTime float64
	reqTimeArr     []float64
}

func New(balancer balancer.Balancer) Server {
	server := Server{
		balancer:       balancer,
		averageReqTime: 0,
		reqTimeArr:     []float64{},
	}
	return server
}

func (s Server) Run(ctx context.Context, httpPort string) {
	router := s.createRouter()
	httpServer := &http.Server{
		Addr:    "localhost:" + httpPort,
		Handler: router,
	}

	wg := sync.WaitGroup{}

	log.Printf("[httpServer] starting on %s\n", httpPort)

	wg.Add(1)
	go func() {
		httpServerRun(httpServer)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		gracefulShutdown(ctx, httpServer)
		wg.Done()
	}()

	wg.Wait()
}

func httpServerRun(httpServer *http.Server) {
	err := httpServer.ListenAndServe()
	if err != nil {
		log.Printf("[httpServer] ListenAndServe: %v\n", err)
	}
}

func gracefulShutdown(ctx context.Context, httpServer *http.Server) {
	<-ctx.Done()
	ctxTo, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("[gracefulShutdown] shutting down")

	err := httpServer.Shutdown(ctxTo)
	if err != nil {
		log.Printf("httpServer.Shutdown: %v\n", err)
	}

	log.Println("[gracefulShutdown] shut down successfully")
}
