package auth_usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/adhyttungga/go-chatapp-service/helpers"
	"github.com/adhyttungga/go-chatapp-service/models/dto"
	"github.com/adhyttungga/go-chatapp-service/models/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func (usecase *AuthUsecaseImpl) Login(c context.Context, login dto.ReqLogin) (httpCode int, resUser dto.ResUser, token string, err error) {
	// Validate the request
	if err := usecase.Validate.Struct(login); err != nil {
		// Return bad request and the message error
		return http.StatusBadRequest, dto.ResUser{}, "", err
	}

	userEntity := entity.User{
		UserName: login.UserName,
	}

	// Retrieving user by username
	if err := usecase.UserRepository.FindByUsername(c, &userEntity); err != nil {
		httpCode := http.StatusInternalServerError

		if err == mongo.ErrNoDocuments {
			err = errors.New("Invalid credential")
			httpCode = http.StatusBadRequest
		}

		msg := fmt.Sprintf("An error occurred while retrieving user: %s", err.Error())
		log.Printf("Error in usecase.AuthUsecase.Login (%s)", msg)
		return httpCode, dto.ResUser{}, "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userEntity.Password), []byte(login.Password)); err != nil {
		return http.StatusBadRequest, dto.ResUser{}, "", errors.New("Invalid credential")
	}

	// Generate jwt token
	token, err = helpers.GenerateToken(userEntity.ID.Hex())
	if err != nil {
		msg := fmt.Sprintf("An error occurred while generating jwt token: %s", err.Error())
		log.Printf("Error in usecase.AuthUsecase.Login (%s)", msg)
		return http.StatusInternalServerError, dto.ResUser{}, "", err
	}

	resUser = dto.ResUser{
		ID:         userEntity.ID.Hex(),
		FullName:   userEntity.FullName,
		UserName:   userEntity.UserName,
		ProfilePic: userEntity.ProfilePic,
	}

	return http.StatusOK, resUser, token, nil
}
