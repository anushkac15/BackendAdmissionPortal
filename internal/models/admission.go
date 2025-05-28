package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PersonalDetails struct {
	FirstName   string `bson:"firstName" json:"firstName"`
	LastName    string `bson:"lastName" json:"lastName"`
	Email       string `bson:"email" json:"email"`
	Phone       string `bson:"phone" json:"phone"`
	DateOfBirth string `bson:"dateOfBirth" json:"dateOfBirth"`
	Gender      string `bson:"gender" json:"gender"`
	Nationality string `bson:"nationality" json:"nationality"`
	Address     struct {
		Street  string `bson:"street" json:"street"`
		City    string `bson:"city" json:"city"`
		State   string `bson:"state" json:"state"`
		ZipCode string `bson:"zipCode" json:"zipCode"`
		Country string `bson:"country" json:"country"`
	} `bson:"address" json:"address"`
}

type AcademicDetails struct {
	HighestQualification string   `bson:"highestQualification" json:"highestQualification"`
	Institution          string   `bson:"institution" json:"institution"`
	YearOfCompletion     int      `bson:"yearOfCompletion" json:"yearOfCompletion"`
	Percentage           float64  `bson:"percentage" json:"percentage"`
	Documents            []string `bson:"documents" json:"documents"`
}

type Documents struct {
	Photo                     string   `bson:"photo" json:"photo"`
	IDProof                   string   `bson:"idProof" json:"idProof"`
	AddressProof              string   `bson:"addressProof" json:"addressProof"`
	QualificationCertificates []string `bson:"qualificationCertificates" json:"qualificationCertificates"`
}

type Admission struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	StudentID       primitive.ObjectID `bson:"studentId" json:"studentId"`
	CourseID        primitive.ObjectID `bson:"courseId" json:"courseId"`
	PersonalDetails PersonalDetails    `bson:"personalDetails" json:"personalDetails"`
	AcademicDetails AcademicDetails    `bson:"academicDetails" json:"academicDetails"`
	Documents       Documents          `bson:"documents" json:"documents"`
	Status          string             `bson:"status" json:"status"`
	Comments        string             `bson:"comments,omitempty" json:"comments,omitempty"`
	CreatedAt       time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt       time.Time          `bson:"updatedAt" json:"updatedAt"`
}
