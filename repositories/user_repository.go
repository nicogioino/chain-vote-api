package repositories

import (
	. "chain-vote-api/models"
	"fmt"
	"github.com/google/uuid"
)

func SaveUser(user *User) (*User, error) {

	err := DB.Create(&user).Error

	if err != nil {
		fmt.Println("Error saving user ")
		return nil, err
	}

	return user, nil
}

func GetUserById(uuid uuid.UUID) (*User, error) {

	user := User{}

	err := DB.Where("id = ?", uuid).First(&user).Error

	if err != nil {
		fmt.Println("Error getting user ")
		return nil, err
	}

	return &user, nil
}

func UpdateUser(user *User, ethAddress string) error {
	err := DB.Model(&user).Update("eth_address", ethAddress).Error
	if err != nil {
		fmt.Println("Error updating user ")
		return err
	}
	return nil
}
