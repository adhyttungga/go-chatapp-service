package auth_usecase

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/adhyttungga/go-chatapp-service/helpers"
	"github.com/adhyttungga/go-chatapp-service/models/dto"
	"github.com/adhyttungga/go-chatapp-service/models/entity"
	"golang.org/x/crypto/bcrypt"
)

func (usecase *AuthUsecaseImpl) Signup(c context.Context, req dto.ReqSignup) (res dto.ResUser, token string, err error) {
	payload := entity.User{
		FullName: req.FullName,
		UserName: req.UserName,
		Gender:   req.Gender,
	}

	// Generate hash password
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		msg := fmt.Sprintf("An error occurred while creating hash password: %s", err.Error())
		log.Printf("Error in usecase.AuthUsecase.Signup (%s)", msg)
		return dto.ResUser{}, "", err
	}

	// Generate profile pic base on gender
	boyProfilePic := fmt.Sprintf(`https://avatar.iran.liara.run/public/boy?username=%s`, req.UserName)
	girlProfilePic := fmt.Sprintf(`https://avatar.iran.liara.run/public/girl?username=%s`, req.UserName)
	profilePic := boyProfilePic
	if strings.EqualFold(req.Gender, "female") {
		profilePic = girlProfilePic
	}

	payload.Password = string(hashedPwd)
	payload.ProfilePic = profilePic

	resID, err := usecase.AuthRepository.CreateUser(c, &payload)
	if err != nil {
		msg := fmt.Sprintf("An error occurred while creating user: %s", err.Error())
		log.Printf("Error in usecase.AuthUsecase.Signup (%s)", msg)
		return dto.ResUser{}, "", err
	}

	// Generate jwt token
	token, err = helpers.GenerateToken(resID)
	if err != nil {
		msg := fmt.Sprintf("An error occurred while generating jwt token: %s", err.Error())
		log.Printf("Error in usecase.AuthUsecase.Signup (%s)", msg)
		return dto.ResUser{}, "", err
	}

	res = dto.ResUser{
		ID:         resID,
		FullName:   payload.FullName,
		UserName:   payload.UserName,
		ProfilePic: payload.ProfilePic,
	}

	return res, token, err
}
