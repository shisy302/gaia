package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AutomationTemplateKindValue represents the possible values for attribute "kind".
type AutomationTemplateKindValue string

const (
	// AutomationTemplateKindAction represents the value Action.
	AutomationTemplateKindAction AutomationTemplateKindValue = "Action"

	// AutomationTemplateKindCondition represents the value Condition.
	AutomationTemplateKindCondition AutomationTemplateKindValue = "Condition"
)

// AutomationTemplateIdentity represents the Identity of the object.
var AutomationTemplateIdentity = elemental.Identity{
	Name:     "automationtemplate",
	Category: "automationtemplates",
	Package:  "sephiroth",
	Private:  false,
}

// AutomationTemplatesList represents a list of AutomationTemplates
type AutomationTemplatesList []*AutomationTemplate

// Identity returns the identity of the objects in the list.
func (o AutomationTemplatesList) Identity() elemental.Identity {

	return AutomationTemplateIdentity
}

// Copy returns a pointer to a copy the AutomationTemplatesList.
func (o AutomationTemplatesList) Copy() elemental.Identifiables {

	copy := append(AutomationTemplatesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AutomationTemplatesList.
func (o AutomationTemplatesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AutomationTemplatesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*AutomationTemplate))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AutomationTemplatesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AutomationTemplatesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the AutomationTemplatesList converted to SparseAutomationTemplatesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AutomationTemplatesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseAutomationTemplatesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseAutomationTemplate)
	}

	return out
}

// Version returns the version of the content.
func (o AutomationTemplatesList) Version() int {

	return 1
}

// AutomationTemplate represents the model of a automationtemplate
type AutomationTemplate struct {
	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Contains the entitlements needed for executing the function.
	Entitlements map[string][]elemental.Operation `json:"entitlements" msgpack:"entitlements" bson:"-" mapstructure:"entitlements,omitempty"`

	// Function contains the code.
	Function string `json:"function" msgpack:"function" bson:"-" mapstructure:"function,omitempty"`

	// Contains the unique identifier key for the template.
	Key string `json:"key" msgpack:"key" bson:"-" mapstructure:"key,omitempty"`

	// Represents the kind of template.
	Kind AutomationTemplateKindValue `json:"kind" msgpack:"kind" bson:"-" mapstructure:"kind,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Contains the computed parameters.
	Parameters map[string]interface{} `json:"parameters" msgpack:"parameters" bson:"-" mapstructure:"parameters,omitempty"`

	// Contains all the steps with parameters.
	Steps []*UIStep `json:"steps" msgpack:"steps" bson:"-" mapstructure:"steps,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewAutomationTemplate returns a new *AutomationTemplate
func NewAutomationTemplate() *AutomationTemplate {

	return &AutomationTemplate{
		ModelVersion: 1,
		Entitlements: map[string][]elemental.Operation{},
		Kind:         AutomationTemplateKindCondition,
		Parameters:   map[string]interface{}{},
		Steps:        []*UIStep{},
	}
}

// Identity returns the Identity of the object.
func (o *AutomationTemplate) Identity() elemental.Identity {

	return AutomationTemplateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AutomationTemplate) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AutomationTemplate) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AutomationTemplate) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesAutomationTemplate{}

	s.Description = o.Description
	s.Name = o.Name

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AutomationTemplate) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesAutomationTemplate{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Description = s.Description
	o.Name = s.Name

	return nil
}

// Version returns the hardcoded version of the model.
func (o *AutomationTemplate) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *AutomationTemplate) BleveType() string {

	return "automationtemplate"
}

// DefaultOrder returns the list of default ordering fields.
func (o *AutomationTemplate) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *AutomationTemplate) Doc() string {

	return `Templates that can be used in automations.`
}

func (o *AutomationTemplate) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetDescription returns the Description of the receiver.
func (o *AutomationTemplate) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *AutomationTemplate) SetDescription(description string) {

	o.Description = description
}

// GetName returns the Name of the receiver.
func (o *AutomationTemplate) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *AutomationTemplate) SetName(name string) {

	o.Name = name
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *AutomationTemplate) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAutomationTemplate{
			Description:  &o.Description,
			Entitlements: &o.Entitlements,
			Function:     &o.Function,
			Key:          &o.Key,
			Kind:         &o.Kind,
			Name:         &o.Name,
			Parameters:   &o.Parameters,
			Steps:        &o.Steps,
		}
	}

	sp := &SparseAutomationTemplate{}
	for _, f := range fields {
		switch f {
		case "description":
			sp.Description = &(o.Description)
		case "entitlements":
			sp.Entitlements = &(o.Entitlements)
		case "function":
			sp.Function = &(o.Function)
		case "key":
			sp.Key = &(o.Key)
		case "kind":
			sp.Kind = &(o.Kind)
		case "name":
			sp.Name = &(o.Name)
		case "parameters":
			sp.Parameters = &(o.Parameters)
		case "steps":
			sp.Steps = &(o.Steps)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseAutomationTemplate to the object.
func (o *AutomationTemplate) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAutomationTemplate)
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Entitlements != nil {
		o.Entitlements = *so.Entitlements
	}
	if so.Function != nil {
		o.Function = *so.Function
	}
	if so.Key != nil {
		o.Key = *so.Key
	}
	if so.Kind != nil {
		o.Kind = *so.Kind
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Parameters != nil {
		o.Parameters = *so.Parameters
	}
	if so.Steps != nil {
		o.Steps = *so.Steps
	}
}

// DeepCopy returns a deep copy if the AutomationTemplate.
func (o *AutomationTemplate) DeepCopy() *AutomationTemplate {

	if o == nil {
		return nil
	}

	out := &AutomationTemplate{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *AutomationTemplate.
func (o *AutomationTemplate) DeepCopyInto(out *AutomationTemplate) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy AutomationTemplate: %s", err))
	}

	*out = *target.(*AutomationTemplate)
}

// Validate valides the current information stored into the structure.
func (o *AutomationTemplate) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("kind", string(o.Kind), []string{"Action", "Condition"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = errors.Append(err)
	}

	for _, sub := range o.Steps {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*AutomationTemplate) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AutomationTemplateAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AutomationTemplateLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*AutomationTemplate) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AutomationTemplateAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *AutomationTemplate) ValueForAttribute(name string) interface{} {

	switch name {
	case "description":
		return o.Description
	case "entitlements":
		return o.Entitlements
	case "function":
		return o.Function
	case "key":
		return o.Key
	case "kind":
		return o.Kind
	case "name":
		return o.Name
	case "parameters":
		return o.Parameters
	case "steps":
		return o.Steps
	}

	return nil
}

// AutomationTemplateAttributesMap represents the map of attribute for AutomationTemplate.
var AutomationTemplateAttributesMap = map[string]elemental.AttributeSpecification{
	"Description": {
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Entitlements": {
		AllowedChoices: []string{},
		ConvertedName:  "Entitlements",
		Description:    `Contains the entitlements needed for executing the function.`,
		Exposed:        true,
		Name:           "entitlements",
		SubType:        "_automation_entitlements",
		Type:           "external",
	},
	"Function": {
		AllowedChoices: []string{},
		ConvertedName:  "Function",
		Description:    `Function contains the code.`,
		Exposed:        true,
		Name:           "function",
		Type:           "string",
	},
	"Key": {
		AllowedChoices: []string{},
		ConvertedName:  "Key",
		Description:    `Contains the unique identifier key for the template.`,
		Exposed:        true,
		Name:           "key",
		Type:           "string",
	},
	"Kind": {
		AllowedChoices: []string{"Action", "Condition"},
		ConvertedName:  "Kind",
		DefaultValue:   AutomationTemplateKindCondition,
		Description:    `Represents the kind of template.`,
		Exposed:        true,
		Name:           "kind",
		Type:           "enum",
	},
	"Name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Parameters": {
		AllowedChoices: []string{},
		ConvertedName:  "Parameters",
		Description:    `Contains the computed parameters.`,
		Exposed:        true,
		Name:           "parameters",
		SubType:        "map[string]interface{}",
		Type:           "external",
	},
	"Steps": {
		AllowedChoices: []string{},
		ConvertedName:  "Steps",
		Description:    `Contains all the steps with parameters.`,
		Exposed:        true,
		Name:           "steps",
		SubType:        "uistep",
		Type:           "refList",
	},
}

// AutomationTemplateLowerCaseAttributesMap represents the map of attribute for AutomationTemplate.
var AutomationTemplateLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"description": {
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"entitlements": {
		AllowedChoices: []string{},
		ConvertedName:  "Entitlements",
		Description:    `Contains the entitlements needed for executing the function.`,
		Exposed:        true,
		Name:           "entitlements",
		SubType:        "_automation_entitlements",
		Type:           "external",
	},
	"function": {
		AllowedChoices: []string{},
		ConvertedName:  "Function",
		Description:    `Function contains the code.`,
		Exposed:        true,
		Name:           "function",
		Type:           "string",
	},
	"key": {
		AllowedChoices: []string{},
		ConvertedName:  "Key",
		Description:    `Contains the unique identifier key for the template.`,
		Exposed:        true,
		Name:           "key",
		Type:           "string",
	},
	"kind": {
		AllowedChoices: []string{"Action", "Condition"},
		ConvertedName:  "Kind",
		DefaultValue:   AutomationTemplateKindCondition,
		Description:    `Represents the kind of template.`,
		Exposed:        true,
		Name:           "kind",
		Type:           "enum",
	},
	"name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"parameters": {
		AllowedChoices: []string{},
		ConvertedName:  "Parameters",
		Description:    `Contains the computed parameters.`,
		Exposed:        true,
		Name:           "parameters",
		SubType:        "map[string]interface{}",
		Type:           "external",
	},
	"steps": {
		AllowedChoices: []string{},
		ConvertedName:  "Steps",
		Description:    `Contains all the steps with parameters.`,
		Exposed:        true,
		Name:           "steps",
		SubType:        "uistep",
		Type:           "refList",
	},
}

// SparseAutomationTemplatesList represents a list of SparseAutomationTemplates
type SparseAutomationTemplatesList []*SparseAutomationTemplate

// Identity returns the identity of the objects in the list.
func (o SparseAutomationTemplatesList) Identity() elemental.Identity {

	return AutomationTemplateIdentity
}

// Copy returns a pointer to a copy the SparseAutomationTemplatesList.
func (o SparseAutomationTemplatesList) Copy() elemental.Identifiables {

	copy := append(SparseAutomationTemplatesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAutomationTemplatesList.
func (o SparseAutomationTemplatesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAutomationTemplatesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAutomationTemplate))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAutomationTemplatesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAutomationTemplatesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseAutomationTemplatesList converted to AutomationTemplatesList.
func (o SparseAutomationTemplatesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAutomationTemplatesList) Version() int {

	return 1
}

// SparseAutomationTemplate represents the sparse version of a automationtemplate.
type SparseAutomationTemplate struct {
	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Contains the entitlements needed for executing the function.
	Entitlements *map[string][]elemental.Operation `json:"entitlements,omitempty" msgpack:"entitlements,omitempty" bson:"-" mapstructure:"entitlements,omitempty"`

	// Function contains the code.
	Function *string `json:"function,omitempty" msgpack:"function,omitempty" bson:"-" mapstructure:"function,omitempty"`

	// Contains the unique identifier key for the template.
	Key *string `json:"key,omitempty" msgpack:"key,omitempty" bson:"-" mapstructure:"key,omitempty"`

	// Represents the kind of template.
	Kind *AutomationTemplateKindValue `json:"kind,omitempty" msgpack:"kind,omitempty" bson:"-" mapstructure:"kind,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Contains the computed parameters.
	Parameters *map[string]interface{} `json:"parameters,omitempty" msgpack:"parameters,omitempty" bson:"-" mapstructure:"parameters,omitempty"`

	// Contains all the steps with parameters.
	Steps *[]*UIStep `json:"steps,omitempty" msgpack:"steps,omitempty" bson:"-" mapstructure:"steps,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseAutomationTemplate returns a new  SparseAutomationTemplate.
func NewSparseAutomationTemplate() *SparseAutomationTemplate {
	return &SparseAutomationTemplate{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAutomationTemplate) Identity() elemental.Identity {

	return AutomationTemplateIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAutomationTemplate) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAutomationTemplate) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseAutomationTemplate) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseAutomationTemplate{}

	if o.Description != nil {
		s.Description = o.Description
	}
	if o.Name != nil {
		s.Name = o.Name
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseAutomationTemplate) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseAutomationTemplate{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.Description != nil {
		o.Description = s.Description
	}
	if s.Name != nil {
		o.Name = s.Name
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseAutomationTemplate) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAutomationTemplate) ToPlain() elemental.PlainIdentifiable {

	out := NewAutomationTemplate()
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Entitlements != nil {
		out.Entitlements = *o.Entitlements
	}
	if o.Function != nil {
		out.Function = *o.Function
	}
	if o.Key != nil {
		out.Key = *o.Key
	}
	if o.Kind != nil {
		out.Kind = *o.Kind
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Parameters != nil {
		out.Parameters = *o.Parameters
	}
	if o.Steps != nil {
		out.Steps = *o.Steps
	}

	return out
}

// GetDescription returns the Description of the receiver.
func (o *SparseAutomationTemplate) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseAutomationTemplate) SetDescription(description string) {

	o.Description = &description
}

// GetName returns the Name of the receiver.
func (o *SparseAutomationTemplate) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseAutomationTemplate) SetName(name string) {

	o.Name = &name
}

// DeepCopy returns a deep copy if the SparseAutomationTemplate.
func (o *SparseAutomationTemplate) DeepCopy() *SparseAutomationTemplate {

	if o == nil {
		return nil
	}

	out := &SparseAutomationTemplate{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAutomationTemplate.
func (o *SparseAutomationTemplate) DeepCopyInto(out *SparseAutomationTemplate) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAutomationTemplate: %s", err))
	}

	*out = *target.(*SparseAutomationTemplate)
}

type mongoAttributesAutomationTemplate struct {
	Description string `bson:"description"`
	Name        string `bson:"name"`
}
type mongoAttributesSparseAutomationTemplate struct {
	Description *string `bson:"description,omitempty"`
	Name        *string `bson:"name,omitempty"`
}
