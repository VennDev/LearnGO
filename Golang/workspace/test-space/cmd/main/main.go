package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/venndev/test-space/internal/app"
	"github.com/venndev/test-space/internal/config"
)

func main() {
	// Khởi tạo cấu hình
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Không thể tải cấu hình: %v", err)
	}

	// Khởi tạo ứng dụng
	app := app.New(cfg)

	// Thiết lập routes
	http.HandleFunc("/", app.HomeHandler)

	// Khởi động server
	fmt.Printf("Server đang chạy tại http://localhost:%d\n", cfg.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil)
	if err != nil {
		log.Fatalf("Không thể khởi động server: %v", err)
	}
}
