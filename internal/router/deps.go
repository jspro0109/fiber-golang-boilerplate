package router

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"fiber-golang-boilerplate/config"
	"fiber-golang-boilerplate/internal/handler"
	"fiber-golang-boilerplate/pkg/health"
)

type Deps struct {
	AuthHandler   *handler.AuthHandler
	UserHandler   *handler.UserHandler
	UploadHandler *handler.UploadHandler
	AdminHandler  *handler.AdminHandler
	Config        *config.Config
	Pool          *pgxpool.Pool
	Health        *health.Checker
}
