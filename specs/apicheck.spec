# Model
model:
  rest_name: apicheck
  resource_name: apichecks
  entity_name: APICheck
  package: squall
  description: |-
    This API allows to verify is a client identitied by his token is allowed to do
    some operations on some apis. For example, it allows third party system to
    impersonate a user and ensure a proxfied request should be allowed.

# Attributes
attributes:
  v1:
  - name: authorized
    description: Authorized contains the results of the check.
    type: external
    exposed: true
    subtype: authorized_identities
    read_only: true
    autogenerated: true

  - name: claims
    description: Claims contains the decoded claims used.
    type: list
    exposed: true
    subtype: string
    read_only: true

  - name: namespace
    description: Namespace is the namespace to use to check the api authentication.
    type: string
    exposed: true
    required: true
    example_value: /namespace

  - name: operation
    description: Operation is the operation you want to check.
    type: enum
    exposed: true
    stored: true
    allowed_choices:
    - Create
    - Delete
    - Info
    - Patch
    - Retrieve
    - RetrieveMany
    - Update
    filterable: true
    orderable: true

  - name: targetIdentities
    description: |-
      TargetIdentities contains the list of identities you want to check the
      authorization.
    type: external
    exposed: true
    subtype: identity_list
    required: true
    example_value:
    - processingunit
    - enforcer

  - name: token
    description: Token is the token to use to check api authentication.
    type: string
    exposed: true
    required: true
    example_value: valid.jwt.token