package router

import (
	"assignment-golang-backend/database"
	"assignment-golang-backend/handler"
	"assignment-golang-backend/middleware"
	"assignment-golang-backend/repository"
	"assignment-golang-backend/usecase"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	db := database.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	transRepository := repository.NewTransactionRepository(db)
	transUsecase := usecase.NewTransactionUsecase(transRepository, userRepository)
	transHandler := handler.NewTransactionHandler(transUsecase)

	authRoute := r.Group("/")
	{
		authRoute.POST("/login", userHandler.Login)
		authRoute.POST("/register", userHandler.Register)
	}

	transRoute := r.Group("/transaction")
	transRoute.Use(middleware.CheckAuth())
	{
		transRoute.POST("/topup", transHandler.TopUpAmount)
		transRoute.POST("/transfer", transHandler.Transfer)
	}

	r.Run(":8081")

}
