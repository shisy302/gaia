package squallmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

// FlowStatisticMetricValue represents the possible values for attribute "metric".
type FlowStatisticMetricValue string

const (
	// FlowStatisticMetricFlows represents the value Flows.
	FlowStatisticMetricFlows FlowStatisticMetricValue = "Flows"

	// FlowStatisticMetricPorts represents the value Ports.
	FlowStatisticMetricPorts FlowStatisticMetricValue = "Ports"
)

// FlowStatisticTypeValue represents the possible values for attribute "type".
type FlowStatisticTypeValue string

const (
	// FlowStatisticTypeRepartition represents the value Repartition.
	FlowStatisticTypeRepartition FlowStatisticTypeValue = "Repartition"

	// FlowStatisticTypeSerie represents the value Serie.
	FlowStatisticTypeSerie FlowStatisticTypeValue = "Serie"
)

// FlowStatisticModeValue represents the possible values for attribute "mode".
type FlowStatisticModeValue string

const (
	// FlowStatisticModeAccepted represents the value Accepted.
	FlowStatisticModeAccepted FlowStatisticModeValue = "Accepted"

	// FlowStatisticModeRejected represents the value Rejected.
	FlowStatisticModeRejected FlowStatisticModeValue = "Rejected"
)

// FlowStatisticIdentity represents the Identity of the object
var FlowStatisticIdentity = elemental.Identity{
	Name:     "flowstatistic",
	Category: "flowstatistics",
}

// FlowStatisticsList represents a list of FlowStatistics
type FlowStatisticsList []*FlowStatistic

// ContentIdentity returns the identity of the objects in the list.
func (o FlowStatisticsList) ContentIdentity() elemental.Identity {
	return FlowStatisticIdentity
}

// List convert the object to and elemental.IdentifiablesList.
func (o FlowStatisticsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// FlowStatistic represents the model of a flowstatistic
type FlowStatistic struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"-"`

	// DataPoints is a list of time/value pairs that represent the flow events over time.
	DataPoints []map[string]interface{} `json:"dataPoints" bson:"-"`

	// DestinationID is the ID of the destination.
	DestinationID string `json:"destinationID" bson:"-"`

	// DestinationTags contains the tags used to identify destination
	DestinationTags map[string]string `json:"destinationTags" bson:"-"`

	// Metric is the kind of metric the statistic represents.
	Metric FlowStatisticMetricValue `json:"metric" bson:"-"`

	// Mode defines if the metric is for accepted or rejected flows.
	Mode FlowStatisticModeValue `json:"mode" bson:"-"`

	// SourceID is the source of the stats.
	SourceID string `json:"sourceID" bson:"-"`

	// SourceTags contains the tags used to identify the source.
	SourceTags map[string]string `json:"sourceTags" bson:"-"`

	// Type is the type of represenation
	Type FlowStatisticTypeValue `json:"type" bson:"-"`

	// UserIdentifier can be set by the user as a query parameter. It will be returned in the FlowStatistic object.
	UserIdentifier string `json:"userIdentifier" bson:"-"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`
}

// NewFlowStatistic returns a new *FlowStatistic
func NewFlowStatistic() *FlowStatistic {

	return &FlowStatistic{
		ModelVersion: 1.0,
		DataPoints:   []map[string]interface{}{},
		Metric:       "Flows",
		Mode:         "Accepted",
		Type:         "Serie",
	}
}

// Identity returns the Identity of the object.
func (o *FlowStatistic) Identity() elemental.Identity {

	return FlowStatisticIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *FlowStatistic) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *FlowStatistic) SetIdentifier(ID string) {

	o.ID = ID
}

// Version returns the hardcoded version of the model
func (o *FlowStatistic) Version() float64 {

	return 1.0
}

// Doc returns the documentation for the object
func (o *FlowStatistic) Doc() string {
	return `FlowStatistic retrieves the flows between two nodes in a specified interval of time.`
}

func (o *FlowStatistic) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *FlowStatistic) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("metric", string(o.Metric), []string{"Flows", "Ports"}, true); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("mode", string(o.Mode), []string{"Accepted", "Rejected"}, true); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Repartition", "Serie"}, true); err != nil {
		errors = append(errors, err)
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
func (FlowStatistic) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return FlowStatisticAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (FlowStatistic) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return FlowStatisticAttributesMap
}

// FlowStatisticAttributesMap represents the map of attribute for FlowStatistic.
var FlowStatisticAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
		Unique:         true,
	},
	"DataPoints": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `DataPoints is a list of time/value pairs that represent the flow events over time.`,
		Exposed:        true,
		Name:           "dataPoints",
		ReadOnly:       true,
		SubType:        "datapoints_list",
		Type:           "external",
	},
	"DestinationID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `DestinationID is the ID of the destination.`,
		Exposed:        true,
		Format:         "free",
		Name:           "destinationID",
		ReadOnly:       true,
		Type:           "string",
	},
	"DestinationTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `DestinationTags contains the tags used to identify destination`,
		Exposed:        true,
		Name:           "destinationTags",
		ReadOnly:       true,
		SubType:        "selectors_list",
		Type:           "external",
	},
	"Metric": elemental.AttributeSpecification{
		AllowedChoices: []string{"Flows", "Ports"},
		Autogenerated:  true,
		DefaultValue:   FlowStatisticMetricValue("Flows"),
		Description:    `Metric is the kind of metric the statistic represents.`,
		Exposed:        true,
		Name:           "metric",
		ReadOnly:       true,
		Type:           "enum",
	},
	"Mode": elemental.AttributeSpecification{
		AllowedChoices: []string{"Accepted", "Rejected"},
		Autogenerated:  true,
		DefaultValue:   FlowStatisticModeValue("Accepted"),
		Description:    `Mode defines if the metric is for accepted or rejected flows.`,
		Exposed:        true,
		Name:           "mode",
		ReadOnly:       true,
		Type:           "enum",
	},
	"SourceID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `SourceID is the source of the stats.`,
		Exposed:        true,
		Format:         "free",
		Name:           "sourceID",
		ReadOnly:       true,
		Type:           "string",
	},
	"SourceTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `SourceTags contains the tags used to identify the source. `,
		Exposed:        true,
		Name:           "sourceTags",
		ReadOnly:       true,
		SubType:        "selectors_list",
		Type:           "external",
	},
	"Type": elemental.AttributeSpecification{
		AllowedChoices: []string{"Repartition", "Serie"},
		Autogenerated:  true,
		DefaultValue:   FlowStatisticTypeValue("Serie"),
		Description:    `Type is the type of represenation`,
		Exposed:        true,
		Name:           "type",
		ReadOnly:       true,
		Type:           "enum",
	},
	"UserIdentifier": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `UserIdentifier can be set by the user as a query parameter. It will be returned in the FlowStatistic object.`,
		Exposed:        true,
		Format:         "free",
		Name:           "userIdentifier",
		Orderable:      true,
		Type:           "string",
	},
}