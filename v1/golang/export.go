package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
)

// ExportIdentity represents the Identity of the object.
var ExportIdentity = elemental.Identity{
	Name:     "export",
	Category: "export",
	Private:  false,
}

// ExportsList represents a list of Exports
type ExportsList []*Export

// Identity returns the identity of the objects in the list.
func (o ExportsList) Identity() elemental.Identity {

	return ExportIdentity
}

// Copy returns a pointer to a copy the ExportsList.
func (o ExportsList) Copy() elemental.Identifiables {

	copy := append(ExportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ExportsList.
func (o ExportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ExportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Export))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ExportsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ExportsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o ExportsList) Version() int {

	return 1
}

// Export represents the model of a export
type Export struct {
	// APIVersion of the api used for the exported data.
	APIVersion int `json:"APIVersion" bson:"-" mapstructure:"APIVersion,omitempty"`

	// List of all exported data.
	Data map[string][]map[string]interface{} `json:"data" bson:"-" mapstructure:"data,omitempty"`

	// The list of identities to export.
	Identities []string `json:"identities" bson:"identities" mapstructure:"identities,omitempty"`

	// Label allows to define a unique label for this export. When importing the
	// content of the export, this label will be added as a tag that will be used to
	// recognize imported object in a later import.
	Label string `json:"label" bson:"-" mapstructure:"label,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewExport returns a new *Export
func NewExport() *Export {

	return &Export{
		ModelVersion: 1,
		Data:         map[string][]map[string]interface{}{},
	}
}

// Identity returns the Identity of the object.
func (o *Export) Identity() elemental.Identity {

	return ExportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Export) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Export) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *Export) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Export) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Export) Doc() string {
	return `Export the policies and related objects in a given namespace.`
}

func (o *Export) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Export) Validate() error {

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
func (*Export) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ExportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ExportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Export) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ExportAttributesMap
}

// ExportAttributesMap represents the map of attribute for Export.
var ExportAttributesMap = map[string]elemental.AttributeSpecification{
	"APIVersion": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "APIVersion",
		Description:    `APIVersion of the api used for the exported data.`,
		Exposed:        true,
		Name:           "APIVersion",
		ReadOnly:       true,
		Type:           "integer",
	},
	"Data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Data",
		Description:    `List of all exported data.`,
		Exposed:        true,
		Name:           "data",
		SubType:        "exported_data_content",
		Type:           "external",
	},
	"Identities": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Identities",
		Description:    `The list of identities to export.`,
		Exposed:        true,
		Name:           "identities",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Label": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Label",
		Description: `Label allows to define a unique label for this export. When importing the
content of the export, this label will be added as a tag that will be used to
recognize imported object in a later import.`,
		Exposed:  true,
		Name:     "label",
		ReadOnly: true,
		Type:     "string",
	},
}

// ExportLowerCaseAttributesMap represents the map of attribute for Export.
var ExportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"apiversion": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "APIVersion",
		Description:    `APIVersion of the api used for the exported data.`,
		Exposed:        true,
		Name:           "APIVersion",
		ReadOnly:       true,
		Type:           "integer",
	},
	"data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Data",
		Description:    `List of all exported data.`,
		Exposed:        true,
		Name:           "data",
		SubType:        "exported_data_content",
		Type:           "external",
	},
	"identities": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Identities",
		Description:    `The list of identities to export.`,
		Exposed:        true,
		Name:           "identities",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"label": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Label",
		Description: `Label allows to define a unique label for this export. When importing the
content of the export, this label will be added as a tag that will be used to
recognize imported object in a later import.`,
		Exposed:  true,
		Name:     "label",
		ReadOnly: true,
		Type:     "string",
	},
}
