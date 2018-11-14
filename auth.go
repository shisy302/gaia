package gaia

import (
	"fmt"
	"sync"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
	"go.aporeto.io/midgard-lib/claims"
)

// AuthIdentity represents the Identity of the object.
var AuthIdentity = elemental.Identity{
	Name:     "auth",
	Category: "auth",
	Package:  "midgard",
	Private:  false,
}

// AuthsList represents a list of Auths
type AuthsList []*Auth

// Identity returns the identity of the objects in the list.
func (o AuthsList) Identity() elemental.Identity {

	return AuthIdentity
}

// Copy returns a pointer to a copy the AuthsList.
func (o AuthsList) Copy() elemental.Identifiables {

	copy := append(AuthsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AuthsList.
func (o AuthsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AuthsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Auth))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AuthsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AuthsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the AuthsList converted to SparseAuthsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AuthsList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o AuthsList) Version() int {

	return 1
}

// Auth represents the model of a auth
type Auth struct {
	// Claims are the claims.
	Claims *claims.MidgardClaims `json:"claims" bson:"-" mapstructure:"claims,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewAuth returns a new *Auth
func NewAuth() *Auth {

	return &Auth{
		ModelVersion: 1,
		Claims:       claims.NewMidgardClaims(),
	}
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
func (o *Auth) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *Auth) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Auth) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Auth) Doc() string {
	return `This API verifies if the given token is valid or not.`
}

func (o *Auth) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Auth) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAuth{
			Claims: &o.Claims,
		}
	}

	sp := &SparseAuth{}
	for _, f := range fields {
		switch f {
		case "claims":
			sp.Claims = &(o.Claims)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseAuth to the object.
func (o *Auth) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAuth)
	if so.Claims != nil {
		o.Claims = *so.Claims
	}
}

// DeepCopy returns a deep copy if the Auth.
func (o *Auth) DeepCopy() *Auth {

	if o == nil {
		return nil
	}

	out := &Auth{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Auth.
func (o *Auth) DeepCopyInto(out *Auth) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Auth: %s", err))
	}

	*out = *target.(*Auth)
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
func (*Auth) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AuthAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AuthLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Auth) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AuthAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Auth) ValueForAttribute(name string) interface{} {

	switch name {
	case "claims":
		return o.Claims
	}

	return nil
}

// AuthAttributesMap represents the map of attribute for Auth.
var AuthAttributesMap = map[string]elemental.AttributeSpecification{
	"Claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Claims",
		Description:    `Claims are the claims.`,
		Exposed:        true,
		Name:           "claims",
		ReadOnly:       true,
		SubType:        "claims",
		Type:           "external",
	},
}

// AuthLowerCaseAttributesMap represents the map of attribute for Auth.
var AuthLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Claims",
		Description:    `Claims are the claims.`,
		Exposed:        true,
		Name:           "claims",
		ReadOnly:       true,
		SubType:        "claims",
		Type:           "external",
	},
}

// SparseAuthsList represents a list of SparseAuths
type SparseAuthsList []*SparseAuth

// Identity returns the identity of the objects in the list.
func (o SparseAuthsList) Identity() elemental.Identity {

	return AuthIdentity
}

// Copy returns a pointer to a copy the SparseAuthsList.
func (o SparseAuthsList) Copy() elemental.Identifiables {

	copy := append(SparseAuthsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAuthsList.
func (o SparseAuthsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAuthsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAuth))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAuthsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAuthsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseAuthsList converted to AuthsList.
func (o SparseAuthsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAuthsList) Version() int {

	return 1
}

// SparseAuth represents the sparse version of a auth.
type SparseAuth struct {
	// Claims are the claims.
	Claims **claims.MidgardClaims `json:"claims,omitempty" bson:"-" mapstructure:"claims,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseAuth returns a new  SparseAuth.
func NewSparseAuth() *SparseAuth {
	return &SparseAuth{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAuth) Identity() elemental.Identity {

	return AuthIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAuth) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAuth) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseAuth) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAuth) ToPlain() elemental.PlainIdentifiable {

	out := NewAuth()
	if o.Claims != nil {
		out.Claims = *o.Claims
	}

	return out
}

// DeepCopy returns a deep copy if the SparseAuth.
func (o *SparseAuth) DeepCopy() *SparseAuth {

	if o == nil {
		return nil
	}

	out := &SparseAuth{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAuth.
func (o *SparseAuth) DeepCopyInto(out *SparseAuth) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAuth: %s", err))
	}

	*out = *target.(*SparseAuth)
}