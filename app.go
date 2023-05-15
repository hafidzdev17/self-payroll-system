package main

import (
	"fmt"
	"log"
	"net/http"
	"self-payrol/config"
	"self-payrol/delivery"
	"self-payrol/repository"
	"self-payrol/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	server struct {
		httpServer *echo.Echo
		cfg        config.Config
	}

	Server interface {
		Run()
	}
)

func InitServer(cfg config.Config) Server {
	e := echo.New()
	e.HideBanner = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return &server{
		httpServer: e,
		cfg:        cfg,
	}
}

func (s *server) Run() {

	s.httpServer.GET("", func(e echo.Context) error {

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status":  "success",
			"message": "Hello, World!" + s.cfg.ServiceName() + " " + s.cfg.ServiceEnvironment(),
		})
	})

	positionRepo := repository.NewPositionRepository(s.cfg)
	positionUsecase := usecase.NewPositionUsecase(positionRepo)
	positionDelivery := delivery.NewPositionDelivery(positionUsecase)
	positionGroup := s.httpServer.Group("/positions")
	positionDelivery.Mount(positionGroup)

	companyRepo := repository.NewCompanyRepository(s.cfg)
	companyUsecase := usecase.NewCompanyUsecase(companyRepo)
	companyDelivery := delivery.NewCompanyDelivery(companyUsecase)
	companyGroup := s.httpServer.Group("/company")
	companyDelivery.Mount(companyGroup)

	// TODO: panggil user repository, user usecase, user derlivery, dan mount ke router
	userRepo := repository.NewUserRepository(s.cfg)
	userUsecase := usecase.NewUserUsecase(userRepo, positionRepo, companyRepo)
	userDelivery := delivery.NewUserDelivery(userUsecase)
	userGroup := s.httpServer.Group("/employee")
	userDelivery.Mount(userGroup)

	transactionRepo := repository.NewTransactionRepository(s.cfg)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepo)
	transactionDelivery := delivery.NewTransactionDelivery(transactionUsecase)
	transactionGroup := s.httpServer.Group("/transactions")
	transactionDelivery.Mount(transactionGroup)

	if err := s.httpServer.Start(fmt.Sprintf(":%d", s.cfg.ServicePort())); err != nil {
		log.Panic(err)
	}
}
