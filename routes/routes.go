package routes

import (
	"strings"
	"time"

	"github.com/adhyttungga/go-chatapp-service/config"
	auth_delivery "github.com/adhyttungga/go-chatapp-service/delivery/auth"
	message_delivery "github.com/adhyttungga/go-chatapp-service/delivery/message"
	user_delivery "github.com/adhyttungga/go-chatapp-service/delivery/user"
	"github.com/adhyttungga/go-chatapp-service/middleware"
	auth_repository "github.com/adhyttungga/go-chatapp-service/repository/auth"
	message_repository "github.com/adhyttungga/go-chatapp-service/repository/message"
	user_repository "github.com/adhyttungga/go-chatapp-service/repository/user"
	auth_usecase "github.com/adhyttungga/go-chatapp-service/usecase/auth"
	message_usecase "github.com/adhyttungga/go-chatapp-service/usecase/message"
	user_usecase "github.com/adhyttungga/go-chatapp-service/usecase/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(DB *mongo.Database, validate *validator.Validate) *gin.Engine {
	router := gin.Default()
	allowOrigins := strings.Split(config.Config.Origin.AllowOrigin, ",")

	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-length"},
		MaxAge:           12 * time.Hour,
		AllowCredentials: true,
	}))

	authRepository := auth_repository.NewAuthRepository(DB)
	authUsecase := auth_usecase.NewAuthUsecase(authRepository)
	authDelivery := auth_delivery.NewAuthDelivery(authUsecase, validate)

	messageRepository := message_repository.NewMessageRepository(DB)
	messageUsecase := message_usecase.NewMessageUsecase(messageRepository)
	messageDelivery := message_delivery.NewMessageDelivery(messageUsecase, validate)

	userRepository := user_repository.NewUserRepository(DB)
	userUsecase := user_usecase.NewUserUsecase(userRepository)
	userDelivery := user_delivery.NewUserDelivery(userUsecase)

	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/signup", authDelivery.Signup)
		authRoutes.POST("/login", authDelivery.Login)
		authRoutes.POST("/logout", authDelivery.Logout)
	}
	messageRoutes := router.Group("/api/message").Use(middleware.ProtectRoute())
	{
		messageRoutes.POST("/send/:id", messageDelivery.SendMessage)
		messageRoutes.GET("/:id", messageDelivery.GetMessages)
	}
	userRoutes := router.Group("/api/user").Use(middleware.ProtectRoute())
	{
		userRoutes.GET("/", userDelivery.GetUser)
	}

	return router
}
