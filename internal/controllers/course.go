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
)

func CreateCourse(c *gin.Context) {
	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set timestamps
	course.CreatedAt = time.Now()
	course.UpdatedAt = time.Now()

	// Insert into database
	collection := config.GetCollection("courses")
	result, err := collection.InsertOne(context.Background(), course)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating course"})
		return
	}

	course.ID = result.InsertedID.(primitive.ObjectID)
	c.JSON(http.StatusCreated, course)
}

func GetCourses(c *gin.Context) {
	collection := config.GetCollection("courses")
	findOptions := options.Find()

	cursor, err := collection.Find(context.Background(), bson.M{}, findOptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while fetching courses"})
		return
	}
	defer cursor.Close(context.Background())

	var courses []models.Course
	if err = cursor.All(context.Background(), &courses); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while decoding courses"})
		return
	}

	c.JSON(http.StatusOK, courses)
}

func GetCourse(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	collection := config.GetCollection("courses")
	var course models.Course
	err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&course)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(http.StatusOK, course)
}

func UpdateCourse(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"name":                course.Name,
			"description":         course.Description,
			"duration":            course.Duration,
			"seats":               course.Seats,
			"eligibilityCriteria": course.EligibilityCriteria,
			"fees":                course.Fees,
			"updated_at":          time.Now(),
		},
	}

	collection := config.GetCollection("courses")
	result, err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id": objectID},
		update,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while updating course"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course updated successfully"})
}

func DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	collection := config.GetCollection("courses")
	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while deleting course"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}
