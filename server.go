package ServiceAPI

import (
	"context"
	"net/http"
	"time"
)
/*
Абстракция на сервером HTTP.
Создаем основной метод для запуска сервера, в который входит метод ListenAndServe (для запуска HTTP сервера).
Для сервера устанавливаем основные параметры таймаутов, адреса, портов и обработчиков. В поле Handler записываются уже инициализированные обработчики из пакета обработчиков.
*/
type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
