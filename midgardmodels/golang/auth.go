package midgardmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

// AuthIdentity represents the Identity of the object
var AuthIdentity = elemental.Identity{
	Name:     "auth",
	Category: "auth",
}

// AuthsList represents a list of Auths
type AuthsList []*Auth

// Auth represents the model of a auth
type Auth struct {
	// Claims are the claims.
	Claims interface{} `json:"claims" cql:"-" bson:"-"`
}

// NewAuth returns a new *Auth
func NewAuth() *Auth {

	return &Auth{}
}

// Identity returns the Identity of the object.
func (o *Auth) Identity() elemental.Identity {

	return AuthIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Auth) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Auth) SetIdentifier(ID string) {

}

func (o *Auth) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Auth) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (Auth) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return AuthAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (Auth) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AuthAttributesMap
}

// AuthAttributesMap represents the map of attribute for Auth.
var AuthAttributesMap = map[string]elemental.AttributeSpecification{
	"Claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Claims are the claims.`,
		Exposed:        true,
		Name:           "claims",
		ReadOnly:       true,
		SubType:        "claims",
		Type:           "external",
	},
}
