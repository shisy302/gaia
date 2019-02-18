# Model
model:
  rest_name: sshcertificate
  resource_name: sshcertificates
  entity_name: SSHCertificate
  package: barret
  group: internal/ssh
  description: Internal api to deliver SSH certificate.
  private: true

# Attributes
attributes:
  v1:
  - name: certificate
    description: Contains the signed SSH certificate in OpenSSH Format.
    type: string
    exposed: true
    read_only: true
    autogenerated: true

  - name: extensions
    description: List of extensions to set in the ssh certificate.
    type: external
    exposed: true
    subtype: map_of_string_of_strings

  - name: options
    description: List of options to set in the ssh certificate.
    type: external
    exposed: true
    subtype: map_of_string_of_strings

  - name: principals
    description: List of principals to set in the ssh certificate.
    type: list
    exposed: true
    subtype: string

  - name: publicKey
    description: Contains the public key to sign in OpenSSH Format.
    type: string
    exposed: true
    required: true
    example_value: ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCytT my key

  - name: signerID
    description: The identifier of the CA to use to sign the certificate.
    type: string
    exposed: true
    required: true
    example_value: xxx-xxx-xxx-xxx

  - name: type
    description: Type of SSH certificate.
    type: enum
    exposed: true
    allowed_choices:
    - User
    - Host
    default_value: User

  - name: validity
    description: Set the validity of the SSH certificate.
    type: string
    exposed: true
    default_value: 1h
    validations:
    - $timeDuration
