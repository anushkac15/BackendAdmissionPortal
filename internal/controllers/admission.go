package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"admission-portal-backend/internal/config"
	"admission-portal-backend/internal/models"
	"log"
)

func ApplyAdmission(c *gin.Context) {
	// Extract userID from JWT (set by middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	studentID, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var req struct {
		CourseID        string                 `json:"courseId"`
		PersonalDetails models.PersonalDetails `json:"personalDetails"`
		AcademicDetails models.AcademicDetails `json:"academicDetails"`
		Documents       models.Documents       `json:"documents"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	courseID, err := primitive.ObjectIDFromHex(req.CourseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	admission := models.Admission{
		StudentID:       studentID,
		CourseID:        courseID,
		PersonalDetails: req.PersonalDetails,
		AcademicDetails: req.AcademicDetails,
		Documents:       req.Documents,
		Status:          "pending",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	// Insert admission into DB...
	collection := config.GetCollection("admissions")
	result, err := collection.InsertOne(context.Background(), admission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while applying for admission"})
		return
	}
	admission.ID = result.InsertedID.(primitive.ObjectID)
	c.JSON(http.StatusCreated, admission)
}

func GetAdmissions(c *gin.Context) {
	userID, _ := c.Get("userID")
	studentID, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	collection := config.GetCollection("admissions")
	findOptions := options.Find()

	cursor, err := collection.Find(context.Background(), bson.M{"studentId": studentID}, findOptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while fetching admissions"})
		return
	}
	defer cursor.Close(context.Background())

	var admissions []models.Admission
	if err = cursor.All(context.Background(), &admissions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while decoding admissions"})
		return
	}

	c.JSON(http.StatusOK, admissions)
}

func GetAdmission(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admission ID"})
		return
	}

	userID, _ := c.Get("userID")
	studentID, err := primitive.ObjectIDFromHex(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	collection := config.GetCollection("admissions")
	var admission models.Admission
	err = collection.FindOne(context.Background(), bson.M{
		"_id":       objectID,
		"studentId": studentID,
	}).Decode(&admission)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admission not found"})
		return
	}

	c.JSON(http.StatusOK, admission)
}

func UpdateAdmissionStatus(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admission ID"})
		return
	}

	var updateData struct {
		Status   string `json:"status" binding:"required,oneof=pending approved rejected"`
		Comments string `json:"comments"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"status":     updateData.Status,
			"comments":   updateData.Comments,
			"updated_at": time.Now(),
		},
	}

	collection := config.GetCollection("admissions")
	result, err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id": objectID},
		update,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while updating admission"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admission not found"})
		return
	}

	// Log notification
	log.Printf("Notification: Admission status updated. AdmissionID: %s, NewStatus: %s, Comments: %s", id, updateData.Status, updateData.Comments)

	c.JSON(http.StatusOK, gin.H{"message": "Admission status updated successfully"})
}
