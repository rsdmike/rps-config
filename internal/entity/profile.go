package entity
// type Profile struct {
// 	ProfileName                string               `json:"profileName,omitempty" example:"My Profile"`
// 	AMTPassword                string               `json:"amtPassword,omitempty" example:"my_password"`
// 	GenerateRandomPassword     bool                 `json:"generateRandomPassword,omitempty" example:"true"`
// 	CIRAConfigName             string               `json:"ciraConfigName,omitempty" example:"My CIRA Config"`
// 	Activation                 string               `json:"activation" example:"activate"`
// 	MEBXPassword               string               `json:"mebxPassword,omitempty" example:"my_password"`
// 	GenerateRandomMEBxPassword bool                 `json:"generateRandomMEBxPassword,omitempty" example:"true"`
// 	CIRAConfigObject           *CIRAConfig          `json:"ciraConfigObject,omitempty"`
// 	Tags                       []string             `json:"tags,omitempty" example:"tag1,tag2"`
// 	DhcpEnabled                bool                 `json:"dhcpEnabled,omitempty" example:"true"`
// 	WifiConfigs                []ProfileWifiConfigs `json:"wifiConfigs,omitempty"`
// 	TenantId                   string               `json:"tenantId" example:"abc123"`
// 	TLSMode                    TlsMode              `json:"tlsMode,omitempty" example:"1"`
// 	TLSCerts                   *TLSCerts            `json:"tlsCerts,omitempty"`
// 	TLSSigningAuthority        TlsSigningAuthority  `json:"tlsSigningAuthority,omitempty" example:"SelfSigned"`
// 	UserConsent                AMTUserConsent       `json:"userConsent,omitempty" example:"All"`
// 	IDEREnabled                bool                 `json:"iderEnabled,omitempty" example:"true"`
// 	KVMEnabled                 bool                 `json:"kvmEnabled,omitempty" example:"true"`
// 	SOLEnabled                 bool                 `json:"solEnabled,omitempty" example:"true"`
// 	Ieee8021xProfileName       string               `json:"ieee8021xProfileName,omitempty" example:"My Profile"`
// 	Ieee8021xProfileObject     *Ieee8021xConfig     `json:"ieee8021xProfileObject,omitempty"`
// 	Version                    string               `json:"version,omitempty" example:"1.0.0"`
// }

// type TLSCerts struct {
// 	RootCertificate   CertCreationResult
// 	IssuedCertificate CertCreationResult
// 	Version           string `json:",omitempty"`
// }

type TlsMode int

const (
	InvalidTlsMode    TlsMode = -1
	NoneTlsMode       TlsMode = 0
	ServerOnlyTlsMode TlsMode = 1
	ServerAllowNonTLS TlsMode = 2
	MutualOnlyTlsMode TlsMode = 3
	MutualAllowNonTLS TlsMode = 4
)

type TlsSigningAuthority string

const (
	SelfSignedTlsSigningAuthority  TlsSigningAuthority = "SelfSigned"
	MicrosoftCATlsSigningAuthority TlsSigningAuthority = "MicrosoftCA"
)

type AMTRedirectionServiceEnabledStates int

const (
	DisabledAMTRedirectionServiceEnabledState       AMTRedirectionServiceEnabledStates = 32768
	OnlyIDERAMTRedirectionServiceEnabledState       AMTRedirectionServiceEnabledStates = 32769
	OnlySOLAMTRedirectionServiceEnabledState        AMTRedirectionServiceEnabledStates = 32770
	BothIDERAndSOLAMTRedirectionServiceEnabledState AMTRedirectionServiceEnabledStates = 32771
)
