package routes

import (
	"strings"
	"time"

	"github.com/adhyttungga/go-chatapp-service/config"
	auth_delivery "github.com/adhyttungga/go-chatapp-service/delivery/auth"
	message_delivery "github.com/adhyttungga/go-chatapp-service/delivery/message"
	user_delivery "github.com/adhyttungga/go-chatapp-service/delivery/user"
	"github.com/adhyttungga/go-chatapp-service/middleware"
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

func NewRouter(DB *mongo.Database) *gin.Engine {
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

	// Add validator
	validate := validator.New()

	// Repository
	userRepository := user_repository.NewUserRepository(DB)
	messageRepository := message_repository.NewMessageRepository(DB)

	// Usecase
	authUsecase := auth_usecase.NewAuthUsecase(userRepository, validate)
	messageUsecase := message_usecase.NewMessageUsecase(messageRepository, validate)
	userUsecase := user_usecase.NewUserUsecase(userRepository)

	// Delivery
	authDelivery := auth_delivery.NewAuthDelivery(authUsecase)
	messageDelivery := message_delivery.NewMessageDelivery(messageUsecase)
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
		userRoutes.GET("/", userDelivery.FindAllExcludeId)
	}

	return router
}
