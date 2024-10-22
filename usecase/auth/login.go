package auth_usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/adhyttungga/go-chatapp-service/helpers"
	"github.com/adhyttungga/go-chatapp-service/models/dto"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func (usecase *AuthUsecaseImpl) Login(c context.Context, req dto.ReqLogin) (httpCode int, res dto.ResUser, token string, err error) {
	user, err := usecase.AuthRepository.GetUser(c, req.UserName)
	if err != nil {
		httpCode := http.StatusInternalServerError
		if err == mongo.ErrNoDocuments {
			err = errors.New("Invalid credential")
			httpCode = http.StatusBadRequest
		}
		msg := fmt.Sprintf("An error occurred while retrieving user: %s", err.Error())
		log.Printf("Error in usecase.AuthUsecase.Login (%s)", msg)
		return httpCode, dto.ResUser{}, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return http.StatusBadRequest, dto.ResUser{}, "", errors.New("Invalid credential")
	}

	// Generate jwt token
	token, err = helpers.GenerateToken(user.ID.Hex())
	if err != nil {
		msg := fmt.Sprintf("An error occurred while generating jwt token: %s", err.Error())
		log.Printf("Error in usecase.AuthUsecase.Login (%s)", msg)
		return http.StatusInternalServerError, dto.ResUser{}, "", err
	}

	res = dto.ResUser{
		ID:         user.ID.Hex(),
		FullName:   user.FullName,
		UserName:   user.UserName,
		ProfilePic: user.ProfilePic,
	}

	return http.StatusOK, res, token, nil
}
