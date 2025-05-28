package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EligibilityCriteria struct {
	MinimumPercentage float64  `bson:"minimumPercentage" json:"minimumPercentage"`
	RequiredSubjects  []string `bson:"requiredSubjects" json:"requiredSubjects"`
	EntranceExam      bool     `bson:"entranceExam" json:"entranceExam"`
}

type Fees struct {
	TuitionFee   float64 `bson:"tuitionFee" json:"tuitionFee"`
	AdmissionFee float64 `bson:"admissionFee" json:"admissionFee"`
	OtherFees    float64 `bson:"otherFees" json:"otherFees"`
}

type Course struct {
	ID                  primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
	Name                string              `bson:"name" json:"name" binding:"required"`
	Description         string              `bson:"description" json:"description"`
	Duration            string              `bson:"duration" json:"duration"`
	Seats               int                 `bson:"seats" json:"seats"`
	EligibilityCriteria EligibilityCriteria `bson:"eligibilityCriteria" json:"eligibilityCriteria"`
	Fees                Fees                `bson:"fees" json:"fees"`
	CreatedAt           time.Time           `bson:"created_at" json:"created_at"`
	UpdatedAt           time.Time           `bson:"updated_at" json:"updated_at"`
}
