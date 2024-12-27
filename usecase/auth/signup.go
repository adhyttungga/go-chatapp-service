package auth_usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/adhyttungga/go-chatapp-service/helpers"
	"github.com/adhyttungga/go-chatapp-service/models/dto"
	"github.com/adhyttungga/go-chatapp-service/models/entity"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func (usecase *AuthUsecaseImpl) Signup(c context.Context, signup dto.ReqSignup) (resUser dto.ResUser, token string, httpCode int, err error) {
	userEntity := entity.User{
		FullName: signup.FullName,
		UserName: signup.UserName,
		Gender:   signup.Gender,
	}

	// Validate the request
	if err := usecase.Validate.Struct(signup); err != nil {
		// Return bad request and the error message
		return dto.ResUser{}, "", http.StatusBadRequest, err
	}

	// Validate password and confirm password should match
	if signup.Password != signup.ConfirmPassword {
		// Return bad request and the error message if not match
		return dto.ResUser{}, "", http.StatusBadRequest, errors.New("Password doesn't match")
	}

	// Validate username should unique
	if err := usecase.UserRepository.FindByUsername(c, &userEntity); err != nil {
		if err != mongo.ErrNoDocuments {
			// Print error message to log unless the error is no document found
			msg := fmt.Sprintf("An error occurred while retrieving an existing username: %s", err.Error())
			log.Printf("Error in usecase.AuthUsecase.Signup (%s)", msg)
			return dto.ResUser{}, "", http.StatusInternalServerError, err
		}
	} else { // Return bad request when username already exist
		return dto.ResUser{}, "", http.StatusBadRequest, errors.New("Username already exist, please use another username")
	}

	// Generate hash password
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(signup.Password), 10)
	if err != nil {
		msg := fmt.Sprintf("An error occurred while creating hash password: %s", err.Error())
		log.Printf("Error in usecase.AuthUsecase.Signup (%s)", msg)
		return dto.ResUser{}, "", http.StatusInternalServerError, err
	}

	// Generate profile pic base on gender
	boyProfilePic := fmt.Sprintf(`https://avatar.iran.liara.run/public/boy?username=%s`, signup.UserName)
	girlProfilePic := fmt.Sprintf(`https://avatar.iran.liara.run/public/girl?username=%s`, signup.UserName)
	profilePic := boyProfilePic
	if strings.EqualFold(signup.Gender, "female") {
		profilePic = girlProfilePic
	}

	userEntity.Password = string(hashedPwd)
	userEntity.ProfilePic = profilePic

	if err := usecase.UserRepository.Create(c, &userEntity); err != nil {
		msg := fmt.Sprintf("An error occurred while creating user: %s", err.Error())
		log.Printf("Error in usecase.AuthUsecase.Signup (%s)", msg)
		return dto.ResUser{}, "", http.StatusInternalServerError, err
	}

	// Generate jwt token
	token, err = helpers.GenerateToken(userEntity.ID.Hex())
	if err != nil {
		msg := fmt.Sprintf("An error occurred while generating jwt token: %s", err.Error())
		log.Printf("Error in usecase.AuthUsecase.Signup (%s)", msg)
		return dto.ResUser{}, "", http.StatusInternalServerError, err
	}

	resUser = dto.ResUser{
		ID:         userEntity.ID.Hex(),
		FullName:   userEntity.FullName,
		UserName:   userEntity.UserName,
		ProfilePic: userEntity.ProfilePic,
	}

	return resUser, token, http.StatusCreated, nil
}
