package types

type Info struct {
	Ext *map[string]interface{} `json:"ext,omitempty"`
}

type Payload map[string]interface{}
