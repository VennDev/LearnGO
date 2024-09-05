package app

import (
	"fmt"
	"net/http"

	"github.com/venndev/test-space/internal/config"
)

type App struct {
	config *config.Config
}

func New(cfg *config.Config) *App {
	return &App{config: cfg}
}

func (a *App) HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Chào mừng đến với ứng dụng Go!")
}
