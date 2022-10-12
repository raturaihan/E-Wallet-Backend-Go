package router

import (
	"assignment-golang-backend/database"
	"assignment-golang-backend/handler"
	"assignment-golang-backend/repository"
	"assignment-golang-backend/usecase"

	"github.com/gin-gonic/gin"
)

func Router() {
	db := database.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	transRepository := repository.NewTransactionRepository(db)
	transUsecase := usecase.NewTransactionUsecase(transRepository, userRepository)
	transHandler := handler.NewTransactionHandler(transUsecase)

	r := gin.Default()
	authRoute := r.Group("/")
	{
		authRoute.POST("/login", userHandler.Login)
		authRoute.POST("/register", userHandler.Register)
	}

	transRoute := r.Group("/transaction")
	{
		transRoute.POST("/topup", transHandler.TopUpAmount)
	}
	r.Run(":8080")
}
