package jsonbenchmark

import (
	"errors"

	"github.com/bytedance/sonic"
	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
	segmentjson "github.com/segmentio/encoding/json"
)

type codec struct {
	name         string
	marshal      func(any) ([]byte, error)
	unmarshal    func([]byte, any) error
	concreteOnly bool
	compat       bool
}

var (
	jsoniterCompat = jsoniter.ConfigCompatibleWithStandardLibrary
	jsoniterFast   = jsoniter.ConfigFastest
	sonicCompat    = sonic.ConfigStd
	sonicFast      = sonic.ConfigFastest
)

// thirdPartyCodecs lists libraries shared by both Go 1.23 and Go 1.27 suites.
func thirdPartyCodecs() []codec {
	return []codec{
		{
			name:      "jsoniter/compat",
			marshal:   jsoniterCompat.Marshal,
			unmarshal: jsoniterCompat.Unmarshal,
			compat:    true,
		},
		{
			name:      "jsoniter/fast",
			marshal:   jsoniterFast.Marshal,
			unmarshal: jsoniterFast.Unmarshal,
		},
		{
			name:      "go-json/compat",
			marshal:   gojson.Marshal,
			unmarshal: gojson.Unmarshal,
			compat:    true,
		},
		{
			name: "go-json/fast",
			marshal: func(v any) ([]byte, error) {
				return gojson.MarshalWithOption(v, gojson.UnorderedMap(), gojson.DisableHTMLEscape())
			},
			unmarshal: func(data []byte, v any) error {
				return gojson.UnmarshalNoEscape(data, v)
			},
		},
		{
			name:      "segmentio/compat",
			marshal:   segmentjson.Marshal,
			unmarshal: segmentjson.Unmarshal,
			compat:    true,
		},
		{
			name: "segmentio/fast",
			marshal: func(v any) ([]byte, error) {
				return segmentjson.Append(nil, v, 0)
			},
			unmarshal: func(data []byte, v any) error {
				rest, err := segmentjson.Parse(data, v, segmentjson.ZeroCopy)
				if err == nil && len(rest) != 0 {
					return errors.New("segmentio: trailing JSON data")
				}
				return err
			},
		},
		{
			name:      "sonic/compat",
			marshal:   sonicCompat.Marshal,
			unmarshal: sonicCompat.Unmarshal,
			compat:    true,
		},
		{
			name:      "sonic/fast",
			marshal:   sonicFast.Marshal,
			unmarshal: sonicFast.Unmarshal,
		},
		{
			name: "easyjson/compat",
			marshal: func(v any) ([]byte, error) {
				m, ok := v.(easyjson.Marshaler)
				if !ok {
					return nil, errors.New("easyjson: unsupported non-generated type")
				}
				return easyjson.Marshal(m)
			},
			unmarshal: func(data []byte, v any) error {
				u, ok := v.(easyjson.Unmarshaler)
				if !ok {
					return errors.New("easyjson: unsupported non-generated type")
				}
				return easyjson.Unmarshal(data, u)
			},
			concreteOnly: true,
			compat:       true,
		},
		{
			name: "easyjson/fast",
			marshal: func(v any) ([]byte, error) {
				m, ok := v.(interface{ MarshalEasyJSON(*jwriter.Writer) })
				if !ok {
					return nil, errors.New("easyjson: unsupported non-generated type")
				}
				var w jwriter.Writer
				w.NoEscapeHTML = true
				m.MarshalEasyJSON(&w)
				return w.BuildBytes()
			},
			unmarshal: func(data []byte, v any) error {
				u, ok := v.(interface{ UnmarshalEasyJSON(*jlexer.Lexer) })
				if !ok {
					return errors.New("easyjson: unsupported non-generated type")
				}
				l := jlexer.Lexer{Data: data}
				u.UnmarshalEasyJSON(&l)
				return l.Error()
			},
			concreteOnly: true,
		},
	}
}
