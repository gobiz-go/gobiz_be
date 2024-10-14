package controller

import (
	"context"
	"encoding/json"
	"gocroot/config"
	"gocroot/model"
	"net/http"
	"time"

	whatsauth "github.com/whatsauth/itmodel"
	
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUsers(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var whatsapi whatsauth.Response

	// Decode the JSON request body into the user struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		whatsapi.Response = "Invalid request payload"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   whatsapi.Response,
			"message": "The JSON request body could not be decoded. Please check the structure of your request.",
		})
		return
	}

	// Validasi masing-masing field
	if user.Username == "" {
		whatsapi.Response = "Nama is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   whatsapi.Response,
			"message": "Please provide a valid name.",
		})
		return
	}

	if user.Password == "" {
		whatsapi.Response = "Phone number is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   whatsapi.Response,
			"message": "Please provide a valid phone number.",
		})
		return
	}

	if user.Email == "" {
		whatsapi.Response = "Email is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   whatsapi.Response,
			"message": "Please provide a valid email address.",
		})
		return
	}

	if user.Phone == "" {
		whatsapi.Response = "Password is required"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   whatsapi.Response,
			"message": "Please provide a password.",
		})
		return
	}

	// Hash the user's password before saving it to the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		whatsapi.Response = "Failed to hash password"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   whatsapi.Response,
			"message": "An error occurred while hashing the password.",
		})
		return
	}
	user.Password = string(hashedPassword)

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	collection := config.Mongoconn.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, bson.M{
		"username":     user.Username,
		"password":   	user.Password,
		"role":       	user.Role,
		"email":		user.Email,
		"phone":     	user.Phone,
		"created_at": 	user.CreatedAt,
		"updated_at": 	user.UpdatedAt,
	})

	if err != nil {
		whatsapi.Response = "Failed to insert user"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   whatsapi.Response,
			"message": "An error occurred while inserting the user into the database.",
		})
		return
	}

	response := map[string]interface{}{
		"message": "User registered successfully",
		"user_id": result.InsertedID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}