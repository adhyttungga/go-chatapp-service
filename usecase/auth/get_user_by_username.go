package auth_usecase

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func (usecase *AuthUsecaseImpl) GetUserByUsername(c context.Context, username string) (err error) {
	err = usecase.AuthRepository.GetUserByUsername(c, username)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			// Print error message to log unless the error is no document found
			msg := fmt.Sprintf("An error occurred while retrieving an existing username: %s", err.Error())
			log.Printf("Error in usecase.AuthUsecase.GetUserByUsername (%s)", msg)
		}
	}

	return err
}
