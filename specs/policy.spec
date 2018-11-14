# Model
model:
  rest_name: policy
  resource_name: policies
  entity_name: Policy
  package: squall
  description: Policy represents the policy primitive used by all aporeto policies.
  indexes:
  - - namespace
  - - namespace
    - type
  - - namespace
    - normalizedtags
  - - namespace
    - type
    - allobjecttags
  - - namespace
    - type
    - allsubjecttags
  - - namespace
    - type
    - allobjecttags
    - disabled
  - - namespace
    - type
    - allsubjecttags
    - disabled
  - - namespace
    - type
    - allobjecttags
    - propagated
  - - namespace
    - type
    - allsubjecttags
    - propagated
  - - namespace
    - fallback
  get:
    description: Retrieves the object with the given ID.
  delete:
    description: Deletes the object with the given ID.
    global_parameters:
    - $filtering
  extends:
  - '@base'
  - '@described'
  - '@disabled'
  - '@identifiable-pk-stored'
  - '@metadatable'
  - '@named'
  - '@propagated'
  - '@hidden'
  - '@fallback'
  - '@schedulable'

# Attributes
attributes:
  v1:
  - name: action
    description: Action defines set of actions that must be enforced when a dependency
      is met.
    type: external
    exposed: true
    subtype: actions_list
    stored: true

  - name: allObjectTags
    description: This is a set of all object tags for matching in the DB.
    type: external
    subtype: tags_list
    stored: true

  - name: allSubjectTags
    description: This is a set of all subject tags for matching in the DB.
    type: external
    subtype: tags_list
    stored: true

  - name: object
    description: |-
      Object represents set of entities that another entity depends on. As subjects,
      objects are identified as logical operations on tags when a policy is defined.
    type: external
    exposed: true
    subtype: policies_list
    stored: true
    getter: true
    setter: true

  - name: relation
    description: |-
      Relation describes the required operation to be performed between subjects and
      objects.
    type: external
    exposed: true
    subtype: relations_list
    stored: true

  - name: subject
    description: |-
      Subject represent sets of entities that will have a dependency other entities.
      Subjects are defined as logical operations on tags. Logical operations can
      includes AND/OR.
    type: external
    exposed: true
    subtype: policies_list
    stored: true
    getter: true
    setter: true

  - name: type
    description: Type of the policy.
    type: enum
    exposed: true
    stored: true
    creation_only: true
    allowed_choices:
    - APIAuthorization
    - EnforcerProfile
    - File
    - Hook
    - NamespaceMapping
    - Network
    - ProcessingUnit
    - Quota
    - Service
    - Syscall
    - TokenScope
    - ServiceDependency
    primary_key: true