package handler

import (
	"backend/shop"
	"backend/userType"
	"backend/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// the type that the json request expects before adapting it to the main Type
type userPrimitive struct {
	Login    string
	Password string
	Email    string
}

func UserCreation(writer http.ResponseWriter, request *http.Request) {

	var requestData userPrimitive
	err := json.NewDecoder(request.Body).Decode(&requestData)
	if err != nil {
		http.Error(writer, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	user, err := userType.NewUser(requestData.Login, requestData.Password, requestData.Email)
	if err != nil {
		fmt.Errorf("Error adapting data-user.")
		return
	}
	fmt.Println(user)
	err = user.Setup("./data-user")
	if err != nil {
		fmt.Errorf("Error setting up data-user.")
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	responseUser := userType.User{
		Login:  user.Login,
		Email:  user.Email,
		ID:     user.ID,
		Basket: user.Basket,
		//hide password
	}

	JSON, err := json.Marshal(responseUser)
	if err != nil {
		fmt.Errorf("Error returning data-user.")
		return
	}

	_, err = writer.Write(JSON)
	if err != nil {
		fmt.Errorf("Error while sending response.")
		return
	}

}

func UserDeletion(writer http.ResponseWriter, request *http.Request) {
	var LoginData userType.User
	err := json.NewDecoder(request.Body).Decode(&LoginData)
	if err != nil {
		log.Printf("Error decoding JSON: %s", err)
		http.Error(writer, fmt.Sprintf("Error decoding JSON: %s", err), http.StatusBadRequest)
		return
	}

	err = userType.DeleteUserData("./data-user", LoginData)
	if err != nil {
		log.Printf("Error deleting JSON: %s", err)
		http.Error(writer, "Error deleting user data", http.StatusNotFound)
		return
	}

	// Respond with a JSON message upon successful deletion
	response := map[string]string{"message": "User data deleted successfully"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(writer, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(jsonResponse)
	if err != nil {
		fmt.Errorf("Error while sending response.")
		return
	}
}

func HandleBasketAdd(writer http.ResponseWriter, request *http.Request) {
	var data userType.UserBasketUpdate

	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		http.Error(writer, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	userDataFormat := userType.User{
		ID:     data.ID,
		Email:  data.Email,
		Login:  data.Login,
		Basket: data.Basket,
	}
	prod, err := shop.ExtractJSONShop("./data-shop/" + data.Product_id + ".json")
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}
	userDataFormat.AddToBasket(prod, "./data-user/")
}

func HandleUserLogin(writer http.ResponseWriter, request *http.Request) {
	var data userType.UserLoginPrimitive

	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		http.Error(writer, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	C, err := utils.HashString(data.Password)
	if err != nil {
		return
	}
	data.Password = C
	possibleId, err := utils.GenerateUserID(data.Login)

	user, err := userType.ExtractJSONUser("./data-user/" + possibleId + ".json")
	if err != nil {
		fmt.Errorf(err.Error())
		http.Error(writer, "Incorrect User Data", http.StatusUnauthorized)
		return
	}
	fmt.Println(data)
	if user.Password != data.Password || user.Login != data.Login {
		http.Error(writer, "Incorrect User Data", http.StatusUnauthorized)
		return
	}
	responseUser := userType.User{
		Login:  user.Login,
		Email:  user.Email,
		ID:     user.ID,
		Basket: user.Basket,
		//hide password
	}

	JSON, err := json.Marshal(responseUser)
	if err != nil {
		http.Error(writer, "Error encoding JSON", http.StatusBadRequest)
		fmt.Errorf("Error returning data-user.")
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(JSON)
	if err != nil {
		fmt.Errorf("Error while sending response.")
		return
	}
}

func HandleProductRequest(writer http.ResponseWriter, request *http.Request) {
	var products []shop.Product
	folderPath := "./data-shop/"

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("error walking the path %s: %v", folderPath, err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		product, err := shop.ExtractJSONShop(path)
		if err != nil {
			log.Printf("error processing file %s: %v", path, err)
			return nil
		}
		products = append(products, product)
		return nil
	})

	if err != nil {
		log.Printf("error walking the path %s: %v", folderPath, err)
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	JSON, err := json.Marshal(products)
	if err != nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(JSON)
	if err != nil {
		fmt.Errorf("Error while sending response.")
		return
	}
}

func HandleBasketRemove(writer http.ResponseWriter, request *http.Request) {
	var data userType.UserBasketUpdate
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		http.Error(writer, "Error decoding JSON", http.StatusBadRequest)
		return
	}
}
