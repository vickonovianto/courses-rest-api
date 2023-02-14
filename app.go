package main

import (
	"courses-api/config"
	"courses-api/delivery"
	"courses-api/model"
	"courses-api/repository"
	"courses-api/usecase"
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
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
	adminRepository := repository.NewAdminRepository(s.cfg)
	adminUsecase := usecase.NewAdminUsecase(adminRepository)
	adminDelivery := delivery.NewAdminDelivery(adminUsecase)
	adminGroup := s.httpServer.Group("/admin")
	adminDelivery.Mount(adminGroup)

	// Configure middleware with the custom claims type
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(model.CustomClaims)
		},
		SigningKey: []byte(os.Getenv("SECRET")),
	}

	courseCategoryRepository := repository.NewCourseCategoryRepository(s.cfg)
	courseCategoryUsecase := usecase.NewCourseCategoryUsecase(courseCategoryRepository)
	courseCategoryDelivery := delivery.NewCourseCategoryDelivery(courseCategoryUsecase)
	courseCategoryGroup := s.httpServer.Group("/course-categories")
	courseCategoryGroup.Use(echojwt.WithConfig(config))
	courseCategoryDelivery.Mount(courseCategoryGroup)

	courseRepository := repository.NewCourseRepository(s.cfg)
	courseUsecase := usecase.NewCourseUsecase(courseRepository, courseCategoryRepository)
	courseDelivery := delivery.NewCourseDelivery(courseUsecase)
	courseGroup := s.httpServer.Group("/courses")
	courseGroup.Use(echojwt.WithConfig(config))
	courseDelivery.Mount(courseGroup)

	userRepository := repository.NewUserRepository(s.cfg)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userDelivery := delivery.NewUserDelivery(userUsecase)
	userGroup := s.httpServer.Group("/users")
	userGroup.Use(echojwt.WithConfig(config))
	userDelivery.Mount(userGroup)

	userCourseRepository := repository.NewUserCourseRepository(s.cfg)
	userCourseUsecase := usecase.NewUserCourseUsecase(userCourseRepository, userRepository, courseRepository)
	userCourseDelivery := delivery.NewUserCourseDelivery(userCourseUsecase)
	userCourseGroup := s.httpServer.Group("/user-courses")
	userCourseGroup.Use(echojwt.WithConfig(config))
	userCourseDelivery.Mount(userCourseGroup)

	if err := s.httpServer.Start(fmt.Sprintf(":%d", s.cfg.ServicePort())); err != nil {
		log.Panic(err)
	}
}
