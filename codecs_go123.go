//go:build !go1.27

package jsonbenchmark

import stdjson "encoding/json"

func allCodecs() []codec {
	stdlib := codec{
		name:      "encoding-json",
		marshal:   stdjson.Marshal,
		unmarshal: stdjson.Unmarshal,
		compat:    true,
	}
	return append([]codec{stdlib}, thirdPartyCodecs()...)
}
