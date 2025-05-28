package controllers

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"admission-portal-backend/internal/config"
	"admission-portal-backend/internal/models"
)

func Signup(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Always set role to 'student' for public signup
	student.Role = "student"

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(student.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashing password"})
		return
	}
	student.Password = string(hashedPassword)

	// Set timestamps
	student.CreatedAt = time.Now()
	student.UpdatedAt = time.Now()

	// Insert into database
	collection := config.GetCollection("students")
	result, err := collection.InsertOne(context.Background(), student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating student"})
		return
	}

	student.ID = result.InsertedID.(primitive.ObjectID)
	student.Password = "" // Don't send password back

	c.JSON(http.StatusCreated, student)
}

func Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find student
	collection := config.GetCollection("students")
	var student models.Student
	err := collection.FindOne(context.Background(), bson.M{"email": loginData.Email}).Decode(&student)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(loginData.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": student.ID.Hex(),
		"role":    student.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func GetProfile(c *gin.Context) {
	userID, _ := c.Get("userID")
	objectID, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	collection := config.GetCollection("students")
	var student models.Student
	err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&student)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	student.Password = "" // Don't send password back
	c.JSON(http.StatusOK, student)
}

func UpdateProfile(c *gin.Context) {
	userID, _ := c.Get("userID")
	objectID, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var updateData struct {
		Name        string         `json:"name"`
		Phone       string         `json:"phone"`
		DateOfBirth string         `json:"dateOfBirth"`
		Gender      string         `json:"gender"`
		Address     models.Address `json:"address"`
		Password    string         `json:"password"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	if updateData.Name != "" {
		update["$set"].(bson.M)["name"] = updateData.Name
	}
	if updateData.Phone != "" {
		update["$set"].(bson.M)["phone"] = updateData.Phone
	}
	if updateData.DateOfBirth != "" {
		update["$set"].(bson.M)["dateOfBirth"] = updateData.DateOfBirth
	}
	if updateData.Gender != "" {
		update["$set"].(bson.M)["gender"] = updateData.Gender
	}
	if (updateData.Address != models.Address{}) {
		update["$set"].(bson.M)["address"] = updateData.Address
	}
	if updateData.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateData.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashing password"})
			return
		}
		update["$set"].(bson.M)["password"] = string(hashedPassword)
	}

	collection := config.GetCollection("students")
	result, err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id": objectID},
		update,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while updating profile"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

func CreateAdmin(c *gin.Context) {
	// Check for a secret key in the header
	secret := c.GetHeader("X-Admin-Secret")
	if secret != os.Getenv("ADMIN_SECRET") {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set role to admin
	student.Role = "admin"

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(student.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashing password"})
		return
	}
	student.Password = string(hashedPassword)

	student.CreatedAt = time.Now()
	student.UpdatedAt = time.Now()

	collection := config.GetCollection("students")
	result, err := collection.InsertOne(context.Background(), student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating admin"})
		return
	}

	student.ID = result.InsertedID.(primitive.ObjectID)
	student.Password = ""
	c.JSON(http.StatusCreated, student)
}

func ListAdmins(c *gin.Context) {
	collection := config.GetCollection("students")
	cursor, err := collection.Find(context.Background(), bson.M{"role": "admin"})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching admins"})
		return
	}
	defer cursor.Close(context.Background())

	var admins []models.Student
	if err = cursor.All(context.Background(), &admins); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding admins"})
		return
	}
	for i := range admins {
		admins[i].Password = "" // Don't expose password hashes
	}
	c.JSON(http.StatusOK, admins)
}
