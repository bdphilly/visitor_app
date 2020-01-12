package model

import (
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
)

type VisitorValidator interface {
	Validate(*http.Request) error
}

type Visitor struct {
	ID    string `bson:"_id,omitempty"`
	Name  string `bson:"name"`
	Notes string `bson:"notes"`
}

func (v Visitor) Validate(r *http.Request) error {
	fmt.Println("visitor ", v)
	if !govalidator.IsType(v.Notes, "string") {
		return fmt.Errorf("Error validating notes")
	}

	if govalidator.IsNull(v.Name) {
		return fmt.Errorf("Error validating name")
	}
	return nil
}
