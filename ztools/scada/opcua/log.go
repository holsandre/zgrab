package opcua

type OPCUALog struct {
	IsOPCUA               bool   `json:"is_opcua"`
	SecurityPolicyUri     string `json:"SecurityPolicyUri,omitempty"`
	ServerTimestamp	     string  `json:"ServerTimestamp,omitempty"`
	ServerProtocolVersion string `json:"ServerProtocolVersion,omitempty"`
	ServerNonce           string `json:"ServerNonce,omitempty"`
}
