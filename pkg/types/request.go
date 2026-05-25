package types

type Request struct {
	Target  string  `json:"target"`
	Type    ReqType `json:"type"`
	Cmd     string  `json:"cmd"`
	Act     string  `json:"act"`
	Info    Info    `json:"info"`
	Payload Payload `json:"payload"`
}

// {"target":"cnp:/cnp.app", "type":"REQ", "cmd":".test", "act":"action"}
