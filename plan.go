package gaia

import (
	"fmt"
	"sync"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PlanIdentity represents the Identity of the object.
var PlanIdentity = elemental.Identity{
	Name:     "plan",
	Category: "plans",
	Package:  "vince",
	Private:  false,
}

// PlansList represents a list of Plans
type PlansList []*Plan

// Identity returns the identity of the objects in the list.
func (o PlansList) Identity() elemental.Identity {

	return PlanIdentity
}

// Copy returns a pointer to a copy the PlansList.
func (o PlansList) Copy() elemental.Identifiables {

	copy := append(PlansList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PlansList.
func (o PlansList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PlansList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Plan))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PlansList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PlansList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the PlansList converted to SparsePlansList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PlansList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o PlansList) Version() int {

	return 1
}

// Plan represents the model of a plan
type Plan struct {
	// Description contains the description of the Plan.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// EnforcerQuota contains the maximum number of enforcers available in the Plan.
	EnforcersQuota int `json:"enforcersQuota" bson:"enforcersquota" mapstructure:"enforcersQuota,omitempty"`

	// Key contains the key identifier of the Plan.
	Key string `json:"key" bson:"key" mapstructure:"key,omitempty"`

	// Name contains the name of the Plan.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// PoliciesQuota contains the maximum number of policies available in the Plan.
	PoliciesQuota int `json:"policiesQuota" bson:"policiesquota" mapstructure:"policiesQuota,omitempty"`

	// ProcessingUnitsQuota contains the maximum PUs available in the Plan.
	ProcessingUnitsQuota int `json:"processingUnitsQuota" bson:"processingunitsquota" mapstructure:"processingUnitsQuota,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewPlan returns a new *Plan
func NewPlan() *Plan {

	return &Plan{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Plan) Identity() elemental.Identity {

	return PlanIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Plan) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Plan) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *Plan) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Plan) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Plan) Doc() string {
	return `Plan contains the various billing plans available.`
}

func (o *Plan) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Plan) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePlan{
			Description:          &o.Description,
			EnforcersQuota:       &o.EnforcersQuota,
			Key:                  &o.Key,
			Name:                 &o.Name,
			PoliciesQuota:        &o.PoliciesQuota,
			ProcessingUnitsQuota: &o.ProcessingUnitsQuota,
		}
	}

	sp := &SparsePlan{}
	for _, f := range fields {
		switch f {
		case "description":
			sp.Description = &(o.Description)
		case "enforcersQuota":
			sp.EnforcersQuota = &(o.EnforcersQuota)
		case "key":
			sp.Key = &(o.Key)
		case "name":
			sp.Name = &(o.Name)
		case "policiesQuota":
			sp.PoliciesQuota = &(o.PoliciesQuota)
		case "processingUnitsQuota":
			sp.ProcessingUnitsQuota = &(o.ProcessingUnitsQuota)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparsePlan to the object.
func (o *Plan) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePlan)
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.EnforcersQuota != nil {
		o.EnforcersQuota = *so.EnforcersQuota
	}
	if so.Key != nil {
		o.Key = *so.Key
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.PoliciesQuota != nil {
		o.PoliciesQuota = *so.PoliciesQuota
	}
	if so.ProcessingUnitsQuota != nil {
		o.ProcessingUnitsQuota = *so.ProcessingUnitsQuota
	}
}

// DeepCopy returns a deep copy if the Plan.
func (o *Plan) DeepCopy() *Plan {

	if o == nil {
		return nil
	}

	out := &Plan{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Plan.
func (o *Plan) DeepCopyInto(out *Plan) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Plan: %s", err))
	}

	*out = *target.(*Plan)
}

// Validate valides the current information stored into the structure.
func (o *Plan) Validate() error {

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
func (*Plan) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PlanAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PlanLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Plan) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PlanAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Plan) ValueForAttribute(name string) interface{} {

	switch name {
	case "description":
		return o.Description
	case "enforcersQuota":
		return o.EnforcersQuota
	case "key":
		return o.Key
	case "name":
		return o.Name
	case "policiesQuota":
		return o.PoliciesQuota
	case "processingUnitsQuota":
		return o.ProcessingUnitsQuota
	}

	return nil
}

// PlanAttributesMap represents the map of attribute for Plan.
var PlanAttributesMap = map[string]elemental.AttributeSpecification{
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Description",
		Description:    `Description contains the description of the Plan.`,
		Exposed:        true,
		Name:           "description",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"EnforcersQuota": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "EnforcersQuota",
		Description:    `EnforcerQuota contains the maximum number of enforcers available in the Plan.`,
		Exposed:        true,
		Name:           "enforcersQuota",
		ReadOnly:       true,
		Stored:         true,
		Type:           "integer",
	},
	"Key": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Key",
		Description:    `Key contains the key identifier of the Plan.`,
		Exposed:        true,
		Name:           "key",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Name",
		Description:    `Name contains the name of the Plan.`,
		Exposed:        true,
		Name:           "name",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"PoliciesQuota": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "PoliciesQuota",
		Description:    `PoliciesQuota contains the maximum number of policies available in the Plan.`,
		Exposed:        true,
		Name:           "policiesQuota",
		ReadOnly:       true,
		Stored:         true,
		Type:           "integer",
	},
	"ProcessingUnitsQuota": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ProcessingUnitsQuota",
		Description:    `ProcessingUnitsQuota contains the maximum PUs available in the Plan.`,
		Exposed:        true,
		Name:           "processingUnitsQuota",
		ReadOnly:       true,
		Stored:         true,
		Type:           "integer",
	},
}

// PlanLowerCaseAttributesMap represents the map of attribute for Plan.
var PlanLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Description",
		Description:    `Description contains the description of the Plan.`,
		Exposed:        true,
		Name:           "description",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"enforcersquota": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "EnforcersQuota",
		Description:    `EnforcerQuota contains the maximum number of enforcers available in the Plan.`,
		Exposed:        true,
		Name:           "enforcersQuota",
		ReadOnly:       true,
		Stored:         true,
		Type:           "integer",
	},
	"key": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Key",
		Description:    `Key contains the key identifier of the Plan.`,
		Exposed:        true,
		Name:           "key",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Name",
		Description:    `Name contains the name of the Plan.`,
		Exposed:        true,
		Name:           "name",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"policiesquota": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "PoliciesQuota",
		Description:    `PoliciesQuota contains the maximum number of policies available in the Plan.`,
		Exposed:        true,
		Name:           "policiesQuota",
		ReadOnly:       true,
		Stored:         true,
		Type:           "integer",
	},
	"processingunitsquota": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ProcessingUnitsQuota",
		Description:    `ProcessingUnitsQuota contains the maximum PUs available in the Plan.`,
		Exposed:        true,
		Name:           "processingUnitsQuota",
		ReadOnly:       true,
		Stored:         true,
		Type:           "integer",
	},
}

// SparsePlansList represents a list of SparsePlans
type SparsePlansList []*SparsePlan

// Identity returns the identity of the objects in the list.
func (o SparsePlansList) Identity() elemental.Identity {

	return PlanIdentity
}

// Copy returns a pointer to a copy the SparsePlansList.
func (o SparsePlansList) Copy() elemental.Identifiables {

	copy := append(SparsePlansList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePlansList.
func (o SparsePlansList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePlansList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePlan))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePlansList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePlansList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparsePlansList converted to PlansList.
func (o SparsePlansList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePlansList) Version() int {

	return 1
}

// SparsePlan represents the sparse version of a plan.
type SparsePlan struct {
	// Description contains the description of the Plan.
	Description *string `json:"description,omitempty" bson:"description" mapstructure:"description,omitempty"`

	// EnforcerQuota contains the maximum number of enforcers available in the Plan.
	EnforcersQuota *int `json:"enforcersQuota,omitempty" bson:"enforcersquota" mapstructure:"enforcersQuota,omitempty"`

	// Key contains the key identifier of the Plan.
	Key *string `json:"key,omitempty" bson:"key" mapstructure:"key,omitempty"`

	// Name contains the name of the Plan.
	Name *string `json:"name,omitempty" bson:"name" mapstructure:"name,omitempty"`

	// PoliciesQuota contains the maximum number of policies available in the Plan.
	PoliciesQuota *int `json:"policiesQuota,omitempty" bson:"policiesquota" mapstructure:"policiesQuota,omitempty"`

	// ProcessingUnitsQuota contains the maximum PUs available in the Plan.
	ProcessingUnitsQuota *int `json:"processingUnitsQuota,omitempty" bson:"processingunitsquota" mapstructure:"processingUnitsQuota,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparsePlan returns a new  SparsePlan.
func NewSparsePlan() *SparsePlan {
	return &SparsePlan{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePlan) Identity() elemental.Identity {

	return PlanIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePlan) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePlan) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparsePlan) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePlan) ToPlain() elemental.PlainIdentifiable {

	out := NewPlan()
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.EnforcersQuota != nil {
		out.EnforcersQuota = *o.EnforcersQuota
	}
	if o.Key != nil {
		out.Key = *o.Key
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.PoliciesQuota != nil {
		out.PoliciesQuota = *o.PoliciesQuota
	}
	if o.ProcessingUnitsQuota != nil {
		out.ProcessingUnitsQuota = *o.ProcessingUnitsQuota
	}

	return out
}

// DeepCopy returns a deep copy if the SparsePlan.
func (o *SparsePlan) DeepCopy() *SparsePlan {

	if o == nil {
		return nil
	}

	out := &SparsePlan{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePlan.
func (o *SparsePlan) DeepCopyInto(out *SparsePlan) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePlan: %s", err))
	}

	*out = *target.(*SparsePlan)
}