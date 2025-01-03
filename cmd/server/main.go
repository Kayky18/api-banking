package main

import (
	"kayky18/api-banking/configs"
	"kayky18/api-banking/internal/database"
	"kayky18/api-banking/internal/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := configs.LoadConfig(".")

	if err != nil {
		log.Fatal(err)
	}

	db, err := configs.InitDataBase()

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	user := database.NewUserDB(db)
	UserDB := handlers.NewUserHandler(user, config.TokenAuth, config.JwtExpireIn)

	transaction := database.NewTransactionDB(db)
	TransactionDB := handlers.NewTransactionHandler(transaction)

	jwtHandler := handlers.NewJWTHandler(config.TokenAuth, config.JwtExpireIn)

	basePath := "api/v1/"

	r.Use(func(c *gin.Context) {
		c.Set("jwt", config.TokenAuth)
		c.Set("JwtExperiesIn", config.JwtExpireIn)
		c.Next()
	})

	v1 := r.Group(basePath)
	{
		//User
		usergin := v1.Group("user/")
		{
			usergin.POST("/", UserDB.Create)
			usergin.POST("/generate-jwt", UserDB.GetJwt)
			// usergin.GET("/list", UserDB.GetUsers)
			// usergin.GET("/:id", UserDB.GetUser)
			// usergin.PUT("/:id", UserDB.UpdateUser)
			// usergin.DELETE("/:id", UserDB.DeleteUser)

		}

		//Transaction
		transactiongin := v1.Group("transaction/")
		{

			transactiongin.POST("/", jwtHandler.JWTAuthMiddleware(), TransactionDB.CreateTransaction)
			transactiongin.GET("/", jwtHandler.JWTAuthMiddleware(), TransactionDB.GetTransaction)
		}

	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
