package types

type Request struct {
	Target  string  `json:"target"`
	Cmd     string  `json:"cmd"`
	Act     Act     `json:"act"`
	Info    Info    `json:"info"`
	Payload Payload `json:"payload"`
}

type Act string

const (
	QUERY  Act = "@Qry"
	MAKE   Act = "@Mak"
	UPDATE Act = "@Udt"
	REMOVE Act = "@Rmv"
)

// {"target":"cnp:/cnp.app", "type":"REQ", "cmd":".test", "act":"action"}
