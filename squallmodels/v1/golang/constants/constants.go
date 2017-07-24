package constants

// RenderedPolicyType defines the type of the RenderedPolicyType
type RenderedPolicyType string

const (
	// RenderedPolicyTypeNetwork that the RenderedPolicyType is network
	RenderedPolicyTypeNetwork RenderedPolicyType = "networkPolicyRules"
	// RenderedPolicyTypeFile that the RenderedPolicyType is file
	RenderedPolicyTypeFile RenderedPolicyType = "filePolicyRules"
	// RenderedPolicyTypeSystemCall that the RenderedPolicyType is systemCall
	RenderedPolicyTypeSystemCall RenderedPolicyType = "systemCallPolicyRules"
)

// Vulnerability defines the security vulnerability of an image or a layer
type Vulnerability int

const (
	// VulnerabilityNone defines that security vulnerability is None
	VulnerabilityNone Vulnerability = iota

	// VulnerabilityUnknown defines that security vulnerability is Unknown
	VulnerabilityUnknown

	// VulnerabilityNegligible defines that security vulnerability is Negligible
	VulnerabilityNegligible

	// VulnerabilityLow defines that security vulnerability is Low
	VulnerabilityLow

	// VulnerabilityMedium defines that security vulnerability is Medium
	VulnerabilityMedium

	// VulnerabilityHigh defines that security vulnerability is High
	VulnerabilityHigh

	// VulnerabilityCritical defines that security vulnerability is Critical
	VulnerabilityCritical

	// VulnerabilityDefcon1 defines that security vulnerability is Defcon1
	VulnerabilityDefcon1
)

const (
	// UserLabelPrefix is the label prefix for all user defined labels
	UserLabelPrefix = "@usr:"

	// SystemLabelPrefix is the label prefix for all system defined labels
	SystemLabelPrefix = "@sys:"

	// AttributeLabelPrefix is the prefix for all attributes
	AttributeLabelPrefix = "$"
)

const (

	// StatsTagKeyDestinationID represents the destination ID tag.
	StatsTagKeyDestinationID = "@destid"

	// StatsTagKeyDestinationIP represents the destination IP tag.
	StatsTagKeyDestinationIP = "@destip"

	// StatsTagKeySourceID represents the source ID tag.
	StatsTagKeySourceID = "@srcid"

	// StatsTagKeySourceIP represents the source IP tag.
	StatsTagKeySourceIP = "@srcip"

	// StatsTagKeyDestinationPort represents the destination port tag.
	StatsTagKeyDestinationPort = "@destport"

	// StatsTagKeyProcessingUnitID represents the processing unit ID tag.
	StatsTagKeyProcessingUnitID = "$id"

	// StatsTagKeyIP represents the IP tag.
	StatsTagKeyIP = "@ip"

	// StatsTagKeyAction represents the action tag.
	StatsTagKeyAction = "@action"

	// StatsTagKeyReason represents the reason tag.
	StatsTagKeyReason = "@reason"

	// StatsTagKeyEvent represents the event tag.
	StatsTagKeyEvent = "@event"

	// StatsTagKeyImage represents the image tag.
	StatsTagKeyImage = "@sys:image"

	// StatsTagKeyServerID represents the server ID tag.
	StatsTagKeyServerID = "@serverid"

	// StatsTagKeyNamespace represents the namespace tag.
	StatsTagKeyNamespace = "$namespace"

	// StatsTagKeyAporetoContextID represents the aporeto context ID.
	StatsTagKeyAporetoContextID = "@squallid"

	// StatsTagKeyFileAccessPath represents the file path.
	StatsTagKeyFileAccessPath = "@path"

	// StatsTagKeyFileAccessHost represents the host where the file is stored.
	StatsTagKeyFileAccessHost = "@host"

	// StatsTagKeyFileAccessProtocol represents the protocol used to access the file.
	StatsTagKeyFileAccessProtocol = "@protocol"

	// StatsTagKeyFileAccessMode represents the mode of file.
	StatsTagKeyFileAccessMode = "@mode"

	// StatsTagKeySyscallAccessName represents the number of the syscall.
	StatsTagKeySyscallAccessName = "@syscall"

	// StatsTagKeySyscallAccessPID represents the pid of the related process.
	StatsTagKeySyscallAccessPID = "@pid"

	// StatsTagKeySyscallAccessProcessName represents the name of the related process.
	StatsTagKeySyscallAccessProcessName = "@processname"

	// StatsTagKeyReasonExternalServiceDestination represents the extdst value.
	StatsTagKeyReasonExternalServiceDestination = "extdst"

	// StatsTagKeyReasonExternalServiceSource represents the extsrc value.
	StatsTagKeyReasonExternalServiceSource = "extsrc"
)
