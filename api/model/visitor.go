package model

import (
	"errors"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VisitorValidator interface {
	Validate() error
}

type Visitor struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Notes       string             `bson:"notes"`
	CreatedAt   time.Time          `bson:"created_at,omitempty"`
	SignOutTime *time.Time         `bson:"sign_out_time,omitempty"`
}

func (v Visitor) Validate() error {
	stripped := strings.Replace(v.Name, " ", "", -1)
	if govalidator.IsNull(stripped) || !govalidator.IsAlpha(stripped) {
		return errors.New("Name is required and must be letters only.")
	}

	return nil
}
