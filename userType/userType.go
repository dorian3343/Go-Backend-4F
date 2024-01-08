package userType

import (
	"backend/shop"
	"backend/utils"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// base type for users, build off this
type User struct {
	Login    string
	Password string
	Email    string
	ID       string
	Basket   []shop.Product
}

// primitive type to marshal JSON without extra data
type UserLoginPrimitive struct {
	Login    string
	Password string
}

type UserBasketUpdate struct {
	User
	Product_id string
}

func NewUser(login, password, email string) (User, error) {
	if login == "" || password == "" || email == "" {
		return User{}, errors.New("Empty login / password / email.")
	}

	hashedPassword, err := utils.HashString(password)
	if err != nil {
		return User{}, err
	}

	// Generate a unique ID
	id, err := utils.GenerateUserID(login)
	if err != nil {
		return User{}, err
	}

	return User{
		Login:    login,
		Password: string(hashedPassword),
		Email:    email,
		ID:       id,
		Basket:   []shop.Product{},
	}, nil
}

func (u User) Setup(path string) error {
	// Create the file
	f, err := os.Create(path + "/" + u.ID + ".json")
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	// Marshal the entire User struct to JSON
	userJSON, err := json.MarshalIndent(u, "", "    ")
	if err != nil {
		return err
	}

	// Write the JSON data-user to the file
	err = os.WriteFile(f.Name(), userJSON, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (u User) AddToBasket(product shop.Product, path string) error {
	u.Basket = append(u.Basket, product)
	userJSON, err := json.MarshalIndent(u, "", "    ")
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = os.WriteFile(path+"/"+u.ID+".json", userJSON, 0644)
	if err != nil {
		return err
	}

	return nil

}
func (u User) Censor() User {
	u.Password = ""
	return u
}
func (u User) RemoveFromBasket(productToRemove shop.Product, path string) error {
	// Find the index of the product in the basket
	index := -1
	for i, product := range u.Basket {
		if product.Id == productToRemove.Id {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("product not found in basket")
	}

	// Remove the product from the basket
	u.Basket = append(u.Basket[:index], u.Basket[index+1:]...)
	userJSON, err := json.MarshalIndent(u, "", "    ")
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = os.WriteFile(path+"/"+u.ID+".json", userJSON, 0644)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUserData(path string, user User) error {
	// Generate the same ID for the given login
	fmt.Println(user)
	err := os.Remove(path + "/" + user.ID + ".json")
	if err != nil {
		return err
	}

	return nil
}

func ExtractJSONUser(path string) (User, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return User{}, err
	}

	var data User

	err = json.Unmarshal(content, &data)
	if err != nil {
		return User{}, err
	}

	return data, nil
}
