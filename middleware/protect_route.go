package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/adhyttungga/go-chatapp-service/config"
	"github.com/adhyttungga/go-chatapp-service/models/dto"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ProtectRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("jwt")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ResError{Error: "Token invalid"})
			return
		}

		key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(config.Config.PublicKey))
		if err != nil {
			msg := fmt.Sprintf("An error occurred while parsing public key: %s", err.Error())
			log.Printf("Error in protect route middleware (%s)", msg)
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ResError{Error: msg})
			return
		}

		parsedToken, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
			if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unexpected method: %s", jwtToken.Header["alg"])
			}

			return key, nil
		})
		if err != nil {
			msg := fmt.Sprintf("An error occurred while parsing token: %s", err.Error())
			log.Printf("Error in protect route middleware (%s)", msg)
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ResError{Error: "Token invalid"})
			return
		}

		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok || !parsedToken.Valid {
			err := errors.New("Validate invalid")
			msg := fmt.Sprintf("An error occurred while validating: %s", err.Error())
			log.Printf("Error in protect route middleware (%s)", msg)
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ResError{Error: "Token invalid"})
			return
		}

		c.Set("user_id", claims["dat"])
		return
	}
}
