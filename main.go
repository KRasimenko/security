package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"

	httpTransport "example.com/employee-service/http"
	"example.com/employee-service/implementation"
	"example.com/employee-service/repository"
	"example.com/employee-service/transport"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stdout)

	// Создаём хранилище
	repo := repository.NewMemoryRepository()

	// Создаём бизнес-логику
	svc := implementation.NewService(repo, logger)

	// Создаём endpoints
	endpoints := transport.MakeServerEndpoints(svc)

	// Создаём HTTP-обработчик
	handler := httpTransport.MakeHTTPHandler(endpoints)

	logger.Log("msg", "HTTP server started", "port", 8080)

	// Запускаем сервер
	if err := http.ListenAndServe(":8080", handler); err != nil {
		logger.Log("error", err)
		os.Exit(1)
	}
}
