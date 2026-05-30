package types

import "github.com/Sephy314/cnps/pkg/philippines/cebu"

type Info struct {
	Cebu *cebu.Cebu              `json:"cebu,omitempty"`
	Ext  *map[string]interface{} `json:"ext,omitempty"`
}

type Payload map[string]interface{}
