# Model
model:
  rest_name: policyrenderer
  resource_name: policyrenderers
  entity_name: PolicyRenderer
  package: squall
  description: Returns the list of hook policies that matches the given set of tags.
  private: true

# Attributes
attributes:
  v1:
  - name: policies
    description: List of policies rendered for the given set of tags.
    type: external
    exposed: true
    subtype: policy_rules_list
    read_only: true
    autogenerated: true

  - name: tags
    description: List of tags of the object to render the hook policy for.
    type: list
    exposed: true
    subtype: string
    required: true
    example_value:
    - a=a
    - b=b

  - name: type
    description: Type of the policy to render.
    type: enum
    exposed: true
    required: true
    allowed_choices:
    - APIAuthorization
    - EnforcerProfile
    - File
    - Hook
    - NamespaceMapping
    - Network
    - ProcessingUnit
    - Quota
    - Syscall
    - TokenScope
    example_value: APIAuthorization