package logger

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"
	"time"
)

func (l Log) Print() {
	log.SetFlags(0)
	logged := map[string]any{
		"level": l.Level,
		"time":  time.Now(),
		"msg":   l.Msg,
	}

	var buf bytes.Buffer

	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)

	err := encoder.Encode(logged)
	if err != nil {
		log.Printf("Error marshalling logged: %s", err)
		return
	}

	log.Print(strings.TrimSpace(buf.String()))
}

func (r ResponseLog) Print() {
	log.SetFlags(0)
	logged := map[string]any{
		"level":      r.Level,
		"time":       time.Now(),
		"status":     r.Status,
		"msg":        r.Msg,
		"request_id": r.ReqID,
	}

	marshalled, err := json.Marshal(logged)
	if err != nil {
		log.Printf("Error marshalling logged: %s", err)
	}

	log.Printf("%+v", string(marshalled))
}
