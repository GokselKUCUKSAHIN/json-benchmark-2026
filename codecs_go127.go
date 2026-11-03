//go:build go1.27

package jsonbenchmark

import (
	jsontext "encoding/json/jsontext"
	jsonv2 "encoding/json/v2"
)

func allCodecs() []codec {
	stdlib := []codec{
		{
			name: "encoding-json-v2/default-safe",
			marshal: func(v any) ([]byte, error) {
				return jsonv2.Marshal(v)
			},
			unmarshal: func(data []byte, v any) error {
				return jsonv2.Unmarshal(data, v)
			},
			compat: true,
		},
		{
			name: "encoding-json-v2/deterministic",
			marshal: func(v any) ([]byte, error) {
				return jsonv2.Marshal(v, jsonv2.Deterministic(true))
			},
			unmarshal: func(data []byte, v any) error {
				return jsonv2.Unmarshal(data, v)
			},
		},
		{
			name: "encoding-json-v2/fast-relaxed",
			marshal: func(v any) ([]byte, error) {
				return jsonv2.Marshal(v,
					jsontext.AllowDuplicateNames(true),
					jsontext.AllowInvalidUTF8(true),
					jsontext.EscapeForHTML(false),
				)
			},
			unmarshal: func(data []byte, v any) error {
				return jsonv2.Unmarshal(data, v,
					jsontext.AllowDuplicateNames(true),
					jsontext.AllowInvalidUTF8(true),
				)
			},
		},
	}
	return append(stdlib, thirdPartyCodecs()...)
}
