package bootstrap

import (
	"fmt"
	"net"
	"os"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"vladislove-gRPC/internal/services"
)

type User struct {
	ID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name  string
	Email string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type App struct {
	logger *log.Logger
	cfg    *Config
	db     *gorm.DB
}

func NewApp(config *Config) *App {
	return &App{
		cfg: config,
	}
}

func (a *App) Run() error {
	a.initLogger()

	a.logger.Info("Инициализация Базы Данных...")
	if err := a.initDB(); err != nil {
		return fmt.Errorf("init db connection: %w", err)
	}

	a.logger.Info("Инициализация gRPC	...")
	if err := a.initGRPC(); err != nil {
		return fmt.Errorf("init gRPC error: %w", err)
	}

	return nil
}

func (a *App) initLogger() {
	var logger = &log.Logger{
		Out: os.Stdout,
		Formatter: &nested.Formatter{
			HideKeys:        true,
			TimestampFormat: "2006-01-02 15:04:05",
			FieldsOrder:     []string{"component", "category"},
			NoFieldsSpace:   true,
			NoColors:        false,
			TrimMessages:    true,
		},
		Hooks: make(log.LevelHooks),
		Level: log.DebugLevel,
	}

	a.logger = logger
}

func (a *App) initDB() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", a.cfg.DBHost, a.cfg.DBUser, a.cfg.DBPass,
		a.cfg.DBName, a.cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	if err := db.AutoMigrate(&User{}); err != nil {
		return err
	}

	a.db = db

	return nil
}

func (a *App) initGRPC() error {
	listener, listenErr := net.Listen("tcp", ":50051")
	if listenErr != nil {
		a.logger.Fatalf("Ошибка запуска сервера: %v", listenErr)
	}

	grpcServer := grpc.NewServer()

	services.RegisterServices(grpcServer, a.logger)

	a.logger.Info("gRPC сервер запущен на порту 50051")
	if err := grpcServer.Serve(listener); err != nil {
		a.logger.Fatalf("Ошибка запуска сервера: %v", err)
	}

	return nil
}
