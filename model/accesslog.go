package model

import (
	"encoding/json"
	"time"
)

// Mapping defines the ES mapping for AccesLogEntry logs
const Mapping = `
{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
		"log":{
			"properties":{
				"msg":{
					"type":"text"
				},
				"code":{
					"type": "integer"
				},
				"timestamp":{
					"type":"date"
				}
			}
		}
	}
}
`

// AccessLogEntry is a schema for individual log
type AccessLogEntry struct {
	Msg       string    `json:"msg"`
	Code      int       `json:"code"`
	TimeStamp time.Time `json:"timestamp"`
	encoded   []byte
	err       error
}

func (ale *AccessLogEntry) ensureEncoded() {
	if ale.err == nil && ale.encoded == nil {
		ale.encoded, ale.err = json.Marshal(ale)
	}
}

func (ale *AccessLogEntry) Length() int {
	ale.ensureEncoded()
	return len(ale.encoded)
}
func (ale *AccessLogEntry) Encode() ([]byte, error) {
	ale.ensureEncoded()
	return ale.encoded, ale.err
}
