package types

type Cebu struct {
	Kid     string `json:"kid"`
	AuthPvd string `json:"auth_pvd"`
	Token   string `json:"token"`
}

type Info struct {
	Cebu *Cebu                   `json:"cebu,omitempty"`
	Ext  *map[string]interface{} `json:"ext,omitempty"`
}

type Payload map[string]interface{}

type ResType string

const (
	// ResTypeEr ResTypeRES ResType = "RES"
	ResTypeEr ResType = "ERR"
)

type ReqType string

const (
// ReqTypeREQ ReqType = "REQ"
)
