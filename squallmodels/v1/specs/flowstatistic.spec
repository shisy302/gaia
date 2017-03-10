{
    "attributes": [
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": true,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "DataPoints is a list of time/value pairs that represent the flow events over time.",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "dataPoints",
            "orderable": false,
            "primary_key": null,
            "read_only": true,
            "required": null,
            "secret": false,
            "setter": null,
            "stored": false,
            "subtype": "datapoints_list",
            "transient": false,
            "type": "external",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": true,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "DestinationID is the ID of the destination.",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": "free",
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "destinationID",
            "orderable": false,
            "primary_key": null,
            "read_only": true,
            "required": null,
            "secret": false,
            "setter": null,
            "stored": false,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": true,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "DestinationTags contains the tags used to identify destination",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "destinationTags",
            "orderable": false,
            "primary_key": null,
            "read_only": true,
            "required": false,
            "secret": false,
            "setter": null,
            "stored": false,
            "subtype": "selectors_list",
            "transient": false,
            "type": "external",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": [
                "Flows",
                "Ports"
            ],
            "autogenerated": true,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": "Flows",
            "deprecated": null,
            "description": "Metric is the kind of metric the statistic represents.",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "metric",
            "orderable": false,
            "primary_key": null,
            "read_only": true,
            "required": false,
            "secret": false,
            "setter": null,
            "stored": false,
            "subtype": null,
            "transient": false,
            "type": "enum",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": [
                "Accepted",
                "Rejected"
            ],
            "autogenerated": true,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": "Accepted",
            "deprecated": null,
            "description": "Mode defines if the metric is for accepted or rejected flows.",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "mode",
            "orderable": false,
            "primary_key": null,
            "read_only": true,
            "required": null,
            "secret": false,
            "setter": null,
            "stored": false,
            "subtype": null,
            "transient": false,
            "type": "enum",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": true,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "SourceID is the source of the stats.",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": "free",
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "sourceID",
            "orderable": false,
            "primary_key": null,
            "read_only": true,
            "required": null,
            "secret": false,
            "setter": null,
            "stored": false,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": true,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "SourceTags contains the tags used to identify the source. ",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "sourceTags",
            "orderable": false,
            "primary_key": null,
            "read_only": true,
            "required": null,
            "secret": false,
            "setter": null,
            "stored": false,
            "subtype": "selectors_list",
            "transient": false,
            "type": "external",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": [
                "Repartition",
                "Serie"
            ],
            "autogenerated": true,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": "Serie",
            "deprecated": null,
            "description": "Type is the type of represenation",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": null,
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "type",
            "orderable": false,
            "primary_key": null,
            "read_only": true,
            "required": null,
            "secret": false,
            "setter": null,
            "stored": false,
            "subtype": null,
            "transient": false,
            "type": "enum",
            "unique": null,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": true,
            "channel": null,
            "creation_only": null,
            "default_order": null,
            "default_value": null,
            "deprecated": null,
            "description": "UserIdentifier can be set by the user as a query parameter. It will be returned in the FlowStatistic object.",
            "exposed": true,
            "filterable": false,
            "foreign_key": null,
            "format": "free",
            "getter": null,
            "identifier": null,
            "index": null,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "userIdentifier",
            "orderable": true,
            "primary_key": null,
            "read_only": null,
            "required": null,
            "secret": false,
            "setter": null,
            "stored": false,
            "subtype": null,
            "transient": false,
            "type": "string",
            "unique": null,
            "uniqueScope": null
        }
    ],
    "children": [],
    "model": {
        "aliases": [],
        "create": false,
        "delete": false,
        "description": "Returns network access statistics on a particular processing unit.",
        "entity_name": "FlowStatistic",
        "extends": [
            "@identifiable-nopk-nostored"
        ],
        "get": false,
        "package": "Visualization",
        "resource_name": "flowstatistics",
        "rest_name": "flowstatistic",
        "root": null,
        "update": false
    }
}