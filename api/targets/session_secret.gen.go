// Code generated by "make api"; DO NOT EDIT.
package targets

import "encoding/json"

type SessionSecret struct {
	Raw     json.RawMessage        `json:"raw,omitempty"`
	Decoded map[string]interface{} `json:"decoded,omitempty"`
}