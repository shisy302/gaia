{
    "attributes": [
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "Action defines set of actions that must be enforced when a dependency is met.",
            "exposed": true,
            "filterable": false,
            "foreign_key": false,
            "format": null,
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "action",
            "orderable": false,
            "primary_key": false,
            "read_only": false,
            "required": true,
            "secret": false,
            "setter": false,
            "stored": true,
            "subtype": "actions_list",
            "transient": false,
            "type": "external",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "This is a set of all object tags for matching in the DB",
            "exposed": false,
            "filterable": false,
            "foreign_key": false,
            "format": null,
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "allObjectTags",
            "orderable": false,
            "primary_key": false,
            "read_only": false,
            "required": true,
            "secret": false,
            "setter": false,
            "stored": true,
            "subtype": "tags_list",
            "transient": false,
            "type": "external",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "This is a set of all subject tags for matching in the DB",
            "exposed": false,
            "filterable": false,
            "foreign_key": false,
            "format": null,
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "allSubjectTags",
            "orderable": false,
            "primary_key": false,
            "read_only": false,
            "required": true,
            "secret": false,
            "setter": false,
            "stored": true,
            "subtype": "tags_list",
            "transient": false,
            "type": "external",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "Object represents set of entities that another entity depends on. As subjects, objects are identified as logical operations on tags when a policy is defined.",
            "exposed": true,
            "filterable": false,
            "foreign_key": false,
            "format": null,
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "object",
            "orderable": false,
            "primary_key": false,
            "read_only": false,
            "required": false,
            "secret": false,
            "setter": false,
            "stored": true,
            "subtype": "policies_list",
            "transient": false,
            "type": "external",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "Relation describes the required operation to be performed between subjects and objects",
            "exposed": true,
            "filterable": false,
            "foreign_key": false,
            "format": null,
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "relation",
            "orderable": false,
            "primary_key": false,
            "read_only": false,
            "required": false,
            "secret": false,
            "setter": false,
            "stored": true,
            "subtype": "relations_list",
            "transient": false,
            "type": "external",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": null,
            "autogenerated": false,
            "channel": null,
            "creation_only": false,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "Subject represent sets of entities that will have a dependency other entities. Subjects are defined as logical operations on tags. Logical operations can includes AND/OR",
            "exposed": true,
            "filterable": false,
            "foreign_key": false,
            "format": null,
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "subject",
            "orderable": false,
            "primary_key": false,
            "read_only": false,
            "required": true,
            "secret": false,
            "setter": false,
            "stored": true,
            "subtype": "policies_list",
            "transient": false,
            "type": "external",
            "unique": false,
            "uniqueScope": null
        },
        {
            "allowed_chars": null,
            "allowed_choices": [
                "APIAuthorization",
                "EnforcerProfile",
                "File",
                "Hook",
                "NamespaceMapping",
                "Network",
                "ProcessingUnit",
                "Quota",
                "Syscall",
                "TokenScope"
            ],
            "autogenerated": false,
            "channel": null,
            "creation_only": true,
            "default_order": false,
            "default_value": null,
            "deprecated": false,
            "description": "Type of the policy",
            "exposed": true,
            "filterable": true,
            "foreign_key": false,
            "format": null,
            "getter": false,
            "identifier": false,
            "index": false,
            "max_length": null,
            "max_value": null,
            "min_length": null,
            "min_value": null,
            "name": "type",
            "orderable": false,
            "primary_key": true,
            "read_only": false,
            "required": true,
            "secret": false,
            "setter": false,
            "stored": true,
            "subtype": null,
            "transient": false,
            "type": "enum",
            "unique": false,
            "uniqueScope": null
        }
    ],
    "children": [],
    "model": {
        "aliases": [],
        "create": false,
        "delete": true,
        "description": "[nodoc]",
        "entity_name": "Policy",
        "exposed": true,
        "extends": [
            "@base",
            "@described",
            "@disabled",
            "@identifiable-pk-stored",
            "@metadatable",
            "@named",
            "@propagated",
            "@schedulable"
        ],
        "get": true,
        "package": "squall",
        "private": false,
        "resource_name": "policies",
        "rest_name": "policy",
        "root": null,
        "update": false
    }
}